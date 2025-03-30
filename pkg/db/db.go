package db

import "database/sql"

var DB *sql.DB

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT 0
    );`

	_, err := DB
}
