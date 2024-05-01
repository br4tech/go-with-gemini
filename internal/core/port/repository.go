package port

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/model"
)

type (
	IOpinionRepository interface {
		Find(id int) (*domain.Opinion, error)
		CreateOpinion(opinion *domain.Opinion) (*model.Opinion, error)
	}

	IProductRepository interface {
		Find(id int) (*domain.Product, error)
		CreateProduct(product *domain.Product) (*model.Product, error)
	}

	ISummaryRepository interface {
		Find(id int) (*domain.Summary, error)
		CreateSummary(summary *domain.Summary) (*model.Summary, error)
	}
)
