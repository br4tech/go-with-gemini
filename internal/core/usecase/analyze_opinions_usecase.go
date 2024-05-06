package usecase

import (
	"fmt"

	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
)

type AnalyzeOpinionUseCase struct {
	modelo port.IModeloGenerativo
}

func NewAnalyzeOpinionUseCase(modelo port.IModeloGenerativo) port.IAnalyzeOpinionUseCase {
	return &AnalyzeOpinionUseCase{modelo: modelo}
}

func (uc AnalyzeOpinionUseCase) Positive(opinions []domain.Opinion) (*domain.SummaryPositive, error) {

	prompt := `Você será apresentado à opinions de um produto. 
    Seu trabalho é responder um json contendo:
		1) Os principais pontos positivos.
		   Escreva os pontos positivos iniciando com o trecho "Os clientes".
			 Utilize o json schema: {"positivos": "array"}
			 Opinioes: `

	for i, opinion := range opinions {
		prompt += fmt.Sprintf("%d. %s\n", i+1, opinion)
	}

	// var x = uc.modelo.Prompt(prompt)

	// log.Fatal(x)

	return nil, nil
}
