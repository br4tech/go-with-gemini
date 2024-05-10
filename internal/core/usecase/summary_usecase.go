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

func (uc SummaryUseCase) Negative(productID int) (*domain.SummaryNegative, error) {
	opinions, err := uc.calculateTokenUseCase.CalculateToken(productID)

	if err != nil {
		return nil, err
	}

	summary, err := uc.analyzeOpinionUseCase.Negative(opinions)

	if err != nil {
		return nil, err
	}

	return summary, nil
}
