package methods

import (
	"bankapp/models"
	"database/sql"
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
        `INSERT INTO accounts (user_id, account_password, level, balance, card_name, card_number,
        card_cvv, card_expiry, card_limit, invoice_due_date, invoice_total)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        userId, a.AccountPassword, a.Level, a.Balance, a.CreditCard.HolderName,
        a.CreditCard.CardNumber, a.CreditCard.Cvv, a.CreditCard.ExpiryDate,
        a.CreditCard.Limit, a.Invoice.DueDate, a.Invoice.Total,
    )
    if err != nil {
        log.Fatal(err)
    }
}

func GetAccount(db *sql.DB, email string) (models.Account, error) {
    query := `SELECT a.balance, a.level, a.card_number, a.card_name, 
              a.card_cvv, a.card_expiry, a.card_limit, 
              a.invoice_due_date, a.invoice_total, a.account_password
              FROM accounts a
              JOIN users u ON u.id = a.user_id
              WHERE u.email = ?`

    var a models.Account
    err := db.QueryRow(query, email).Scan(
        &a.Balance, &a.Level, &a.CreditCard.CardNumber, &a.CreditCard.HolderName,
        &a.CreditCard.Cvv, &a.CreditCard.ExpiryDate, &a.CreditCard.Limit,
        &a.Invoice.DueDate, &a.Invoice.Total, &a.AccountPassword,
    )
    if err != nil {
        return models.Account{}, err
    }
    a.AccountUser.Email = email
    return a, nil
}

func GetEmailById(db *sql.DB, id int) string {
    var email string
    err := db.QueryRow(`SELECT email FROM users WHERE id = ?`, id).Scan(&email)
    if err != nil {
        return "desconhecido"
    }
    return email
}

func PaidInvoice(db *sql.DB, account models.Account) bool {
    tx, err := db.Begin()
    if err != nil {
        return false
    }

    _, err = tx.Exec(`
        UPDATE accounts
        SET 
            invoice_paid = 1,
            balance = balance - ?
        WHERE user_id = ?
    `, account.Invoice.Total, account.AccountUser.Id)

    if err != nil {
        tx.Rollback()
        return false
    }

    return tx.Commit() == nil
}
