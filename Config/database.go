package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:lTTbg73BVYWIzJDe1gvc@tcp(containers-us-west-126.railway.app:7015)/railway")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
