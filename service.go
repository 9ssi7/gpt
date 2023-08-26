package gpt

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type client struct {
	apiKey  string
	baseURL string
}

type Config struct {
	ApiKey  string
	BaseURL string
}

func New(cnf Config) Client {
	if cnf.ApiKey == "" {
		panic("GPT: API key is required")
	}
	if cnf.BaseURL == "" {
		cnf.BaseURL = BaseURL
	}
	return &client{
		apiKey:  cnf.ApiKey,
		baseURL: cnf.BaseURL,
	}
}

func (c *client) makeUrl(path string) string {
	return c.baseURL + path
}

func (s *client) Complete(request CompletionRequest) (*CompletionResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	url := s.makeUrl(Paths.Completion)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response CompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
