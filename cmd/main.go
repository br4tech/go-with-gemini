package main

import (
	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/adapter"
	"github.com/br4tech/go-with-gemini/server"
)

func main() {
	cfg := config.GetConfig()

	db := adapter.NewPostgresDatabase(&cfg)
	gemini := adapter.NewGeminiAdapter(&cfg)

	opinionHandler := InitializeOpinionHandler(db.GetDb())
	productHandler := InitializeProductHandler(db.GetDb())
	summaryHandler := InitializeSummaryHandler(db.GetDb())

	server.NewEchoServer(
		&cfg,
		db.GetDb(),
		gemini.GetClient(),
		opinionHandler,
		productHandler,
		summaryHandler,
	).Start()
}
