package main

import (
	"m_p1/cli"
	"m_p1/config"
	"m_p1/handler"
)

func main() {
	db := config.ConnectDb()
	defer db.Close()

	userHandler := handler.Handler{UserHandler: db}
	
	app := cli.Cli{Handler: &userHandler}
	app.MainGateInterface()
}