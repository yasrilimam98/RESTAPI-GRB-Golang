package main

import "github/yasrilimam98/grb-restapi/routes"



func main() {
	e:= routes.Init()
	e.Logger.Fatal(e.Start(":5000"))
}
