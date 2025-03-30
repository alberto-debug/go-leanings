package db

import "database/sql"

var DB *sql.DB

func initDB() {
	var err error

	DB, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/todos")
}
