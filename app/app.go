package app

import "fmt"

type User struct {
	id int
	Name string
	Email string
	Cpf string
	Password string
}

type Account struct {
	AccountUser User
	CreditCard string
}

func app() {
	fmt.Printf("All things")
}