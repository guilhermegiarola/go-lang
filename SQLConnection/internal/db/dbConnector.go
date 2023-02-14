package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func SqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return db
}
