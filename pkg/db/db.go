package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/todos")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Createtable() {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT 0
    );`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
