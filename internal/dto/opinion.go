package dto

import "github.com/br4tech/go-with-gemini/internal/core/domain"

type OpinionDTO struct {
	Content   string `json:"content"`
	ProductId int    `json:"product_id"`
}

func (dto OpinionDTO) ToDomain() *domain.Opinion {
	return &domain.Opinion{
		Content:   dto.Content,
		ProductID: dto.ProductId,
	}
}

func (dto *OpinionDTO) FromDomain(domain *domain.Opinion) {
	dto.Content = domain.Content
	dto.ProductId = domain.ProductID
}
