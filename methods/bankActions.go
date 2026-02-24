package methods

import (
	"bankapp/authentication"
	"bankapp/models"
	"database/sql"
	"fmt"
	"strconv"
	_ "modernc.org/sqlite"
)

func NewTransfer(db *sql.DB, account models.Account, email string, amount string) (string, bool) {
	uExists, err := authentication.UserExists(db, email)
	if err != nil {
		return "Erro ao encontrar email", false
	}

	if !uExists {
		msg := fmt.Sprintf("O email %s não existe, tente novamente.", email)
		return msg, false
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	
	if err != nil {
    	return "Valor inválido", false
	}

	userRecipient, err := GetAccount(db, email)

	if err != nil {
		return "Erro ao buscar destinatário", false
	}

	if amountFloat <= 0 {
		return "O valor fornecido precisa ser maior do que 0", false
	}

	if account.Balance <= amountFloat {
		return "Seu saldo é menor do que o valor fornecido", false
	}

	_, err = db.Exec(`UPDATE accounts SET balance = balance - ? WHERE user_id = (SELECT id FROM users WHERE email = ?)`, amountFloat, account.AccountUser.Email)
	if err != nil {
		return "Erro na transferência", false
	}

	_, err = db.Exec(`UPDATE accounts SET balance = balance + ? WHERE user_id = (SELECT id FROM users WHERE email = ?)`, amountFloat, email)
	if err != nil {
		return "Erro na transferência", false
	}

	msg := fmt.Sprintf("Transferência de R$%.2f para %s concluída com sucesso", amountFloat, userRecipient.AccountUser.Email)
	return msg, true
}

func CheckBalance(db *sql.DB, email string) string {
	account, err := GetAccount(db, email)
	if err != nil {
		return "Erro ao buscar saldo"
	}
	return fmt.Sprintf("Seu saldo atual é de R$%.2f", account.Balance)
}

func DisplayTransactions(db  *sql.DB, email string){
	account, err := GetAccount(db, email)
	if err != nil {
		fmt.Println("Erro ao buscar faturas")
		return
	}

	transactions := account.Transactions.Transactions
	if len(transactions) == 0 {
		fmt.Println("Você não possui nenhuma transação bancária até o momento")
		return
	}
	
	for _, t := range transactions {
		from := GetEmailById(db, t.From)
		to := GetEmailById(db, t.To)
		fmt.Printf("Data: %v - Valor: R$%.2f - De: %s - Para: %s\n", t.Date, t.Amount, from, to)
	}
}

