package repository

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/model"
	"gorm.io/gorm"
)

type SummaryRepository struct {
	db *gorm.DB
}

func NewSummaryRepository(db *gorm.DB) port.ISummaryRepository {
	return &SummaryRepository{db: db}
}

func (r *SummaryRepository) Find(id int) (*domain.Summary, error) {
	var summary model.Summary

	if err := r.db.Find(id).First(&summary).Error; err != nil {
		return nil, err
	}

	return summary.ToDomain(), nil
}

func (r *SummaryRepository) CreateSummary(summary *domain.Summary) (*model.Summary, error) {
	summaryModel := new(model.Summary)
	summaryModel.FromDomain(summary)

	if err := r.db.Create(summaryModel).Error; err != nil {
		return nil, err
	}

	return summaryModel, nil
}
