package sql

import (
    "github.com/jinzhu/gorm"
    "invoice_export/config"
    "fmt"
)

func Connect(conf *config.SqlConfig) *gorm.DB {
  dbCon, err := gorm.Open("mssql", "sqlserver://"+conf.Username+":"+conf.Password+"@"+conf.Host+":1433?database="+conf.Database)
  if err != nil {
      fmt.Println("From Connect() attempt: " + err.Error())
  }

  return dbCon
}
