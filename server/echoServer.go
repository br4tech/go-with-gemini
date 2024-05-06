package server

import (
	"fmt"

	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app            *echo.Echo
	db             *gorm.DB
	cfg            *config.Config
	opinionHandler port.IOpinionHandler
	productHandler port.IProductHandler
	summaryHandler port.ISummaryHandler
}

func NewEchoServer(
	cfg *config.Config,
	db *gorm.DB,
	opinionHandler port.IOpinionHandler,
	productHandler port.IProductHandler,
	summaryHandler port.ISummaryHandler,
) Server {
	return &echoServer{
		app:            echo.New(),
		db:             db,
		cfg:            cfg,
		opinionHandler: opinionHandler,
		productHandler: productHandler,
		summaryHandler: summaryHandler,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	s.app.POST("/product", s.productHandler.CreateProduct)
	s.app.GET("/product/:id", s.productHandler.Find)

	s.app.POST("/opinion", s.opinionHandler.CreateOpinion)
	s.app.GET("/:product_id/opinions", s.opinionHandler.FindByProductId)
	s.app.GET("/opinion/:id", s.opinionHandler.Find)

	s.app.POST("/summary", s.summaryHandler.CreateSummary)
	s.app.GET("/summary", s.summaryHandler.Find)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
