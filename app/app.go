package app

import (
	"bankapp/authentication"
	"bankapp/database"
	"bankapp/methods"
	"bankapp/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func App() {
	reader := bufio.NewReader(os.Stdin)
	db := database.StartDB()
	defer db.Close()
	logged := false
	for {
		fmt.Println("Bem-vindo ao GoBank")
		fmt.Println("1 - Criar usuário")
		fmt.Println("2 - Fazer Login")
		fmt.Println("3 - Sair")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Opção Inválida. Tente Novamente")
			continue
		}

		if choice == 1 {
			fmt.Print("Nome Completo: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Seu E-mail: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Seu CPF: ")
			cpf, _ := reader.ReadString('\n')
			cpf = strings.TrimSpace(cpf)

			fmt.Print("Senha: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			fmt.Print("Confirme sua senha: ")
			confirmPassword, _ := reader.ReadString('\n')
			confirmPassword = strings.TrimSpace(confirmPassword)

			if confirmPassword != password {
				fmt.Println("As senhas não coincidem")
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
				fmt.Println("Falha ao criar usuário")
				continue
			}

			fmt.Println("Agora crie sua senha bancária, ela será usada para confirmar suas ações no GoBank")
			fmt.Println("Sua senha deve conter 6 DÍGITOS!")
			fmt.Println("Sua senha(ex: 924501): ")
			accountPassword, _ := reader.ReadString('\n')
			accountPassword = strings.TrimSpace(accountPassword)
			
			p, ok := authentication.ValidatePassword(accountPassword)
			if !ok {
				fmt.Println("Senha bancária inválida")
				continue
			}
			account.AccountPassword = p

			methods.InsertAccount(db, account)
			fmt.Printf("Usuário %s criado com sucesso, seja bem-vindo ao GoBank \n", user.Email)
			logged = true
			bankManager(reader, logged)

		} else if choice == 2 {
			fmt.Print("Seu E-mail: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Senha: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			user := models.User {
				Email:    email,
				Password: password,
			}

			login := authentication.Login(db, user)
			if !login {
				fmt.Println("Falha no login, tente novamente")
				continue
			}

			fmt.Printf("Bem-vindo novamente ao GoBank %s \n", user.Email)
			logged = true
			bankManager(reader, logged)

		}  else if choice == 3 {
			fmt.Println("Até mais, obrigado por usar o GoBank")
			break
		} else {
			fmt.Println("Opção Inválida, por favor tente novamente")
			continue
		}
	}
}

func bankManager(reader *bufio.Reader , logged bool) {
	if logged {
		for {
			fmt.Println("1 - Transferir")
			fmt.Println("2 - Consultar Saldo")
			fmt.Println("3 - Ver Faturas")
			fmt.Println("4 - Pagar Fatura")
			fmt.Println("5 - Sair")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			choice, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Opção Inválida. Tente Novamente")
				continue
			}
			if choice == 1 {
				methods.NewTransfer()
			} else if choice == 2 {
				methods.CheckBalance()
			} else if choice == 3 {
				methods.CheckInvoices()
			} else if choice == 4 {
				methods.PayInvoice()
			} else if choice == 5 {
				fmt.Println("Até mais, obrigado por usar o GoBank")
				break
			} else {
				fmt.Println("Opção Inválida, por favor tente novamente")
				continue
			}
		}
	}
}
