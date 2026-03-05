package models

import (
	"bankapp/generators"
	"time"
)

type Product struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
}

type AccountLevel string

const (
	Bronze AccountLevel = "Bronze"
	Silver AccountLevel = "Silver"
	Gold AccountLevel = "Gold"
	Platinum AccountLevel = "Platinum"
	Diamond AccountLevel = "Diamond"
	Premium AccountLevel = "Premium"
)

var AccountPrices = map[AccountLevel]int{
	Bronze:   0,
	Silver:   30,
	Gold:     50,
	Platinum: 70,
	Diamond:  120,
	Premium:  200,
}

type User struct {
	Id int
	Name string
	Email string
	Cpf string
	Password string
}

type CreditCard struct {
	HolderName string
	CardNumber string
	Cvv string
	ExpiryDate string
	Limit float64
}

type NormalTransaction struct {
	From int
	To int
	ToName string
	Amount float64
	Date time.Time
}

type BankTransactions struct {
	Transactions []NormalTransaction
}

type CreditCardTransaction struct {
    To     string
    Amount float64
    Date  string
	Installments int
	RemainingInstallments int
}

type Invoice struct {
	Transactions []CreditCardTransaction
	Total float64
	DueDate string
	Paid bool
}

type Account struct {
	Id int
	AccountUser User
	AccountPassword string
	Level AccountLevel
	CreditCard CreditCard
	Transactions BankTransactions
	Invoice Invoice
	Balance float64
}

func NewAccount(user User) Account {
	name, number, cvv, date := generators.GenerateCreditCard(user.Name)
	creditCard := CreditCard{HolderName:name, 
		CardNumber: 
		number, 
		Cvv: cvv, 
		ExpiryDate: date,
		Limit: 400.00,
	}
	return Account {
		AccountUser: user,
		AccountPassword: "",
		Level: Bronze,
		CreditCard: creditCard,
		Transactions: BankTransactions{
			Transactions: []NormalTransaction{},
		},
		Invoice: Invoice {
			Transactions: []CreditCardTransaction{},
            Total:        0,
            DueDate:      time.Now().AddDate(0, 1, 0).Format("2006-01-02"),
			Paid: false,
		},
		Balance: 200,
	}
}