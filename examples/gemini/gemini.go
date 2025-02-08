package main

import (
	"fmt"
	"log"
	"os"

	"github.com/loag/ai-lib/pkg/client"
)

func main() {

	conf := client.AIConfig{
		ApiKey:       os.Getenv("GEMINI_KEY"),
		Model:        "gemini-1.5-flash",
		SystemPrompt: "You are a helpful assistant.",
	}

	ai, err := client.NewAI(conf)
	if err != nil {
		log.Fatal(err)
	}

	res, err := ai.GetCompletion("Hello, how are you?")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
