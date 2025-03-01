package config

import "github.com/go-sql-driver/mysql"

func MySQLCfg(cfg *Config) *mysql.Config {
	return &mysql.Config{
		User:   cfg.HTTPServer.User,
		Passwd: cfg.HTTPServer.Password,
		Addr:   cfg.HTTPServer.Address,
		DBName: cfg.HTTPServer.DBName,
		Net:    cfg.HTTPServer.Net,
	}
}
