package usecase

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
)

type SummaryUseCase struct {
	summaryRepository     port.ISummaryRepository
	calculateTokenUseCase port.ICalculateTokenUseCase
	analyzeOpinionUseCase port.IAnalyzeOpinionUseCase
}

func NewSummaryUseCase(
	summaryRepository port.ISummaryRepository,
	calculateTokenUseCase port.ICalculateTokenUseCase,
	analyzeOpinionUseCase port.IAnalyzeOpinionUseCase,
) port.ISummaryUseCase {
	return &SummaryUseCase{
		summaryRepository:     summaryRepository,
		calculateTokenUseCase: calculateTokenUseCase,
		analyzeOpinionUseCase: analyzeOpinionUseCase,
	}
}

func (uc SummaryUseCase) Positive(productID int) (*domain.SummaryPositive, error) {
	opinions, err := uc.calculateTokenUseCase.CalculateToken(productID)

	if err != nil {
		return nil, err
	}

	summary, err := uc.analyzeOpinionUseCase.Positive(opinions)

	if err != nil {
		return nil, err
	}

	return summary, nil
}

// func (uc SummaryUseCase) CreateSummary(summaryDTO *dto.SummarytDTO) (*domain.Summary, error) {

// 	opinions := make([]domain.Opinion, len(summaryDTO.Opinions))
// 	for i, opinionDTO := range summaryDTO.Opinions {
// 		opinions[i] = *opinionDTO.ToDomain()
// 	}

// 	summary := domain.NewSummary(
// 		summaryDTO.Positive,
// 		summaryDTO.Negative,
// 		opinions,
// 	)

// 	if err := validator.ValidateStruct(summary); err != nil {
// 		return nil, err
// 	}

// 	createdSummary, err := uc.summaryRepository.CreateSummary(summary)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return createdSummary.ToDomain(), nil
// }
