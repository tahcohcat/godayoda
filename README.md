# GoDaYoda 🎭

A fun CLI application that fetches random quotes and transforms them with different personalities using Ollama LLM.

## Features

- Fetches random quotes from dummyjson.com
- Transforms quotes with random tones (pirate, sarcastic, Shakespeare, etc.)
- Beautiful colored output
- Uses Ollama for LLM transformations
- Modular architecture with LLM interface for easy extensibility

## Architecture

The application is structured with clean, modular components:

```
├── cmd/godayoda/          # Main application entry point
├── internal/
│   ├── llm/               # LLM interface and implementations
│   │   ├── interface.go   # LLM interface definition
│   │   └── ollama/        # Ollama implementation
│   └── quotes/            # Quote fetching and tone management
└── go.mod
```

- **LLM Interface**: Allows easy swapping between different LLM providers
- **Quotes Module**: Handles quote fetching and tone selection
- **Clean Separation**: Business logic separated from CLI concerns

## Prerequisites

- Go 1.21 or higher
- Ollama running locally on port 11434
- The `llama3.2` model installed in Ollama

### Install Ollama and Model

1. Install Ollama: https://ollama.ai/
2. Pull the model: `ollama pull llama3.2`
3. Start Ollama: `ollama serve` (if not running as service)

## Installation

```bash
# Clone and build
cd /Users/willemkrige/go/src/github/tahcohcat/godayoda
go mod tidy
go build -o godayoda ./cmd/godayoda
```

## Usage

```bash
# Run the application
./godayoda
```

## Example Output

```
✨ "Arrr, there be no great genius without a touch of madness, ye savvy landlubber!" [as a pirate]

Original: "There is no great genius without some touch of madness." - Seneca the Younger
```

## Available Tones

The app randomly selects from these tones:
- As a pirate
- With sarcasm  
- Like Shakespeare
- As a robot
- Like a valley girl
- As a wise old sage
- With extreme enthusiasm
- Like a grumpy old man
- As a superhero
- Like a news anchor
- With a French accent
- As a motivational speaker
- Like a detective noir
- As a medieval knight
- Like a surfer dude