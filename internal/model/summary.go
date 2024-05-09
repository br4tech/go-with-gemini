package model

import (
	"strings"

	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"gorm.io/gorm"
)

type Summary struct {
	gorm.Model

	Positives string    `gorm:"type:text"`
	Negatives string    `gorm:"type:text"`
	Opinions  []Opinion `gorm:"many2many:opinion_summaries;"`
}

func (model *Summary) ToDomain() *domain.Summary {
	domain := &domain.Summary{
		Positive: strings.Split(model.Positives, ","),
		Negative: strings.Split(model.Negatives, ","),
	}

	for _, opinion := range model.Opinions {
		domain.Opinions = append(domain.Opinions, *opinion.ToDomain())
	}

	return domain
}

func (model *Summary) FromDomain(domain *domain.Summary) {
	model.Positives = strings.Join(domain.Positive, ",")
	model.Negatives = strings.Join(domain.Negative, ",")

	for _, domainOpinion := range domain.Opinions {
		opinion := &Opinion{}
		opinion.FromDomain(&domainOpinion)

		model.Opinions = append(model.Opinions, *opinion)
	}
}
