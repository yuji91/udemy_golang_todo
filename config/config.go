package config

type ConfigList struct {
  Port      string
  SQLDriver string
  DbName    string
  Logfile   string
}

var Config ConfigList // global variable to access
