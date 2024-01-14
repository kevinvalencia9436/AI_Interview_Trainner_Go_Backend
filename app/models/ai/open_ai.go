package ai_model

import "time"

type OpenAiTtsRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Voice string `json:"voice"`
}

type OpenAiThread struct {
	Messages []MessageRequest `json:"messages"`
}

type OpenAiThreadRequest struct {
	AssistantID string       `json:"assistant_id"`
	Thread      OpenAiThread `json:"thread"`
}

type OpenAiThreadResponse struct {
	Id       string `json:"id"`
	ThreadId string `json:"thread_id"`
}

type OpenAiThreadRunStepResponse struct {
	Object  string       `json:"object"`
	Data    []OpenAiStep `json:"data"`
	FirstId string       `json:"first_id"`
	LastId  string       `json:"last_id"`
	HasMore bool         `json:"has_more"`
}

type OpenAiStep struct {
	ID          string            `json:"id"`
	Object      string            `json:"object"`
	CreatedAt   int64             `json:"created_at"`
	RunID       string            `json:"run_id"`
	AssistantID string            `json:"assistant_id"`
	ThreadID    string            `json:"thread_id"`
	Type        string            `json:"type"`
	Status      string            `json:"status"`
	CancelledAt *time.Time        `json:"cancelled_at"`
	CompletedAt int64             `json:"completed_at"`
	ExpiresAt   *time.Time        `json:"expires_at"`
	FailedAt    *time.Time        `json:"failed_at"`
	LastError   *string           `json:"last_error"`
	StepDetails OpenAiStepDetails `json:"step_details"`
}

type OpenAiStepDetails struct {
	Type            string                `json:"type"`
	MessageCreation OpenAiMessageCreation `json:"message_creation"`
}

type OpenAiMessageCreation struct {
	MessageID string `json:"message_id"`
}

type OpenAiThreadMessageResponse struct {
	ID          string              `json:"id"`
	Object      string              `json:"object"`
	CreatedAt   int64               `json:"created_at"`
	ThreadID    string              `json:"thread_id"`
	Role        string              `json:"role"`
	Content     []OpenAiContentItem `json:"content"`
	FileIDs     []string            `json:"file_ids"`
	AssistantID string              `json:"assistant_id"`
	RunID       string              `json:"run_id"`
	Metadata    OpenAiMetadata      `json:"metadata"`
}

type OpenAiContentItem struct {
	Type string            `json:"type"`
	Text OpenAiTextContent `json:"text"`
}

type OpenAiTextContent struct {
	Value       string             `json:"value"`
	Annotations []OpenAiAnnotation `json:"annotations"`
}

type OpenAiAnnotation struct {
	// Define the structure of annotations if they have a specific structure
}

type OpenAiMetadata struct {
	// Define the structure of metadata if it has a specific structure
}

type MessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAiRequest struct {
	Model    string           `json:"model"`
	Messages []MessageRequest `json:"messages"`
}

type OpenAiChatResponse struct {
	Choices []OpenAiChoice `json:"choices"`
	Created int64          `json:"created"`
	ID      string         `json:"id"`
	Model   string         `json:"model"`
	Object  string         `json:"object"`
	Usage   OpenAiUsage    `json:"usage"`
}

type OpenAiChoice struct {
	FinishReason string          `json:"finish_reason"`
	Index        int             `json:"index"`
	Message      MessageResponse `json:"message"`
}

type MessageResponse struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type OpenAiUsage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
