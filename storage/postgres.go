package storage

import (
	"database/sql"
	"fmt"

	"github.com/arimatakao/bankapi/types"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgresStorage(user, dbname, password string) (*Postgres, error) {
	conn := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) Init() error {
	return nil
}

func (p *Postgres) CreateAccount(a *types.Account) error {
	return nil
}

func (p *Postgres) GetAcocuntByID(id string) (*types.Account, error) {
	return nil, nil
}

func (p *Postgres) UpdateAccount(a *types.Account) error {
	return nil
}

func (p *Postgres) DeleteAccount(id string) error {
	return nil
}
