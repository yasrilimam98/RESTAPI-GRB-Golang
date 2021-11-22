package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/yasrilimam98/grb-restapi/database"
	"github.com/yasrilimam98/grb-restapi/routes"
)

func main() {

	database.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":5000"))
}
