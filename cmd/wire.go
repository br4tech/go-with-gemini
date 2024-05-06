//go:build wireinject
// +build wireinject

package main

import (
	"github.com/br4tech/go-with-gemini/internal/core/usecase"
	"github.com/br4tech/go-with-gemini/internal/handler"
	"github.com/br4tech/go-with-gemini/internal/repository"
	"github.com/google/wire"
	"gorm.io/gorm"
)

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

func InitializeSummaryHandler(db *gorm.DB) *handler.SummaryHandler {
	wire.Build(
		repository.NewSummaryRepository,
		usecase.NewSummaryUseCase,
		handler.NewSummaryHandler,
	)

	return &handler.SummaryHandler{}
}
