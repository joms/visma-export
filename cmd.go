package main

import (
    _ "github.com/jinzhu/gorm/dialects/mssql"
    "invoice_export/config"
    "invoice_export/sql"
    "invoice_export/invoice"
)

func main() {
    conf := config.GetConfig()
    conn := sql.Connect(conf.SQL)
    invoice.Export(conn, conf.SQL)
}
