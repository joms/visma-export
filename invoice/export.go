package invoice

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"visma-export/config"

	"github.com/jinzhu/gorm"
)

var (
	db          *gorm.DB
	invoiceList []string
)

// Export invoices into Visma readable format
func Export(dbCon *gorm.DB, dbConf *config.SQLConfig, miscConf *config.MiscConfig) {
	var reportLines []report
	db = dbCon

	invoiceList = loadList()

	reportRows, err := db.Raw(`Select o.OrderID,c.CustomerNo,o.DeliveryDate,o.OrderDate,d.DepartmentNo,d.Name AS DepartmentName,o.Reference
                        From Orders as o
                        left join Customers as c on O.CustomerID = c.CustomerID
                        left join Departments as d on d.DepartmentID = o.DepartmentID
						where o.OrderType = 2 and o.OrderDate >= ?`, dbConf.Oldest).Rows()

	defer reportRows.Close()
	if err != nil {
		fmt.Println(err)
	}

	for reportRows.Next() {
		var reportRow report

		var orderHeading orderHeading
		db.ScanRows(reportRows, &orderHeading)
		reportRow.H = orderHeading
		reportRow.H.HeadStatus = "-1"
		reportRow.H.OrderType = "2"
		reportRow.H.OrderGr5 = reportRow.H.OrderCSOrdNo
		reportRow.H.Semicolon3 = miscConf.CashRegister
		reportRow.H.OrderDelDt = orderDateToVismaDate(reportRow.H.OrderDelDt)

		if isInvoiceDone(reportRow.H.OrderCSOrdNo) {
			continue
		}

		reportRow.A = getAddressRow(reportRow.H.OrderCSOrdNo)
		reportRow.L = getInvoices(reportRow.H.OrderCSOrdNo)

		getInvoices(reportRow.H.OrderCSOrdNo)

		reportLines = append(reportLines, reportRow)
		saveDone(reportRow.H.OrderCSOrdNo)
	}

	printInvoice(miscConf.SaveDir, dbConf.Database, reportLines)
	writeList(invoiceList)
}

func getAddressRow(orderID string) addressLine {
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

	invoiceRows, err := db.Raw(`Select a.ArticleNo,lines.Description,lines.Count,lines.GrossPrice,lines.NetAmount,a.SalesPrice
                                from OrderLines as lines
                                left join Articles as a on lines.ArticleID = a.ArticleID
                                WHERE lines.OrderID = ?`, OrderID).Rows()
	defer invoiceRows.Close()
	if err != nil {
		fmt.Println(err)
	}

	for invoiceRows.Next() {
		var invoice invoiceLine
		db.ScanRows(invoiceRows, &invoice)
		invoice.RecType = "-1"
		invoices = append(invoices, invoice)
	}

	return invoices
}

// Magic
func printInvoice(path string, dbname string, report []report) {
	t := time.Now()
	now := t.Format("2006-01-02T1504")

	if len(report) > 0 {
		// Create our file and ensure that it is empty
		file, err := os.Create(path + "/" + dbname + now + ".edi")
		if err != nil {
			fmt.Println(err)
		}

		// Create a new CSV writer
		writer := csv.NewWriter(file)
		writer.Comma = ';'
		writer.UseCRLF = true

		// Write the header line
		writer.Write([]string{"1", "1"})

		// Write actual content
		for _, elem := range report {
			// Write address header line
			writer.Write(elem.A.ToSlice())

			// Write order Head line
			writer.Write(elem.H.ToSlice())

			// Loop through order lines
			for _, l := range elem.L {
				// Write order line
				writer.Write(l.ToSlice())
			}
		}
		writer.Flush()
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
