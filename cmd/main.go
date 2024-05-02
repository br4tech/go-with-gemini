package main

import (
	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/adapter"
	"github.com/br4tech/go-with-gemini/server"
)

func main() {
	cfg := config.GetConfig()
	db := adapter.NewPostgresDatabase(&cfg)

	// userHandler := InitializeUserHandler(db.GetDb())

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
