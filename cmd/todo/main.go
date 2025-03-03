package main

import (
	"todo/internal/config"
)

func main() {
	cfg := config.MustConfig()

	log := logger_cfg.SetupLogger(cfg.Env)

	log.Info("one")
	log.Debug("two")
	log.Error("three")
}
