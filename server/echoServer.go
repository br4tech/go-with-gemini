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
	app            *echo.Echo
	db             *gorm.DB
	cfg            *config.Config
	opinionHandler port.IOpinionHandler
	productHandler port.IProductHandler
	summyHandler   port.ISummaryHandler
}

func NewEchoServer(cfg *config.Config, db *gorm.DB,
	opinionHandler port.IOpinionHandler,
	productHandler port.IProductHandler,
	summyHandler port.ISummaryHandler,
) Server {
	return &echoServer{
		app:            echo.New(),
		db:             db,
		cfg:            cfg,
		opinionHandler: opinionHandler,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	s.app.POST("/product", s.productHandler.CreateOpinion)
	s.app.GET("/product", s.productHandler.Find)

	s.app.POST("/opinion", s.opinionHandler.CreateOpinion)
	s.app.GET("/opinion", s.opinionHandler.Find)

	s.app.POST("/summary", s.summyHandler.CreateOpinion)
	s.app.GET("/summary", s.summyHandler.Find)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
