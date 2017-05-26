package invoice

import (
    "github.com/jinzhu/gorm"
    "fmt"
    "encoding/csv"
    "os"
    "invoice_export/config"
)

var (
    db *gorm.DB
    invoiceList []string
)

// Export invoices into Visma readable format
func Export(dbCon *gorm.DB, dbConf *config.SQLConfig) {
    var reportLines []report
    db = dbCon

    invoiceList = loadList()

    reportRows, err := db.Raw(`Select o.OrderID,c.CustomerNo,o.DeliveryDate
                        From Orders as o
                        left join Customers as c on O.CustomerID = c.CustomerID
                        where o.OrderType = 2 and o.OrderDate >= ?`, dbConf.Oldest).Rows()
    defer reportRows.Close()
    if (err != nil) {
        fmt.Println(err)
    }

    for reportRows.Next() {
        var reportRow report

        var orderHeading orderHeading
        db.ScanRows(reportRows, &orderHeading)
        reportRow.H = orderHeading
        reportRow.H.HeadStatus = "-1"
        reportRow.H.OrderType = "2"

        if (isInvoiceDone(reportRow.H.OrderCSOrdNo)) {
            continue
        }

        reportRow.A = getAddressRow(reportRow.H.OrderCSOrdNo)
        reportRow.L = getInvoices(reportRow.H.OrderCSOrdNo)
        getInvoices(reportRow.H.OrderCSOrdNo)

        reportLines = append(reportLines, reportRow)
        saveDone(reportRow.H.OrderCSOrdNo)
    }

    printInvoice(reportLines)
    writeList(invoiceList)
}

func getAddressRow(orderID string) addressLine{
    var addressHeading addressLine
        db.Raw(`Select c.CustomerNo,C.Name,C.MailingAddress,C.MailingZip,C.MailingCity,YourRef,C.Phone,C.Email,e.EmployeeID, p.Days,projects.ProjectNo From Orders
                left join Customers as c on orders.customerId = c.CustomerID
                left join PaymentTerms as p on orders.PaymentTermId = p.PaymentTermID
                left join Employees as e on orders.EmployeeID = e.EmployeeID
                left join Projects as projects on orders.ProjectID = projects.ProjectID
                where orderType = 2 and OrderID = ?`, orderID).Scan(&addressHeading)

    return addressHeading
}

func getInvoices(OrderID string) []invoiceLine {
    var invoices []invoiceLine

    invoiceRows, err := db.Raw(`Select a.ArticleNo,lines.Description,lines.Count,lines.GrossPrice
                                from OrderLines as lines
                                left join Articles as a on lines.ArticleID = a.ArticleID
                                WHERE lines.OrderID = ?`, OrderID).Rows()
    defer invoiceRows.Close()
    if (err != nil) {
        fmt.Println(err)
    }

    for invoiceRows.Next() {
        var invoice invoiceLine
        db.ScanRows(invoiceRows, &invoice)
        invoices = append(invoices, invoice)
    }

    return invoices
}

func printInvoice(report []report) {
    file, err := os.Create("result.edi")
    if err != nil {
        fmt.Println(err)
    }

    writer := csv.NewWriter(file)
    writer.Comma = ';'

    writer.Write([]string{"1", "1"})
    writer.Flush()

    for i := 0; i < len(report); i++ {
        elem := report[i]

        file.WriteString(elem.A.ToCSV()+"\n")
        file.WriteString(elem.H.ToCSV()+"\n")

        for x := 0; x < len(elem.L); x++ {
            nl := "\n"
            if x == len(elem.L) - 1 && i == len(report) - 1 {
                nl = ""
            }

            file.WriteString(elem.L[x].ToCSV()+nl)
        }
    }

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
}

func isInvoiceDone(invoiceNumber string) bool {
    for _, b := range invoiceList {
        if b == invoiceNumber {
            return true
        }
    }

    return false
}

func saveDone(invoiceNumber string) {
    invoiceList = append(invoiceList, invoiceNumber)
}
