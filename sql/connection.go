package sql

import (
    "github.com/jinzhu/gorm"
    "invoice_export/config"
    "fmt"
)

// Connect to the SQL database
func Connect(conf *config.SQLConfig) *gorm.DB {
  dbCon, err := gorm.Open("mssql", "sqlserver://"+conf.Username+":"+conf.Password+"@"+conf.Host+":1433?database="+conf.Database)
  if err != nil {
      fmt.Println("From Connect() attempt: " + err.Error())
  }

  return dbCon
}
