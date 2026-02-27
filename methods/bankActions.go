package methods

import (
	"bankapp/authentication"
	"bankapp/models"
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
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

func MakePurchase(db *sql.DB, reader *bufio.Reader ,email string, productTitle string,productPrice float64) (string, bool) {
	account, err := GetAccount(db, email)
	if err != nil {
		return "Erro ao acessar conta", false
	}

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice := strings.ToLower(input)
	if choice == "d" {
		if account.Balance < productPrice {
			return "Você não possui saldo o suficiente para realizar esta compra",false
		}
		_, err = db.Exec(`UPDATE accounts SET balance = balance - ? WHERE user_id = (SELECT id FROM users WHERE email = ?)`, productPrice, account.AccountUser.Email)
		if err != nil {
			return "Erro ao realizar pagamento", false
		}
		msg := fmt.Sprintf("Compra: %s no valor de: R$%.2f processada com sucesso", productTitle, productPrice)
		return msg, true
	} else if choice == "c" {
		// NAO SETAR BALANCE, MAS SETAR ->
		fmt.Println("Digite os dados do seu cartão de crédito")
		// Verificar se dados batem primeiro, success prosseguir
		// Verificar se compra é maior que o limit , se sim -> Barrar
		// Senão -> Pedir infos de cartao de crédito(se corretas prosseguir)
		// limit -= product.Price
		// Success -> Options 12x valor divido por numsInstallment 
		// From -> user
		// To -> "GoApp"
		// Amount -> productPrice
		// Date -> time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
		// Enviar setando: Installments := valor divido por quantidade de parcelas
		// transaction := models.CreditCardTransaction(From, To, Amount, Date, Installment)
		// accounts.Invoice.Transaction.append(transaction)
		// Total += productPrice
		// Paid = false

	} else {
		return "Opção inválida, tente novamente", false
	}

	fmt.Println("D para Débito, C para crédito:")
	



}

// Criar um produto aleatório struct {}
// string -> 




