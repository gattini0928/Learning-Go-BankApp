package methods

import (
	"database/sql"
	"bankapp/models"
	"log"	
	_ "modernc.org/sqlite"
)

func InsertUser(db *sql.DB, u models.User) {
	_, err := db.Exec(
		`INSERT INTO users (name, email, cpf, password)
		VALUES (?, ?, ?, ?)`, u.Name, u.Email, u.Cpf, u.Password)
	if err != nil {
		log.Fatal(err)
	}
}