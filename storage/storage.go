package storage

import (
	"github.com/arimatakao/bankapi/types"
)

type Storager interface {
	CreateAccount(a *types.Account) error
	GetAcocuntByID(id string) (*types.Account, error)
	UpdateAccount(a *types.Account) error
	DeleteAccount(id string) error
}
