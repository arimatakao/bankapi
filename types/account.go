package types

import (
	"math/rand"
	"strconv"
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
	CardNumber int    `json:"cardNumber"`
	CreatedAt  string `json:"createdAt"`
}

func NewAccount(login, password string) *Account {
	return &Account{
		ID:         strconv.Itoa(rand.Intn(100)),
		Login:      login,
		Password:   password,
		Balance:    rand.Intn(1000),
		CardNumber: rand.Intn(100000),
		CreatedAt:  "2011",
	}
}
