package db_cfg

import (
	"github.com/go-sql-driver/mysql"
	"todo/internal/config"
)

func MySQLCfg(cfg *config.Config) *mysql.Config {
	return &mysql.Config{
		User:   cfg.HTTPServer.User,
		Passwd: cfg.HTTPServer.Password,
		Addr:   cfg.HTTPServer.Address,
		DBName: cfg.HTTPServer.DBName,
		Net:    cfg.HTTPServer.Net,
	}
}
