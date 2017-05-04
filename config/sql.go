package config

import (
    "fmt"
    "github.com/go-ini/ini"
)

type SqlConfig struct {
    Host string
    Username string
    Password string
    Database string
}

func GenerateSqlConfig(conf *ini.Section, err error) *SqlConfig {
    if (err != nil) {
        fmt.Println(err)
    }

    host, err := conf.GetKey("HOST")
    if (err != nil) {
        fmt.Println(err)
    }

    user, err := conf.GetKey("USERNAME")
    if (err != nil) {
        fmt.Println(err)
    }

    pass, err := conf.GetKey("PASSWORD")
    if (err != nil) {
        fmt.Println(err)
    }

    db, err := conf.GetKey("DATABASE")
    if (err != nil) {
        fmt.Println(err)
    }

    return &SqlConfig{
        Host: host.String(),
        Username: user.String(),
        Password: pass.String(),
        Database: db.String(),
    }
}
