package client

import (
	"errors"
	"strings"

	"github.com/loag/ai-lib/internal"
	"github.com/loag/ai-lib/internal/anthropic"
	"github.com/loag/ai-lib/internal/gemini"
	"github.com/loag/ai-lib/internal/openai"
)

type AIConfig struct {
	ApiKey       string
	Model        string
	SystemPrompt string
	Temperature  float32
	TopP         float32
	MaxTokens    int
}

type AI struct {
	ai internal.AI
}

func NewAI(config AIConfig) (*AI, error) {
	var ai internal.AI

	if config.ApiKey == "" {
		return nil, errors.New("api key is required")
	}

	if config.Model == "" {
		return nil, errors.New("model is required")
	}

	if config.SystemPrompt == "" {
		return nil, errors.New("system prompt is required")
	}

	if config.TopP < 0 || config.TopP > 1 {
		return nil, errors.New("topP must be between 0 and 1")
	}

	if strings.Contains(config.Model, "claude") {
		if config.Temperature < 0 || config.Temperature > 1 {
			return nil, errors.New("temperature must be between 0 and 1")
		}

		base := "https://api.anthropic.com/v1/messages"
		claude := anthropic.NewAnthropic(config.ApiKey, base)
		claude.SetModel(config.Model)
		claude.SetTemperature(config.Temperature)
		claude.SetTopP(config.TopP)
		claude.SetSystemPrompt(config.SystemPrompt)
		claude.SetMaxTokens(config.MaxTokens)

		ai = claude
	} else if strings.Contains(config.Model, "gemini") {
		gemini := gemini.NewGemini(config.ApiKey, "https://generativelanguage.googleapis.com/v1beta/models")
		gemini.SetModel(config.Model)
		gemini.SetSystemPrompt(config.SystemPrompt)

		ai = gemini

	} else {

		if config.Temperature < 0 || config.Temperature > 2 {
			return nil, errors.New("temperature must be between 0 and 1")
		}

		openai := openai.NewOpenAI(config.ApiKey, config.Model)
		openai.SetTemperature(config.Temperature)
		openai.SetTopP(config.TopP)
		openai.SetSystemPrompt(config.SystemPrompt)

		ai = openai
	}

	return &AI{ai: ai}, nil
}

func (a *AI) GetCompletion(prompt string) (string, error) {
	return a.ai.GetCompletion(prompt)
}
