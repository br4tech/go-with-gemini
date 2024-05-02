package adapter

import (
	"context"

	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type geminiAdapter struct {
	Client *genai.Client
}

func NewGeminiAdapter(cfg *config.Config) port.IModeloGenerativo {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.Gemini.ApiKey))
	if err != nil {
		panic("Falha ao criar cliente Gemini")
	}
	return &geminiAdapter{Client: client}
}

func (g *geminiAdapter) GetClient() *genai.Client {
	return g.Client
}
