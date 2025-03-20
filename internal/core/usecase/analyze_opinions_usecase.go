package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/google/generative-ai-go/genai"
)

type AnalyzeOpinionUseCase struct {
	modelo port.IModeloGenerativo
}

type Part_Text struct {
	// Inline text.
	Text string `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

func NewAnalyzeOpinionUseCase(modelo port.IModeloGenerativo) port.IAnalyzeOpinionUseCase {
	return &AnalyzeOpinionUseCase{modelo: modelo}
}

func (uc AnalyzeOpinionUseCase) Positive(opinions []domain.Opinion) (*domain.SummaryPositive, error) {

	prompt := `Você será apresentado à opinions de um produto. 
    Seu trabalho é responder contendo:
		1) No minimo 5 pontos positivos.
		   Escreva os pontos positivos iniciando com o trecho "Os clientes".
			 Utilize o json schema: {\"positive_aspects\": \"array\"} 
			 Opinioes: `

	for i, opinion := range opinions {
		prompt += fmt.Sprintf("%d. %s\n", i+1, opinion)
	}

	var parts = uc.modelo.Prompt(prompt)
	var positives domain.SummaryPositive

	for _, part := range parts {
		switch part.(type) {
		case genai.Text:
			text := part.(genai.Text)
			if err := json.Unmarshal([]byte(text), &positives); err != nil {
				fmt.Println("Erro ao decodificar JSON:", err)
				return nil, nil
			}
		default:
		}

	}

	return &positives, nil
}

func (uc AnalyzeOpinionUseCase) Negative(opinions []domain.Opinion) (*domain.SummaryNegative, error) {

	prompt := `Você será apresentado à opinions de um produto. 
    Seu trabalho é responder contendo:
		1) No minimo 5 pontos negativos.
		   Escreva os pontos negativos iniciando com o trecho Os clientes.
			 Utilize o json schema: {negative_aspects: array} 
			 Opinioes: `

	for i, opinion := range opinions {
		prompt += fmt.Sprintf("%d. %s\n", i+1, opinion)
	}

	var parts = uc.modelo.Prompt(prompt)
	var negatives domain.SummaryNegative

	for _, part := range parts {
		switch part.(type) {
		case genai.Text:
			text := part.(genai.Text)
			if err := json.Unmarshal([]byte(text), &negatives); err != nil {
				fmt.Println("Erro ao decodificar JSON:", err)
				return nil, nil
			}
		default:
		}

	}

	return &negatives, nil
}
