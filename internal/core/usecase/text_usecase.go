package usecase

import "github.com/br4tech/go-with-gemini/internal/core/port"

type TextUseCase struct {
}

func NewTextUseCase() port.ITextUseCase {
	return &TextUseCase{}
}
