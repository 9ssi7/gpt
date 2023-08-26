package gpt

type CompletionRequest struct {
	Model            Model     `json:"model"`
	Messages         []Message `json:"messages"`
	Temperature      float32   `json:"temperature"`
	MaxTokens        int       `json:"max_tokens"`
	TopP             float32   `json:"top_p"`
	FrequencyPenalty float32   `json:"frequency_penalty"`
	PresencePenalty  float32   `json:"presence_penalty"`
}

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type Model string

const (
	ModelGpt3Turbo        Model = "gpt-3.5-turbo"
	ModelGpt3Turbo16K0613 Model = "gpt-3.5-turbo-16k-0613"
	ModelGpt3Turbo16K     Model = "gpt-3.5-turbo-16k"
	ModelGpt3Turbo0613    Model = "gpt-3.5-turbo-0613"
	ModelGpt3Turbo0301    Model = "gpt-3.5-turbo-0301"
)

type Role string

const (
	RoleUser   Role = "user"
	RoleBot    Role = "bot"
	RoleSystem Role = "system"
)
