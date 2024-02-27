// Set the API_KEY env var to your API key taken from ai.google.dev
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/robermar23/langchaingo/llms"
	"github.com/robermar23/langchaingo/llms/googleai"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	prompt := "Who was the second person to walk on the moon?"
	answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
