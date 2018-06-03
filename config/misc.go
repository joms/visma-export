package config

import (
  "fmt"
  "github.com/go-ini/ini"
)

// Misc configuration structure
type MiscConfig struct {
  SaveDir string
  CashRegister string
}

func GenerateMiscConfig(conf *ini.Section, err error) *MiscConfig {
  if (err != nil) {
    fmt.Println(err)
  }

  savedir, err := conf.GetKey("SAVEDIR")
  if (err != nil) {
    fmt.Println(err)
  }

  cashregister, err := conf.GetKey("CASHREGISTER")
  if (err != nil) {
    fmt.Println(err)
  }

  return &MiscConfig{
    SaveDir: savedir.String(),
    CashRegister: cashregister.String(),
  }
}
