package repository

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/model"
	"gorm.io/gorm"
)

type OpinionRepository struct {
	db *gorm.DB
}

func NewOpinionRepository(db *gorm.DB) port.IOpinionRepository {
	return &OpinionRepository{db: db}
}

func (r *OpinionRepository) Find(id int) (*domain.Opinion, error) {
	var opinion model.Opinion

	if err := r.db.Find(id).First(&opinion).Error; err != nil {
		return nil, err
	}

	return opinion.ToDomain(), nil
}

func (r *OpinionRepository) CreateOpinion(opinion *domain.Opinion) (*model.Opinion, error) {
	opinionModel := new(model.Opinion)
	opinionModel.FromDomain(opinion)

	if err := r.db.Create(opinionModel).Error; err != nil {
		return nil, err
	}

	return opinionModel, nil
}