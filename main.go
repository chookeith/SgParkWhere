package main

import (
  "SgParkWhere/router"
)

func main(){
  // Set the router as the default one shipped with Gin
  var connection router.Export
  connection.Init()

}
