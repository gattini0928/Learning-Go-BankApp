package app

import (
	"bankapp/authentication"
	"bankapp/database"
	"bankapp/methods"
	"log"
	"bankapp/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"database/sql"
	_ "modernc.org/sqlite"
)

var itemsToPurchase = map[string]float64 {
	"Banana": 7.00,
	"Book - O Alquimista": 50.0,
	"Mouse": 199.99,
	"Keyboard": 400.00,
	"TV": 750.0,
	"Go Course for Sênior's Developers": 2000,
	"PC": 10000.00,
}

func App() {
	reader := bufio.NewReader(os.Stdin)
	db := database.StartDB()
	defer db.Close()
	logged := false
	for {
		fmt.Println("💰 Bem-vindo ao GoBank 💰")
		fmt.Println("1 - Criar usuário 😉➕")
		fmt.Println("2 - Fazer Login 🚪")
		fmt.Println("3 - Sair 🚪➡️")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Opção Inválida. Tente Novamente")
			continue
		}

		if choice == 1 {
			fmt.Print("🪪 Nome Completo: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("📩 Seu E-mail: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("🔢 Seu CPF: ")
			cpf, _ := reader.ReadString('\n')
			cpf = strings.TrimSpace(cpf)

			fmt.Print("🔑 Senha: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			fmt.Print("🔑 Confirme sua senha: ")
			confirmPassword, _ := reader.ReadString('\n')
			confirmPassword = strings.TrimSpace(confirmPassword)

			if confirmPassword != password {
				fmt.Println("❌ As senhas não coincidem")
				continue
			}

			user := models.User{
				Name:     name,
				Email:    email,
				Cpf:      cpf,
				Password: password}

			account := models.NewAccount(user)

			signup := authentication.SignUp(db, user)
			if !signup {
				fmt.Println("❌ Falha ao criar usuário")
				continue
			}

			fmt.Println("🔢🔑 Agora crie sua senha bancária, ela será usada para confirmar suas ações no GoBank")
			fmt.Println("‼️ Sua senha deve conter 6 DÍGITOS!")
			fmt.Println("🔑 Sua senha(ex: 924501): ")
			accountPassword, _ := reader.ReadString('\n')
			accountPassword = strings.TrimSpace(accountPassword)
			
			p, ok := authentication.ValidatePassword(accountPassword)
			if !ok {
				continue
			}
			account.AccountPassword = p

			methods.InsertAccount(db, account)
			fmt.Printf("😎 Usuário %s criado com sucesso, seja bem-vindo ao GoBank \n", user.Email)
			logged = true
			bankManager(reader, db, account, logged)

		} else if choice == 2 {
			fmt.Print("📩 Seu E-mail: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("🔑 Senha: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			user := models.User {
				Email:    email,
				Password: password,
			}

			login := authentication.Login(db, user)
			if !login {
				fmt.Println("❌ Falha no login, tente novamente")
				continue
			}

			account, err := methods.GetAccount(db, user.Email)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("😎 Bem-vindo novamente ao GoBank %s \n", user.Email)
			logged = true
			bankManager(reader, db, account, logged)

		}  else if choice == 3 {
			fmt.Println("👋 Até mais, obrigado por usar o GoBank")
			break
		} else {
			fmt.Println("❌ Opção Inválida, por favor tente novamente")
			continue
		}
	}
}

func bankManager(reader *bufio.Reader , db *sql.DB, account models.Account, logged bool) {
	if logged {
		for {
			fmt.Println("1 - Transferir 🪪")
			fmt.Println("2 - Consultar Saldo 🔍")
			fmt.Println("3 - Visualizar Transações 🗺️")
			fmt.Println("4 - Visualizar Faturas 🗺️")
			fmt.Println("5 - Pagar Fatura 💳")
			fmt.Println("6 - Realizar Compra 🛍️")
			fmt.Println("7 - Sair ➡️🚪")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			choice, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("❌ Opção Inválida. Tente Novamente")
				continue
			}
			if choice == 1 {
				fmt.Println("📩 Digite o email do destinatário: ")
				email, _ := reader.ReadString('\n')
				email = strings.TrimSpace(email)

				fmt.Println("💵 Digite o valor da transferência: ")
				amount, _ := reader.ReadString('\n')
				amount = strings.TrimSpace(amount)

				fmt.Println("Digite sua senha bancária para confirmar a transferência: ")
				a_password, _ := reader.ReadString('\n')
				a_password = strings.TrimSpace(a_password) 

				msg, ok := methods.NewTransfer(db, account, email, amount, a_password)
				if !ok {
					fmt.Println(msg)
					continue
				}
				fmt.Println(msg)
			} else if choice == 2 {
				fmt.Println(methods.CheckBalance(db, account.AccountUser.Email))
			}else if choice == 3 {
				methods.DisplayTransactions(db, account.AccountUser.Email)
			} else if choice == 4 {
				methods.CheckCreditInvoices(db, account.AccountUser.Email)
			} else if choice == 5 {
				fmt.Println("Digite sua senha bancária para confirmar o pagamento da sua fatura: ")
				a_password, _ := reader.ReadString('\n')
				a_password = strings.TrimSpace(a_password) 
				methods.PayInvoice(db, account.AccountUser.Email, a_password)
			
			}else if choice == 6 {


			} else if choice == 7 {
				fmt.Println("👋 Até mais, obrigado por usar o GoBank")
				logged = false
				break
			} else {
				fmt.Println("❌ Opção Inválida, por favor tente novamente")
				continue
			}
		}
	}
}
