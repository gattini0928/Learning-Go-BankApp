package app

import (
	"bankapp/models"
	"bankapp/databases"
	"bankapp/authentication"
	"bankapp/methods"
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func App() {
	reader := bufio.NewReader(os.Stdin)
	db := databases.ManagerDB()
	defer db.Close()
	for {
		fmt.Println("Bem-vindo ao GoBank")
		fmt.Println("1 - Criar usuário")

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
			Name: name, 
			Email:email, 
			Cpf:cpf, 
			Password: password}

		signup := authentication.SignUp(db,  user)
		if !signup {
			fmt.Println("Falha ao criar usuário")
			continue
		}
		methods.InsertUser(db, user)
		fmt.Printf("Usuário %s criado com sucesso, seja bem-vindo ao GoBank \n", user.Email)
		}
	}
}