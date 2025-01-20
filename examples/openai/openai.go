package main

import (
	"fmt"
	"os"

	"github.com/loag/ai-lib/pkg/client"
	"github.com/rs/zerolog/log"
)

func main() {
	baseurl := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("OPENAI_KEY")
	openai := client.NewOpenAI(baseurl, apiKey)

	openai.SetSystemPrompt("You are a helpful assistant.")

	res, err := openai.GetCompletion("Hello, how are you?")
	if err != nil {
		log.Err(err).Msg("Error sending message")
		return
	}

	fmt.Println(res.Choices[0].Message.Content)
}
