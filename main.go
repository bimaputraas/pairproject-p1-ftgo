package main

import (
	"github.com/bimaxputra/pairproject-p1-ftgo/cli"
	"github.com/bimaxputra/pairproject-p1-ftgo/config"
	"github.com/bimaxputra/pairproject-p1-ftgo/handler"
)

func main() {
	db := config.ConnectDb()
	defer db.Close()

	userHandler := handler.Handler{UserHandler: db}
	
	app := cli.Cli{Handler: &userHandler}
	app.MainGateInterface()
}