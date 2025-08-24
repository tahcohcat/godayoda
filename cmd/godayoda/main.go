package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/tahcohcat/godayoda/internal/llm"
	"github.com/tahcohcat/godayoda/internal/llm/ollama"
	"github.com/tahcohcat/godayoda/internal/quotes"
)

type App struct {
	quotesService *quotes.Service
	llmClient     llm.LLM
}

func newApp() *App {
	quotesService := quotes.NewService("https://dummyjson.com/quotes/random")
	ollamaClient := ollama.New("http://localhost:11434", "llama3.2")

	return &App{
		quotesService: quotesService,
		llmClient:     ollamaClient,
	}
}

func (a *App) runQuoteCommand(cmd *cobra.Command, args []string) error {
	quote, err := a.quotesService.FetchRandom()
	if err != nil {
		return err
	}

	tone := a.quotesService.GetRandomTone()

	transformedQuote, err := a.llmClient.Transform(quote.Quote, tone)
	if err != nil {
		return err
	}

	cyan := color.New(color.FgCyan, color.Bold)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)

	fmt.Println()
	cyan.Printf("✨ [%s] %s ~ %s\n", tone, transformedQuote, quote.Author)
	fmt.Println()
	yellow.Printf("Original: ")
	green.Printf("\"%s\" - %s\n", quote.Quote, quote.Author)
	fmt.Println()

	return nil
}

func main() {
	app := newApp()

	var rootCmd = &cobra.Command{
		Use:   "godayoda",
		Short: "Get a random quote with a twist of personality",
		Long:  "Fetches a random quote and transforms it with a random tone using Ollama LLM",
		RunE:  app.runQuoteCommand,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
