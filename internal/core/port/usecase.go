package port

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/dto"
)

type (
	IOpinionUseCase interface {
		Find(id int) (*domain.Opinion, error)
		FindByProductId(id int) ([]*domain.Opinion, error)
		CreateOpinion(opinionDTO *dto.OpinionDTO) (*domain.Opinion, error)
	}

	ICalculateTokenUseCase interface {
		CalculateToken(productID int) ([]domain.Opinion, error)
	}

	IAnalyzeOpinionUseCase interface {
		Positive(opinions []domain.Opinion) (*domain.SummaryPositive, error)
		Negative(opinions []domain.Opinion) (*domain.SummaryNegative, error)
	}

	IProductUseCase interface {
		Find(id int) (*domain.Product, error)
		CreateProduct(productDTO *dto.ProductDTO) (*domain.Product, error)
	}

	ISummaryUseCase interface {
		Positive(productID int) (*domain.SummaryPositive, error)
		Negative(productID int) (*domain.SummaryNegative, error)
	}
)
