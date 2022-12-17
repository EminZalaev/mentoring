package internal

import (
	"database/sql"
	"fmt"
)

type Store interface {
	GetCurrency() (*[]currency, error)
	PostCurrency(crc *currency) error
	PutCurrency(crc *currency) error
}

type Storage struct {
	DB *sql.DB
}

func NewStorage(cfg *Config) (*Storage, error) {
	db, err := initStorage(cfg)
	if err != nil {
		return nil, err
	}
	return &Storage{
		DB: db,
	}, nil
}

func initStorage(cfg *Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (store *Storage) GetCurrency() (*[]currency, error) {
	result := make([]currency, 0)
	rows, err := store.DB.Query("select currencyfrom, currencyto from currency")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		curr := currency{}
		if err := rows.Scan(&curr); err != nil {
			return nil, err
		}
		result = append(result, curr)
	}
	return &result, nil
}

func (store *Storage) PostCurrency(crc *currency) error {
	_, err := store.DB.Exec("insert into currency (currencyfrom, currencyto, well) values ($1,$2,$3)",
		crc.CurrencyFrom,
		crc.CurrencyTo,
		crc.Value)
	if err != nil {
		return err
	}
	return nil
}

func (store *Storage) PutCurrency(crc *currency) error {
	_, err := store.DB.Exec("update currency set (currencyfrom=$1, currencyto=$2, well=$3)",
		crc.CurrencyFrom,
		crc.CurrencyTo,
		crc.Value)
	if err != nil {
		return err
	}
	return nil
}
