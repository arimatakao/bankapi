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
	q := fmt.Sprintf(`insert into account (login, password, card_number, balance, created_at)
	values (
		'%s', '%s', %v, %v, '%s'
		)`,
		a.Login, a.Password, a.CardNumber, a.Balance, a.CreatedAt)
	return p.checkExecError(q)
}

func (p *Postgres) GetAcocuntByID(id string) (*types.Account, error) {
	return nil, nil
}

func (p *Postgres) UpdateAccount(a *types.Account) error {
	return nil
}

func (p *Postgres) DeleteAccount(id string) error {
	q := fmt.Sprintf(`delete from account where id = %v`, id)
	return p.checkExecError(q)
}

func (p *Postgres) createAccountTable() error {
	q := `create table if not exists account (
		Id serial not null,
		Login varchar(50) not null,
		Password varchar(120) not null,
		Card_number serial not null,
		Balance serial not null,
		Created_at date not null,
		PRIMARY KEY(Id)
		)`

	return p.checkExecError(q)
}

func (p *Postgres) createTrasferHistoryTable() error {
	q := `create table if not exists transfer_history (
		Id serial not null,
		From_card_Number serial not null,
		To_card_Number serial not null,
		Message varchar(240),
		Date date not null,
		PRIMARY KEY(Id)
		)`
	return p.checkExecError(q)
}

func (p *Postgres) checkExecError(query string) error {
	_, err := p.db.Exec(query)
	return err
}
