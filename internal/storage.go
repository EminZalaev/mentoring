package internal

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Store interface {
	GetCurrency() (*[]currencyGetResponse, error)
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
		cfg.DBName)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	return db, err
}

func (store *Storage) CloseDBConnection() error {
	if err := store.DB.Close(); err != nil {
		return fmt.Errorf("error close database: %w", err)
	}
	return nil
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
	defer rows.Close()

	return &result, nil
}

func (store *Storage) PostCurrency(crc *currency) error {
	updatedTime := time.Now()
	well := float64(crc.Value)
	_, err := store.DB.Exec("insert into currency (currencyfrom, currencyto, well,updated_at) values ($1,$2,$3,$4)",
		crc.CurrencyFrom,
		crc.CurrencyTo,
		well,
		updatedTime)
	if err != nil {
		return err
	}

	return nil
}

func (store *Storage) PutCurrency(crc *currency) error {
	updatedTime := time.Now()
	well := float64(crc.Value)
	if crc.Value == 0 {
		well = crc.Well
	}
	_, err := store.DB.Exec("update currency set currencyfrom=$1, currencyto=$2, well = $3,updated_at=$4 where currencyfrom=$1 and currencyto=$2",
		crc.CurrencyFrom,
		crc.CurrencyTo,
		well,
		updatedTime)
	if err != nil {
		return err
	}

	return nil
}
