//go:build wireinject
// +build wireinject

package main

import (
	"fmt"

	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/adapter"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/core/usecase"
	"github.com/br4tech/go-with-gemini/internal/handler"
	"github.com/br4tech/go-with-gemini/internal/repository"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

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

func InitializeSummaryHandler(db *gorm.DB) *handler.SummaryHandler {
	wire.Build(
		repository.NewSummaryRepository,
		usecase.NewSummaryUseCase,
		handler.NewSummaryHandler,
		provideConfig,
		provideOpinionRepository,
		provideModeloGenerativo,
		provideCalculateTokenUseCase,
		providerAnalyzeOpinionUseCase,
	)

	return &handler.SummaryHandler{}
}

func InitializeOpinionHandler(db *gorm.DB) *handler.OpinionHandler {
	wire.Build(
		repository.NewOpinionRepository,
		usecase.NewOpinionUseCase,
		handler.NewOpinionHandler,
	)
	return &handler.OpinionHandler{}
}

func InitializeProductHandler(db *gorm.DB) *handler.ProductHandler {
	wire.Build(
		repository.NewProductRepository,
		usecase.NewProductUseCase,
		handler.NewProductHandler,
	)

	return &handler.ProductHandler{}
}
