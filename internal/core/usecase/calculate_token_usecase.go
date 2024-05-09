package usecase

import (
	"errors"

	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/pkoukk/tiktoken-go"
)

const MAX_TOKENS_PER_REQUEST = 3000

type CalculateTokenUseCase struct {
	opinionRepository port.IOpinionRepository
}

func NewCalculateTokenUseCase(
	opinionRepository port.IOpinionRepository,
) port.ICalculateTokenUseCase {
	return &CalculateTokenUseCase{
		opinionRepository: opinionRepository,
	}
}

func (uc CalculateTokenUseCase) CalculateToken(productID int) ([]domain.Opinion, error) {
	selectedOpinions, err := uc.opinionRepository.FindByProductId(productID)
	if err != nil {
		return nil, err
	}

	if selectedOpinions == nil {
		return nil, errors.New("selectedOpinions is nil")
	}

	usedTokensCount := 0
	var opinions []domain.Opinion

	if selectedOpinions != nil {
		for _, opinion := range selectedOpinions {
			opinionTokens := calculate(opinion.Content)

			if usedTokensCount+opinionTokens > MAX_TOKENS_PER_REQUEST {
				break
			} else {
				opinions = append(opinions, *opinion)
				usedTokensCount += opinionTokens
			}
		}
	}

	return opinions, nil
}

func calculate(content string) int {
	encoding := "cl100k_base"

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		return 0
	}

	token := tke.Encode(content, nil, nil)

	return len(token)
}
