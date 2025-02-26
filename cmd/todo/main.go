package main

import (
	"todo/internal/config"
)

func main() {
	cfg := config.MustConfig()

	log := config.SetupLogger(cfg.Env)

	log.Info("one")
	log.Debug("two")
	log.Error("three")
}
