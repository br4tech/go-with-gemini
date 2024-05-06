package port

import (
	"github.com/google/generative-ai-go/genai"
	"gorm.io/gorm"
)

type (
	IDatabase interface {
		GetDb() *gorm.DB
	}

	IModeloGenerativo interface {
		GetClient() *genai.Client
		Prompt(text string) interface{}
	}
)
