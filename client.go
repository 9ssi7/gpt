package gpt

// Client defines the interface for interacting with the GPT API.
// It provides a method to generate completions based on the provided input.
type Client interface {
	// Complete sends a request to the GPT API to generate a completion based on the provided input.
	// It returns a response containing the model's output or an error if something went wrong.
	Complete(CompletionRequest) (*CompletionResponse, error)
}

type paths struct {
	// Completion is the path for the completions endpoint.
	Completion string
}

// Paths contains the paths for the different GPT API endpoints.
// It is used to construct the full URL for each endpoint.
var Paths = paths{
	Completion: "/v1/chat/completions",
}

// BaseURL is the base URL for the GPT API.
var BaseURL = "https://api.openai.com"
