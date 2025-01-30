package main

import (
	"fmt"
	"os"

	"github.com/loag/ai-lib/pkg/client"
	"github.com/rs/zerolog/log"
)

func main() {
	apiKey := os.Getenv("ANTHROPIC_KEY")

	conf := client.AIConfig{
		ApiKey:       apiKey,
		Model:        "claude-3-5-sonnet-20240620",
		SystemPrompt: "You are a helpful assistant.",
		Temperature:  0.7,
		TopP:         1.0,
		MaxTokens:    2480,
	}

	anthropic, err := client.NewAI(conf)
	if err != nil {
		log.Err(err).Msg("Error creating AI")
		return
	}

	res, err := anthropic.GetCompletion("Hello, how are you?")
	if err != nil {
		log.Err(err).Msg("Error sending message")
		return
	}

	fmt.Println(res)
}
