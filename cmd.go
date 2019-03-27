package main

import (
	"github.com/joms/visma-export/config"
	"github.com/joms/visma-export/invoice"
	"github.com/joms/visma-export/sql"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	conf := config.GetConfig()
	conn := sql.Connect(conf.SQL)
	invoice.Export(conn, conf.SQL, conf.Misc)
}
