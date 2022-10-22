package main

import (
  "fmt"
  "./config" // import inner package
  "log"
)

func main() {
  fmt.Println(config.Config.Port)
  fmt.Println(config.Config.SQLDriver)
  fmt.Println(config.Config.DbName)
  fmt.Println(config.Config.LogFile)

  log.Println("test")
}
