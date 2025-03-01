package main

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"os"
	"todo/internal/config"
	"todo/storage"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	const OP = "cmd.migrate.Main"
	cfg := config.MustConfig()

	log := config.SetupLogger(cfg.Env)

	storageCfg := config.MySQLCfg(cfg)
	db, err := storage.NewMySQLStorage(storageCfg)
	if err != nil {
		log.Error("failed to init store", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Error(OP, err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		log.Error(OP, err)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Error(OP, err)
		}
		log.Info("migrate up OK")
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Error(OP, err)
		}

		log.Info("migrate down OK")
	}

	os.Exit(0)
}
