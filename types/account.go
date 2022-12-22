package types

import (
	"math/rand"
	"time"
)

type AccountReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Account struct {
	ID         string `json:"ID"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Balance    int    `json:"balance"`
	CardNumber int    `json:"card_number"`
	CreatedAt  string `json:"created_at"`
}

func NewAccount(login, password string) *Account {
	return &Account{
		Login:      login,
		Password:   password,
		Balance:    rand.Intn(1000),
		CardNumber: rand.Intn(100000),
		CreatedAt:  time.Now().UTC().Format("_2/Jan/2006"),
	}
}
