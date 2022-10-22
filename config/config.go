package config

import (
  "log"
  "../utils"

  "gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
  Port      int
  SQLDriver string
  DbName    string
  LogFile   string
}

var Config ConfigList // global variable to access

func init() {
  LoadConfig()
  utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
  cfg, err := ini.Load("config.ini")
  if err != nil {
    log.Fatalln(err)
  }
  Config = ConfigList{
    Port:      cfg.Section("web").Key("port").MustInt(8080),
    DbName:    cfg.Section("db").Key("name").String(),
    SQLDriver: cfg.Section("db").Key("driver").String(),
    LogFile:   cfg.Section("web").Key("logfile").String(),
  }
}
