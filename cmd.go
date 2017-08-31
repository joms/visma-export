package main

import (
	"visma-export/config"
	"visma-export/invoice"
	"visma-export/sql"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	conf := config.GetConfig()
	conn := sql.Connect(conf.SQL)
	invoice.Export(conn, conf.SQL, conf.Misc)
}
