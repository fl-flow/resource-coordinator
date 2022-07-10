package main

import (
  _ "github.com/fl-flow/resource-coordinator/etc"
  "github.com/fl-flow/resource-coordinator/http_server"
)


func main() {
  httpserver.Run()
}
