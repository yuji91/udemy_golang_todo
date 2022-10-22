package main

import (
  "fmt"
  "./config" // import inner package
)

func main() {
  fmt.Println(config.Config.Port)
  fmt.Println(config.Config.SQLDriver)
  fmt.Println(config.Config.DbName)
  fmt.Println(config.Config.LogFile)
}
