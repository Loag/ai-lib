package main

import (
	"fmt"
	"os"

	"github.com/loag/ai-lib/pkg/client"
	"github.com/rs/zerolog/log"
)

func main() {
	apiKey := os.Getenv("ANTHROPIC_KEY")
	baseURL := "https://api.anthropic.com/v1/messages"

	anthropic := client.NewAnthropic(apiKey, baseURL)

	anthropic.SetSystemPrompt("You are a helpful assistant.")

	res, err := anthropic.GetCompletion("Hello, how are you?")
	if err != nil {
		log.Err(err).Msg("Error sending message")
		return
	}

	fmt.Println(res.Content[0].Text)
}
