package internal

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Store interface {
	GetCurrency() (*[]currencyGetResponse, error)
	PostCurrency(crc *currencyRequest) error
	PutCurrency(crc *currencyRequest) error
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

func (store *Storage) GetCurrency() (*[]currencyGetResponse, error) {
	result := make([]currencyGetResponse, 0)
	rows, err := store.DB.Query("select currencyfrom, currencyto from currency")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		curr := currencyGetResponse{}
		if err := rows.Scan(&curr.CurrencyFrom, &curr.CurrencyTo); err != nil {
			return nil, err
		}

		result = append(result, curr)
	}

	return &result, nil
}

func (store *Storage) PostCurrency(crc *currencyRequest) error {
	updatedTime := time.Now()
	time.Parse("2006-01-02 15:04:05-07", updatedTime.String())
	_, err := store.DB.Exec("insert into currency (currencyfrom, currencyto, well,updated_at) values ($1,$2,$3)",
		crc.CurrencyFrom,
		crc.CurrencyTo,
		crc.Value,
		updatedTime)
	if err != nil {
		return err
	}

	return nil
}

func (store *Storage) PutCurrency(crc *currencyRequest) error {
	updatedTime := time.Now()
	time.Parse("2006-01-02 15:04:05-07", updatedTime.String())
	_, err := store.DB.Exec("update currency set currencyfrom=$1, currencyto=$2, well = $3,updated_at=$4",
		crc.CurrencyFrom,
		crc.CurrencyTo,
		crc.Value,
		updatedTime)
	if err != nil {
		return err
	}

	return nil
}
