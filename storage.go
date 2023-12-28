package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)




type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	ClearAccounts()error
	UpdateAccount(*Account) error
	GetAccounts()([]*Account , error)
	GetAccountByID(int)(*Account , error)
}

type PostgressStore struct {
	db *sql.DB
}

func (s *PostgressStore)ClearAccounts() error {
	query := `delete from account`
	_,err := s.db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func NewPostgressStore()(*PostgressStore , error) {
	connStr := "user=postgres dbname=postgres password=fd472992   sslmode=disable"
	db, err := sql.Open("postgres" , connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgressStore{
		db : db,
	}, nil
}


func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp

	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) GetAccounts()([]*Account , error){
	query := `select * from account Limit 10`
	rows , err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccounts(rows)
		if err != nil {
			return nil , err
		}
		accounts = append(accounts , account)
	}
	return accounts , nil
}

func scanIntoAccounts(rows *sql.Rows)(*Account , error){
	account := new(Account)
	err := rows.Scan(
		&account.ID , 
		&account.FirstName , 
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)
	return account , err
}



func (s *PostgressStore) CreateAccount(acc *Account) error{
	query := `insert into account
	(first_name , last_name , number , balance , created_at)
	values ($1 , $2 , $3 , $4 , $5)`
	resp, err := s.db.Query(query , acc.FirstName , acc.LastName , acc.Number , acc.Balance , acc.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Println("%+v\n" , resp)
	return nil
}
func (s *PostgressStore) UpdateAccount(*Account) error{
	return nil
}
func (s *PostgressStore) DeleteAccount(id int) error{
	_ , err:= s.db.Query("delete from account where id = $1" , id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostgressStore) GetAccountByID(id int) (*Account , error){
	rows , err := s.db.Query("select * from account where id = $1" , id)
	if err != nil {
		return nil , err 
	}

	for rows.Next() {
		return scanIntoAccounts(rows)
	}

	return nil , fmt.Errorf("accout %d not found" , id)
}