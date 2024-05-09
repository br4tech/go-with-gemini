package main

import (
	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/adapter"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/model"
)

func main() {
	cfg := config.GetConfig()
	db := adapter.NewPostgresDatabase(&cfg)

	ExecuteMigrate(db)
}

func ExecuteMigrate(db port.IDatabase) {
	db.GetDb().Migrator().CreateTable(
		&model.Product{},
		&model.Opinion{},
		&model.Summary{},
	)
}
