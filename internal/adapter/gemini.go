package adapter

import (
	"context"
	"log"

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

func (g *geminiAdapter) Prompt(text string) []genai.Part {
	ctx := context.Background()

	model := g.Client.GenerativeModel("gemini-pro")
	prompt := genai.Text(text)

	resp, err := model.GenerateContent(ctx, prompt)

	if err != nil {
		log.Fatal(err)
	}

	return resp.Candidates[0].Content.Parts
}
