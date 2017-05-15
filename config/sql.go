package config

import (
    "fmt"
    "github.com/go-ini/ini"
)

// SQLConfig structure
type SQLConfig struct {
    Host string
    Username string
    Password string
    Database string
    Oldest string
}

// GenerateSQLConfig from the ini configuration data
func GenerateSQLConfig(conf *ini.Section, err error) *SQLConfig {
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

    oldest, err := conf.GetKey("OLDESTORDER")
    if (err != nil) {
        fmt.Println(err)
    }

    return &SQLConfig{
        Host: host.String(),
        Username: user.String(),
        Password: pass.String(),
        Database: db.String(),
        Oldest: oldest.String(),
    }
}
