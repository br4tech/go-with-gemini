package model

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name  string `gorm:"column:name,unique;not null"`
	Code  string `gorm:"column:code,unique;not null"`
	Image string `gorm:"column:image"`
}

func (model Product) ToDomain() *domain.Product {
	return &domain.Product{
		Name:  model.Name,
		Code:  model.Code,
		Image: model.Image,
	}
}

func (model *Product) FromDomain(domain *domain.Product) {
	model.Name = domain.Name
	model.Code = domain.Code
	model.Image = domain.Image
}
