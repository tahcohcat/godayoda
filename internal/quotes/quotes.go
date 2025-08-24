package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type Quote struct {
	ID     int      `json:"id"`
	Quote  string   `json:"quote"`
	Author string   `json:"author"`
	Length int      `json:"length"`
	Tags   []string `json:"tags"`
}

type Service struct {
	APIUrl string
	tones  []string
}

func NewService(apiURL string) *Service {
	return &Service{
		APIUrl: apiURL,
		tones: []string{
			"as a pirate",
			"with sarcasm",
			"like Shakespeare",
			"as a robot",
			"like a valley girl",
			"as a wise old sage",
			"with extreme enthusiasm",
			"like a grumpy old man",
			"as a superhero",
			"like a news anchor",
			"with a French accent",
			"as a motivational speaker",
			"like a detective noir",
			"as a medieval knight",
			"like a surfer dude",
		},
	}
}

func (s *Service) FetchRandom() (*Quote, error) {
	resp, err := http.Get(s.APIUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch quote: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var quote Quote
	if err := json.Unmarshal(body, &quote); err != nil {
		return nil, fmt.Errorf("failed to parse quote: %w", err)
	}

	return &quote, nil
}

func (s *Service) GetRandomTone() string {
	rand.Seed(time.Now().UnixNano())
	return s.tones[rand.Intn(len(s.tones))]
}
