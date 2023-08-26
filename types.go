package gpt

// CompletionRequest represents the input data required to generate completions using the GPT API.
type CompletionRequest struct {
	Model            Model     `json:"model"`             // Model specifies the GPT model to be used for generating the completion.
	Messages         []Message `json:"messages"`          // Messages contains the conversation history.
	Temperature      float32   `json:"temperature"`       // Temperature controls the randomness of the model's output.
	MaxTokens        int       `json:"max_tokens"`        // MaxTokens limits the length of the generated output.
	TopP             float32   `json:"top_p"`             // TopP controls the diversity of the model's output.
	FrequencyPenalty float32   `json:"frequency_penalty"` // FrequencyPenalty penalizes more frequent tokens.
	PresencePenalty  float32   `json:"presence_penalty"`  // PresencePenalty penalizes new tokens.
}

// CompletionResponse represents the model's output for a given CompletionRequest.
type CompletionResponse struct {
	Id      string     `json:"id"`      // Id is a unique identifier for the response.
	Object  string     `json:"object"`  // Object indicates the type of the response.
	Created int        `json:"created"` // Created is a timestamp of when the response was generated.
	Model   Model      `json:"model"`   // Model specifies which GPT model was used.
	Choices []Choice   `json:"choices"` // Choices contains the generated outputs.
	Usage   Completion `json:"usage"`   // Usage provides details about token usage for the request.
}

// Choice represents a single generated output.
type Choice struct {
	Index        int     `json:"index"`         // Index is the position of the choice in the list.
	Message      Message `json:"message"`       // Message contains the generated output.
	FinishReason string  `json:"finish_reason"` // FinishReason indicates why the generation was stopped.
}

// Completion provides details about token usage for a request.
type Completion struct {
	PromptTokens     int `json:"prompt_tokens"`     // PromptTokens is the number of tokens used in the input.
	CompletionTokens int `json:"completion_tokens"` // CompletionTokens is the number of tokens in the generated output.
	TotalTokens      int `json:"total_tokens"`      // TotalTokens is the total number of tokens used.
}

// Message represents a single message in a conversation.
type Message struct {
	Role    Role   `json:"role"`    // Role indicates who sent the message (e.g., user, bot).
	Content string `json:"content"` // Content contains the text of the message.
}

// Model represents the different GPT models available for generating completions.
type Model string

const (
	ModelGpt3Turbo        Model = "gpt-3.5-turbo"
	ModelGpt3Turbo16K0613 Model = "gpt-3.5-turbo-16k-0613"
	ModelGpt3Turbo16K     Model = "gpt-3.5-turbo-16k"
	ModelGpt3Turbo0613    Model = "gpt-3.5-turbo-0613"
	ModelGpt3Turbo0301    Model = "gpt-3.5-turbo-0301"
)

// Role represents the different roles that can send messages in a conversation.
type Role string

const (
	RoleUser   Role = "user"   // RoleUser represents a message sent by the user.
	RoleBot    Role = "bot"    // RoleBot represents a message sent by the bot.
	RoleSystem Role = "system" // RoleSystem represents a system-generated message.
)

// Object represents the different types of responses that can be received from the GPT API.
type Object string

const (
	ObjectChatCompletion Object = "chat.completion" // ObjectChatCompletion indicates a completion response.
)
