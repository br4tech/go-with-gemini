package usecase

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/dto"
	validator "github.com/br4tech/go-with-gemini/pkg"
)

type SummaryUseCase struct {
	summaryRepository port.ISummaryRepository
}

func NewSummaryUseCase(summaryRepository port.ISummaryRepository) port.ISummaryUseCase {
	return &SummaryUseCase{summaryRepository: summaryRepository}
}

func (uc SummaryUseCase) Find(id int) (*domain.Summary, error) {
	summary, err := uc.summaryRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return summary, nil
}

func (uc SummaryUseCase) CreateSummary(summaryDTO *dto.SummarytDTO) (*domain.Summary, error) {

	opinions := make([]domain.Opinion, len(summaryDTO.Opinions))
	for i, opinionDTO := range summaryDTO.Opinions {
		opinions[i] = *opinionDTO.ToDomain()
	}

	summary := domain.NewSummary(
		summaryDTO.Positive,
		summaryDTO.Negative,
		opinions,
	)

	if err := validator.ValidateStruct(summary); err != nil {
		return nil, err
	}

	createdSummary, err := uc.summaryRepository.CreateSummary(summary)

	if err != nil {
		return nil, err
	}

	return createdSummary.ToDomain(), nil
}
