package sql

import (
	"fmt"
	"strconv"
	"visma-export/config"

	"github.com/jinzhu/gorm"
)

// Connect to the SQL database
func Connect(conf *config.SQLConfig) *gorm.DB {
	dbCon, err := gorm.Open("mssql", "sqlserver://"+conf.Username+":"+conf.Password+"@"+conf.Host+":"+strconv.Itoa(conf.Port)+"?database="+conf.Database)
	if err != nil {
		fmt.Println("From Connect() attempt: " + err.Error())
	}

	return dbCon
}
