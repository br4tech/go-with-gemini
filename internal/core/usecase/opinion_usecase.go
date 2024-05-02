package usecase

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/dto"
	validator "github.com/br4tech/go-with-gemini/pkg"
)

type OpinionUseCase struct {
	opinionRepository port.IOpinionRepository
}

func NewOpinionUseCase(opinionRepository port.IOpinionRepository) port.IOpinionUseCase {
	return &OpinionUseCase{opinionRepository: opinionRepository}
}

func (uc *OpinionUseCase) Find(id int) (*domain.Opinion, error) {
	opinion, err := uc.opinionRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return opinion, nil
}

func (uc *OpinionUseCase) CreateOpinion(opinionDTO *dto.OpinionDTO) (*domain.Opinion, error) {
	opinion := domain.NewOpinion(
		opinionDTO.Content,
		opinionDTO.ProductId,
	)

	if err := validator.ValidateStruct(opinion); err != nil {
		return nil, err
	}

	createdOpinion, err := uc.opinionRepository.CreateOpinion(opinion)

	if err != nil {
		return nil, err
	}
	return createdOpinion.ToDomain(), nil
}
