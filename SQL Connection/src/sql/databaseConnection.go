package main

import (
	"database/sql"
)

func SqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	return db
}
