package types

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type ChatHistory struct {
	ConversationId string `json:"conversation_id"`
	Title          string `json:"title"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Response struct {
	Model              string  `json:"model"`
	CreatedAt          string  `json:"created_at"`
	Message            Message `json:"message"`
	Done               bool    `json:"done"`
	TotalDuration      int64   `json:"total_duration"`
	LoadDuration       int     `json:"load_duration"`
	PromptEvalCount    int     `json:"prompt_eval_count"`
	PromptEvalDuration int     `json:"prompt_eval_duration"`
	EvalCount          int     `json:"eval_count"`
	EvalDuration       int64   `json:"eval_duration"`
}
