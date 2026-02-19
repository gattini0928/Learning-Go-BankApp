package methods

import (
	"database/sql"
	"bankapp/models"
	"log"	
	_ "modernc.org/sqlite"
)

func InsertUser(db *sql.DB, u models.User) int64 {
	result, err := db.Exec(
		`INSERT INTO users (name, email, cpf, password)
		VALUES (?, ?, ?, ?)`, u.Name, u.Email, u.Cpf, u.Password)

	id, err := result.LastInsertId() 
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func InsertAccount(db *sql.DB, a models.Account) {
    userId := InsertUser(db, a.AccountUser)

    _, err := db.Exec(
        `INSERT INTO accounts (user_id, level, balance, card_name, card_number,
        card_cvv, card_expiry, card_limit, invoice_due_date, invoice_total)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        userId, a.Level, a.Balance, a.CreditCard.HolderName,
        a.CreditCard.CardNumber, a.CreditCard.Cvv, a.CreditCard.ExpiryDate,
        a.CreditCard.Limit, a.Invoice.DueDate, a.Invoice.Total,
    )
    if err != nil {
        log.Fatal(err)
    }
}