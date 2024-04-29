package server

import (
	"fmt"

	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app         *echo.Echo
	db          *gorm.DB
	cfg         *config.Config
	userHandler port.ITextUseCase
}

func NewEchoServer(cfg *config.Config, db *gorm.DB, userHandler port.ITextUseCase) Server {
	return &echoServer{
		app:         echo.New(),
		db:          db,
		cfg:         cfg,
		userHandler: userHandler,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	s.app.POST("/text", s.userHandler.GenerateToken)
	s.app.POST("/image", s.userHandler.CreateUser)
	s.app.GET("/audio", s.userHandler.ValidateAccessToken)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
