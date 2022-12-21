package storage

import "github.com/arimatakao/bankapi/types"

type DBForTests struct {
	//data map[string][]string
}

func NewDBForTests() *DBForTests {
	return nil
	// return &DBForTests{}
}

func (d *DBForTests) CreateAccount(a *types.Account) error {
	return nil
}

func (d *DBForTests) GetAcocuntByID(id string) (*types.Account, error) {
	return nil, nil
}

func (d *DBForTests) UpdateAccount(a *types.Account) error {
	return nil
}

func (d *DBForTests) DeleteAccount(id string) error {
	return nil
}
