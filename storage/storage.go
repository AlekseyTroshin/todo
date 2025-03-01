package storage

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg *mysql.Config) (*sql.DB, error) {
	const OP = "storage.NewMySQL"
	storage, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", OP, err)
	}

	if err := storage.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", OP, err)
	}

	return storage, nil
}
