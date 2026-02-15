package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"

)

func StartDB() *sql.DB {
	db, err := sql.Open("sqlite", "accounts.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			cpf TEXT,
			password TEXT);
	`)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Database de accounts criado com sucesso!")
	return db
}

