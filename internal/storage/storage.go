package storage

import (
	"database/sql"
	"fmt"
	"mentoring/internal/config"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	db, err := initStorage(cfg)
	if err != nil {
		return nil, err
	}
	return &Storage{
		DB: db,
	}, nil
}

func initStorage(cfg *config.Config) (*sql.DB, error) {
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
