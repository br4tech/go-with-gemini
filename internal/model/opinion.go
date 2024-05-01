package model

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"gorm.io/gorm"
)

type Opinion struct {
	gorm.Model

	Id        int    `gorm:"primary_key"`
	Content   string `gorm:"not null"`
	ProductID int    `gorm:"column:product_id"`
}

func (model Opinion) ToDomain() *domain.Opinion {
	return &domain.Opinion{
		Content:   model.Content,
		ProductID: model.ProductID,
	}
}

func (model *Opinion) FromDomain(domain *domain.Opinion) {
	model.Content = domain.Content
	model.ProductID = domain.ProductID
}
