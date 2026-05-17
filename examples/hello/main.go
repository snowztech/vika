// hello shows the smallest possible vikusha agent: no tools, one turn,
// answer a question.
//
//	OPENROUTER_API_KEY=... go run ./examples/hello
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/snowztech/vikusha/core/agent"
	"github.com/snowztech/vikusha/core/llm"
)

func main() {
	_ = godotenv.Load()
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENROUTER_API_KEY not set")
	}

	a, err := agent.New(agent.Options{
		Name:         "hello",
		Model:        "openai/gpt-4o-mini",
		SystemPrompt: "You are a friendly assistant. Reply in one short sentence.",
		Provider:     llm.NewOpenRouter(apiKey),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	reply, err := a.Chat(ctx, "lucas", "Say hello and tell me what model you are.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
