package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

// Config structure for the SQL connection
type Config struct {
	SQL  *SQLConfig
	Misc *MiscConfig
}

// GetConfig fetches configuration from the ini configuration file
func GetConfig() *Config {
	conf := new(Config)

	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
	}

	sql := GenerateSQLConfig(cfg.GetSection("sql"))
	conf.SQL = sql

	misc := GenerateMiscConfig(cfg.GetSection("misc"))
	conf.Misc = misc

	return conf
}
