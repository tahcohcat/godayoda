# ChaosMonk 🐵💥

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Chaos Level](https://img.shields.io/badge/Chaos%20Level-💥%20MAXIMUM-red?style=for-the-badge)]()
[![Ollama](https://img.shields.io/badge/Powered%20by-Ollama-FF6B35?style=for-the-badge&logo=ollama)](https://ollama.ai/)
[![Made with](https://img.shields.io/badge/Made%20with-❤️%20%26%20☕-yellow?style=for-the-badge)]()

[![Personality Count](https://img.shields.io/badge/Personalities-15%2B-purple?style=flat-square)]()
[![API](https://img.shields.io/badge/Quotes%20API-DummyJSON-blue?style=flat-square)](https://dummyjson.com/)
[![CLI](https://img.shields.io/badge/Interface-CLI-brightgreen?style=flat-square)]()
[![Chaos Engineering](https://img.shields.io/badge/Philosophy-Chaos%20Engineering-orange?style=flat-square)]()

A chaotic CLI application that fetches random quotes and transforms them with unpredictable personalities using Ollama LLM. Embrace the chaos!

## Features

- Fetches random quotes from dummyjson.com
- Transforms quotes with chaotic tones (pirate, sarcastic, Shakespeare, robot, and more!)
- Beautiful colored output with personality
- Uses Ollama for LLM transformations
- Modular architecture with LLM interface for easy extensibility

## Philosophy

Like a chaos monkey disrupting systems to make them stronger, ChaosMonk disrupts ordinary quotes to make them more entertaining. It injects randomness and personality into wisdom, creating delightfully unpredictable results.

## Architecture

The application is structured with clean, modular components:

```
├── cmd/chaosmonk/         # Main application entry point
├── internal/
│   ├── llm/               # LLM interface and implementations
│   │   ├── interface.go   # LLM interface definition
│   │   └── ollama/        # Ollama implementation
│   └── quotes/            # Quote fetching and chaos injection
└── go.mod
```

- **LLM Interface**: Allows easy swapping between different LLM providers
- **Quotes Module**: Handles quote fetching and chaotic tone selection
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
cd /Users/willemkrige/go/src/github/tahcohcat/chaosmonk
go mod tidy
go build -o chaosmonk ./cmd/chaosmonk
```

## Usage

```bash
# Unleash the chaos
./chaosmonk
```

## Example Output

```
✨ [as a pirate] "Arrr, there be no great genius without a touch of madness, ye landlubber!" ~ Seneca the Younger

Original: "There is no great genius without some touch of madness." - Seneca the Younger
```

## Available Chaotic Tones

The chaos monkey randomly selects from these unpredictable tones:
- As a pirate 🏴‍☠️
- With sarcasm 😏
- Like Shakespeare 🎭
- As a robot 🤖
- Like a valley girl 💅
- As a wise old sage 🧙‍♂️
- With extreme enthusiasm 🎉
- Like a grumpy old man 😠
- As a superhero 🦸‍♂️
- Like a news anchor 📺
- With a French accent 🥖
- As a motivational speaker 💪
- Like a detective noir 🕵️‍♂️
- As a medieval knight ⚔️
- Like a surfer dude 🏄‍♂️

## Why ChaosMonk?

Just as chaos engineering makes systems more resilient by introducing controlled failures, ChaosMonk makes quotes more memorable by introducing controlled chaos. Sometimes wisdom needs a little shake-up to truly resonate!

## Contributing

Embrace the chaos! Feel free to add more chaotic tones, improve the LLM prompts, or extend the architecture. The more unpredictable, the better!