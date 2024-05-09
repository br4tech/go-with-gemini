package dto

import "github.com/br4tech/go-with-gemini/internal/core/domain"

type SummarytDTO struct {
	Positive []string     `json:"positives"`
	Negative []string     `json:"negatives"`
	Opinions []OpinionDTO `json:"opinions"`
}

func (dto SummarytDTO) ToDomain() *domain.Summary {

	opinions := make([]domain.Opinion, len(dto.Opinions))
	for i, opinionDTO := range dto.Opinions {
		opinions[i] = *opinionDTO.ToDomain()
	}

	return &domain.Summary{
		Positive: dto.Positive,
		Negative: dto.Negative,
		Opinions: opinions,
	}
}

func (dto *SummarytDTO) FromDomain(domain *domain.Summary) {
	dto.Positive = domain.Positive
	dto.Negative = domain.Negative

	opinionDTOs := make([]OpinionDTO, len(domain.Opinions))
	for i, opinionDTO := range domain.Opinions {
		opinionDTOs[i].FromDomain(&opinionDTO)
	}
	dto.Opinions = opinionDTOs
}
