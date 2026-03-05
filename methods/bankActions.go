package methods

import (
	"bankapp/authentication"
	"bankapp/models"
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
	_ "modernc.org/sqlite"
)

func NewTransfer(db *sql.DB, account models.Account, email string, amount string, a_password string) (string, bool) {
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
		fmt.Println("ERRO GetAccount", err)
		return "Erro ao buscar destinatário", false
	}

	if amountFloat <= 0 {
		return "O valor fornecido precisa ser maior do que 0", false
	}

	if account.Balance <= amountFloat {
		return "Seu saldo é menor do que o valor fornecido", false
	}

	if a_password != account.AccountPassword {
		return "Senha incorreta, tente novamente.", false
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
		fmt.Println("ERRO GetAccount", err)
		return "Erro ao buscar saldo"
	}
	return fmt.Sprintf("Seu saldo atual é de R$%.2f", account.Balance)
}

func DisplayTransactions(db  *sql.DB, email string){
	account, err := GetAccount(db, email)
	if err != nil {
		fmt.Println("ERRO GetAccount", err)
		fmt.Println("Erro ao buscar transações")
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

func CheckCreditInvoices(db *sql.DB, email string) {
	account, err := GetAccount(db, email)

	if err != nil {
		fmt.Println("ERRO GetAccount", err)
		fmt.Println("Erro ao buscar fatura")
	}

	invoice := account.Invoice 
	fmt.Printf("Data de Vencimento: %v - Total: R$%.2f", invoice.DueDate, invoice.Total)
}

func PayInvoice(db *sql.DB, email string, a_password string) {
	account, err := GetAccount(db, email)

	if err != nil {
		fmt.Println("ERRO GetAccount", err)
		fmt.Println("Erro ao buscar fatura")
		return
	}

	if account.Balance < account.Invoice.Total {
		fmt.Println("Você não possui saldo o suficiente para pagar esta fatura.")
		return
	}

	if a_password != account.AccountPassword {
		fmt.Println("Senha incorreta, tente novamente.")
		return
	}

	success := PaidInvoice(db, account)
	if !success {
		fmt.Println("Falha ao pagar fatura, tente novamente mais tarde.")
		return
	}

	fmt.Printf("Fatura com valor de: R$%.2f paga com sucesso!", account.Invoice.Total)
}

func MakePurchase(db *sql.DB, reader *bufio.Reader ,email string, productTitle string, productPrice float64) (string, bool) {
	account, err := GetAccount(db, email)
	if err != nil {
		return "Erro ao acessar conta", false
	}

	fmt.Println("D para Débito, C para crédito:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice := strings.ToLower(input)
	switch choice {
		case "d":
			if account.Balance < productPrice {
				return "Você não possui saldo o suficiente para realizar esta compra",false
			}
			_, err = db.Exec(`UPDATE accounts SET balance = balance - ? WHERE user_id = (SELECT id FROM users WHERE email = ?)`, productPrice, account.AccountUser.Email)
			if err != nil {
				return "Erro ao realizar pagamento", false
			}
			msg := fmt.Sprintf("Compra: %s no valor de: R$%.2f processada com sucesso", productTitle, productPrice)
			return msg, true
		case "c":
			fmt.Println("Digite os dados do seu cartão de crédito: ")
			fmt.Println("Nome no cartão: ")
			cardName, _ := reader.ReadString('\n')
			cardName = strings.TrimSpace(cardName)
			
			fmt.Println("Número no cartão: ")
			cardNumber, _ := reader.ReadString('\n')
			cardNumber = strings.TrimSpace(cardNumber)
			
			fmt.Println("Data de expiração")
			expiryDate, _ := reader.ReadString('\n')
			expiryDate = strings.TrimSpace(expiryDate)

			fmt.Println("Cvv: ")
			cvv, _ := reader.ReadString('\n')
			cvv = strings.TrimSpace(cvv)

			fmt.Println("Em quantas vezes você deseja realizar esta compra? Max(12x): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			installments, err := strconv.Atoi(input)
			if err != nil {
					return "Erro ao escolher parcelas", false
				}

			if cardName != account.CreditCard.HolderName || cardNumber != account.CreditCard.CardNumber || 
			expiryDate != account.CreditCard.ExpiryDate ||
			cvv != account.CreditCard.Cvv{
				return "Dados do cartão inválidos", false
			}
			if installments > 12  || installments <= 0{
				return "Número de parcelas inválidas (1x-12x)", false
			}

			expiry, err := time.Parse("2006-01-02", account.CreditCard.ExpiryDate)
			if err != nil {
				return "Erro ao processar data do cartão de crédito", false
			}

			if time.Now().After(expiry) {
				return "A data de validade do seu cartão venceu", false
			}

			if productPrice > account.CreditCard.Limit {
				return "O valor da compra excede o limite do seu cartão", false
			}

			_, err = db.Exec(`UPDATE accounts SET card_limit = card_limit - ? WHERE user_id = (SELECT id FROM users WHERE email = ?)`, productPrice, account.AccountUser.Email)
			if err != nil {
				return "Erro ao realizar pagamento", false
			}

			installmentPrice := productPrice / float64(installments)
			_, err = db.Exec(`UPDATE accounts SET invoice_total = invoice_total + ? WHERE user_id = (SELECT id FROM users WHERE email = ?)`, installmentPrice, account.AccountUser.Email)
			if err != nil {
				return "Erro ao realizar pagamento", false
			}

			to := "GoApp"
			amount := productPrice
			date := time.Now().AddDate(0, 1, 0).Format("2006-01-02")
			
			_, err = db.Exec(`
				INSERT INTO credit_transactions 
				(account_id,to, amount, date, installments)
				VALUES(?,?,?,?,?)
			`, account.Id, to, amount, date, installments)
			if err != nil {
				return "Erro ao realizar pagamento", false
			}
			msg := fmt.Sprintf("Compra de %s, no valor de: R$%.2f realizada com sucesso", productTitle, productPrice)
			return msg, true
		default:
			return "Opção inválida, tente novamente", false
	}
}

func UpdateAccountLevel(db *sql.DB, email string) string {
	account, err := GetAccount(db, email)
	if err != nil {
		return "Erro ao acessar conta"
	}
	var level string
	var price int
	var msg string

	if account.Balance >= 50000 {
		level = "Premium"
		price = 200
		msg = "Parabéns sua conta tem status Premium"
	} else if account.Balance >= 20000 {
		level = "Diamond"
		price = 120
		msg = "Parabéns sua conta tem status Diamante"
	}  else if account.Balance >= 15000 {
		level = "Platinum"
		price = 70
		msg = "Parabéns sua conta tem status Platina"
	} else if account.Balance >= 10000 {
		level = "Gold"
		price = 50
		msg = "Parabéns sua conta tem status Ouro"
	} else if account.Balance >= 5000 {
			level = "Silver"
			price = 30
			msg = "Parabéns sua conta tem status Prata"
	} else if account.Balance < 5000 {
		level = "Bronze"
		price = 0
		msg = "Sua conta tem status Bronze, acumule saldo para subir de nível"
	}
	_, err = db.Exec(`
	UPDATE accounts 
	SET level = ?, invoice_total = invoice_total - ?
	WHERE user_id = (SELECT id FROM users WHERE email = ?)`,
	level, price, email)

	if err != nil {
		return "Erro ao atualizar nível da conta"
	}
	return msg
}





