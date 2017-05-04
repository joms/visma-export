package invoice

import (
    "github.com/jinzhu/gorm"
    "fmt"
    "encoding/csv"
    "os"
)

var (
    db *gorm.DB
)

func Export(dbCon *gorm.DB) {
    var reportLines []Report
    db = dbCon

    reportRows, err := db.Raw(`Select o.OrderID,c.CustomerNo,o.DeliveryDate
                        From Orders as o
                        left join Customers as c on O.CustomerID = c.CustomerID
                        where o.OrderType = 2`).Rows()
    defer reportRows.Close()
    if (err != nil) {
        fmt.Println(err)
    }

    for reportRows.Next() {
        var report Report

        var orderHeading OrderHeading
        db.ScanRows(reportRows, &orderHeading)
        report.H = orderHeading

        report.A = getAddressRow(report.H.OrderCSOrdNo)
        report.L = getInvoices(report.H.OrderCSOrdNo)
        getInvoices(report.H.OrderCSOrdNo)

        reportLines = append(reportLines, report)
    }

    printInvoice(reportLines)
    //fmt.Printf("%+v\n", reportLines)
}

func getAddressRow(orderID string) AddressLine{
    var addressHeading AddressLine
        db.Raw(`Select c.CustomerNo,C.Name,C.MailingAddress,C.MailingZip,C.MailingCity,YourRef,C.Phone,C.Email,e.EmployeeID, p.Days,projects.ProjectNo From Orders
                left join Customers as c on orders.customerId = c.CustomerID
                left join PaymentTerms as p on orders.PaymentTermId = p.PaymentTermID
                left join Employees as e on orders.EmployeeID = e.EmployeeID
                left join Projects as projects on orders.ProjectID = projects.ProjectID
                where orderType = 2 and OrderID = ?`, orderID).Scan(&addressHeading)

    return addressHeading
}

func getInvoices(OrderID string) []InvoiceLine {
    var invoices []InvoiceLine

    invoiceRows, err := db.Raw(`Select a.ArticleNo,lines.Description,lines.Count,lines.GrossPrice
                                from OrderLines as lines
                                left join Articles as a on lines.ArticleID = a.ArticleID
                                WHERE lines.OrderID = ?`, OrderID).Rows()
    defer invoiceRows.Close()
    if (err != nil) {
        fmt.Println(err)
    }

    for invoiceRows.Next() {
        var invoice InvoiceLine
        db.ScanRows(invoiceRows, &invoice)
        invoices = append(invoices, invoice)
    }

    return invoices
}

func printInvoice(report []Report) {
    file, err := os.Create("result.edi")
    if err != nil {
        fmt.Println(err)
    }

    writer := csv.NewWriter(file)
    writer.Comma = ';'

    writer.Write([]string{"1", "1"})

    for _, elem := range report {
        err := writer.Write(elem.A.ToSlice())
        if (err != nil) {
            fmt.Println(err)
        }

        writer.Write(elem.H.ToSlice())

        for _, l := range elem.L {
            writer.Write(l.ToSlice())
        }
    }

    writer.Flush()
}

func structToSlice() {

}

func outputLine() []string {
    var result []string

    return result
}
