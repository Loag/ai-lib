package main

import (
	"fmt"
	"os"

	"github.com/loag/ai-lib/pkg/client"
	"github.com/rs/zerolog/log"
)

func main() {
	apiKey := os.Getenv("OPENAI_KEY")

	conf := client.AIConfig{
		ApiKey:       apiKey,
		Model:        "gpt-4o",
		SystemPrompt: "You are a helpful assistant.",
		Temperature:  0.7,
		TopP:         1.0,
	}

	openai, err := client.NewAI(conf)
	if err != nil {
		log.Err(err).Msg("Error creating AI")
		return
	}

	res, err := openai.GetCompletion("Hello, how are you?")
	if err != nil {
		log.Err(err).Msg("Error sending message")
		return
	}

	fmt.Println(res)
}
