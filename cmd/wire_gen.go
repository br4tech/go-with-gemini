// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/adapter"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/core/usecase"
	"github.com/br4tech/go-with-gemini/internal/handler"
	"github.com/br4tech/go-with-gemini/internal/repository"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeOpinionHandler(db *gorm.DB) *handler.OpinionHandler {
	iOpinionRepository := repository.NewOpinionRepository(db)
	iOpinionUseCase := usecase.NewOpinionUseCase(iOpinionRepository)
	opinionHandler := handler.NewOpinionHandler(iOpinionUseCase)
	return opinionHandler
}

func InitializeProductHandler(db *gorm.DB) *handler.ProductHandler {
	iProductRepository := repository.NewProductRepository(db)
	iProductUseCase := usecase.NewProductUseCase(iProductRepository)
	productHandler := handler.NewProductHandler(iProductUseCase)
	return productHandler
}

func InitializeSummaryHandler(db *gorm.DB) *handler.SummaryHandler {
	iSummaryRepository := repository.NewSummaryRepository(db)
	iOpinionRepository := provideOpinionRepository(db)
	iCalculateTokenUseCase := provideCalculateTokenUseCase(iOpinionRepository)
	config := provideConfig()
	iModeloGenerativo := provideModeloGenerativo(config)
	iAnalyzeOpinionUseCase := providerAnalyzeOpinionUseCase(config, iModeloGenerativo)
	iSummaryUseCase := usecase.NewSummaryUseCase(iSummaryRepository, iCalculateTokenUseCase, iAnalyzeOpinionUseCase)
	summaryHandler := handler.NewSummaryHandler(iSummaryUseCase)
	return summaryHandler
}

// wire.go:

func provideConfig() *config.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("erro fatal ao ler o arquivo de configuração: %v", err))
	}

	cfg := &config.Config{
		App: config.App{
			Port: viper.GetInt("app.server.port"),
		},
		Db: config.Db{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.timezone"),
		},
		Gemini: config.Gemini{
			ApiKey: viper.GetString("gemini.api_key"),
		},
	}

	return cfg
}

func provideOpinionRepository(db *gorm.DB) port.IOpinionRepository {
	return repository.NewOpinionRepository(db)
}

func provideModeloGenerativo(cfg *config.Config) port.IModeloGenerativo {
	return adapter.NewGeminiAdapter(cfg)
}

func provideCalculateTokenUseCase(opinionRepository port.IOpinionRepository) port.ICalculateTokenUseCase {
	return usecase.NewCalculateTokenUseCase(opinionRepository)
}

func providerAnalyzeOpinionUseCase(cfg *config.Config, modelo port.IModeloGenerativo) port.IAnalyzeOpinionUseCase {
	return usecase.NewAnalyzeOpinionUseCase(modelo)
}
