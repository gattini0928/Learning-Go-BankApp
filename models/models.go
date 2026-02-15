package models

type AccountLevel string

const (
	Bronze AccountLevel = "Bronze"
	Silver AccountLevel = "Silver"
	Gold AccountLevel = "Gold"
	Platinum AccountLevel = "Platinum"
	Diamond AccountLevel = "Diamond"
	Premium AccountLevel = "Premium"
)

var AccountPrices = map[AccountLevel]int {
	"Bronze": 0,
	"Silver": 30,
	"Gold": 50,
	"Platinum": 70,
	"Diamon": 120,
	"Premium": 200,
}

type User struct {
	id int
	Name string
	Email string
	Cpf string
	Password string
}

type Account struct {
	AccountUser User
	Level AccountLevel
	CreditCard string
}