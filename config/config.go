package config

import (
    "fmt"
    "github.com/go-ini/ini"
)

type Config struct {
    SQL *SqlConfig
}

func GetConfig() *Config {
    conf := new(Config)

    cfg, err := ini.Load("config.ini")
    if (err != nil) {
        fmt.Println(err)
    }

    sql := GenerateSQLConfig(cfg.GetSection("sql"))
    conf.SQL = sql

    return conf
}
