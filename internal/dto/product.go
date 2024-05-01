package dto

import "github.com/br4tech/go-with-gemini/internal/core/domain"

type ProductDTO struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Image string `json:"image"`
}

func (dto ProductDTO) ToDomain() *domain.Product {
	return &domain.Product{
		Name:  dto.Name,
		Code:  dto.Code,
		Image: dto.Image,
	}
}

func (dto *ProductDTO) FromDomain(domain *domain.Product) {
	dto.Name = domain.Name
	dto.Code = domain.Code
	dto.Image = domain.Image
}
