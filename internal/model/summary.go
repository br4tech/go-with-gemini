package model

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"gorm.io/gorm"
)

type Summary struct {
	gorm.Model

	Id       int       `gorm:"primary_key"`
	Positive []string  `gorm:"column:positive"`
	Negative []string  `gorm:"column:negative"`
	Opinions []Opinion `gorm:"2many:opinion_summaries"`
}

func (model Summary) ToDomain() *domain.Summary {
	domain := &domain.Summary{
		Positive: model.Positive,
		Negative: model.Negative,
	}

	for _, opinion := range model.Opinions {
		domain.Opinions = append(domain.Opinions, *opinion.ToDomain())
	}

	return domain
}

func (model *Summary) FromDomain(domain *domain.Summary) {
	model.Positive = domain.Positive
	model.Negative = domain.Negative

	for _, domainOpinion := range domain.Opinions {
		opinion := &Opinion{}
		opinion.FromDomain(&domainOpinion)

		model.Opinions = append(model.Opinions, *opinion)
	}

}
