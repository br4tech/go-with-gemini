package dto

import "github.com/br4tech/go-with-gemini/internal/core/domain"

type SummarytDTO struct {
	Positive []string `json:"positives"`
	Negative []string `json:"negatives"`
}

func (dto SummarytDTO) ToDomain() *domain.Summary {
	return &domain.Summary{
		Positive: dto.Positive,
		Negative: dto.Negative,
	}
}

func (dto *SummarytDTO) FromDomain(domain *domain.Summary) {
	dto.Positive = domain.Positive
	dto.Negative = domain.Negative
}
