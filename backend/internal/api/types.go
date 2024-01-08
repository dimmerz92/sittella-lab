package api

type ClientEmail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Message   string `json:"message"`
}

type ClientResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTChoice struct {
	Index        int         `json:"index"`
	Message      GPTMessage  `json:"message"`
	LogProbs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type GPTUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GPTResponse struct {
	Id                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int         `json:"created"`
	Model             string      `json:"model"`
	SystemFingerPrint string      `json:"system_fingerprint"`
	Choices           []GPTChoice `json:"choices"`
	Usage             GPTUsage    `json:"usage"`
}

type GPTRequest struct {
	Messages []GPTMessage `json:"messages"`
	Model    string       `json:"model"`
}
