package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	Model   string
}

type request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type response struct {
	Model     string    `json:"model"`
	Tone      string    `json:"tone"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
}

func New(baseURL, model string) *Client {
	return &Client{
		BaseURL: baseURL,
		Model:   model,
	}
}

func (c *Client) Transform(text, tone string) (string, error) {
	prompt := fmt.Sprintf("Please rewrite this quote %s. Keep the meaning intact but change the style and delivery. Keep the length roughly the same as the original. Quote: \"%s\"", tone, text)

	reqBody := request{
		Model:  c.Model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := http.Post(c.BaseURL+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to call Ollama: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read Ollama response: %w", err)
	}

	var ollamaResp response
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return "", fmt.Errorf("failed to parse Ollama response: %w", err)
	}

	ollamaResp.Tone = tone

	return ollamaResp.Response, nil
}
