package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const(
	dbUser = "root"
	dbPassword = ""
	dbHost = "localhost:3306"
	dbName = "milestone_p1_test"
)

func ConnectDb() *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser,dbPassword,dbHost,dbName)

	db,err := sql.Open("mysql",dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}