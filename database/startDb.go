package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"

)

func StartDB() *sql.DB {
	db, err := sql.Open("sqlite", "bank.db")
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

		CREATE TABLE IF NOT EXISTS accounts (
            id          INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id     INTEGER,
			account_password    TEXT,
            level       TEXT,
            balance     REAL,
            card_name   TEXT,
            card_number TEXT,
            card_cvv    TEXT,
            card_expiry TEXT,
            card_limit  REAL,
			invoice_due_date TEXT,
			invoice_total REAL,
			invoice_paid INTEGER NOT NULL DEFAULT 0,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );

		CREATE TABLE IF NOT EXISTS bank_transactions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			account_id INTEGER,
			from_user INTEGER,
			to_user INTEGER,
			amount REAL,
			date TEXT,
			FOREIGN KEY (account_id) REFERENCES accounts(id)
		);

		CREATE TABLE IF NOT EXISTS credit_transactions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			account_id INTEGER,
			from_user INTEGER,
			to TEXT,
			amount REAL,
			date TEXT,
			FOREIGN KEY (account_id) REFERENCES accounts(id)
		);
	`)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Database de bank criado com sucesso!")
	return db
}

