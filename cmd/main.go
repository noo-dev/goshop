package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/quangdangfit/gocommon/logger"
	httpServer "goshop/internal/server/http"
	"goshop/pkg/config"
	"goshop/pkg/database"
)

func main() {
	cfg := config.GetConfig()
	logger.Initialize(cfg.Environment)
	db, err := database.Connect(cfg.DatabaseUri)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	validator := validator.New(validator.WithRequiredStructEnabled())

	httpSrvr := httpServer.NewServer(db, validator)
	if err = httpSrvr.Run(); err != nil {
		logger.Fatal(err)
	}
}
