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

func NewPostgresStorage(host, port, user, password, dbname string) (*Postgres, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
	if err := p.createAccountTable(); err != nil {
		return err
	}
	if err := p.createTrasferHistoryTable(); err != nil {
		return err
	}
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

func (p *Postgres) createAccountTable() error {
	q := `create table if not exists account (
		id serial primary key,
		login varchar(50),
		password varchar(120),
		card_number serial,
		balance serial,
		created_at timestamp
		)`

	return p.checkExecError(q)
}

func (p *Postgres) createTrasferHistoryTable() error {
	q := `create table if not exists transfer_history (
		id serial primary key,
		from_card_Number serial,
		to_card_Number serial,
		message varchar(240),
		date timestamp
		)`
	return p.checkExecError(q)
}

func (p *Postgres) checkExecError(query string) error {
	_, err := p.db.Exec(query)
	return err
}
