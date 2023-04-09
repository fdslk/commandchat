package commandchat

const (
	COMPLETIONS_URL = "https://api.openai.com/v1/completions"
	USER            = "user"
	ASSISTANT       = "assistant"
)

type ClientConfigration struct {
	apiKey string
	org    string
}

// Error is the error standard response from the API
type Error struct {
	Message string      `json:"message,omitempty"`
	Type    string      `json:"type,omitempty"`
	Param   interface{} `json:"param,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type StrArray []string

//   messages=[
//         {"role": "system", "content": "You are a helpful assistant."},
//         {"role": "user", "content": "Who won the world series in 2020?"},
//         {"role": "assistant", "content": "The Los Angeles Dodgers won the World Series in 2020."},
//         {"role": "user", "content": "Where was it played?"}
//     ]

type CompletionsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Prompt           string            `json:"prompt,omitempty"`
	Suffix           string            `json:"suffix,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	LogProbs         int               `json:"logprobs,omitempty"`
	Echo             bool              `json:"echo,omitempty"`
	Stop             StrArray          `json:"stop,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	BestOf           int               `json:"best_of,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
}

type CompletionsResponse struct {
	ID      string `json:"id,omitempty"`
	Object  string `json:"object,omitempty"`
	Created int    `json:"created,omitempty"`
	Model   string `json:"model,omitempty"`
	Choices []struct {
		Text         string      `json:"text,omitempty"`
		Index        int         `json:"index,omitempty"`
		Logprobs     interface{} `json:"logprobs,omitempty"`
		FinishReason string      `json:"finish_reason,omitempty"`
	} `json:"choices,omitempty"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens,omitempty"`
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`

	Error Error `json:"error,omitempty"`
}

func Convert2HistoryMessage(currentHistoryMap map[string][]interface{}, firstTalk bool) []Message {
	assistantHistory := ReverseSlice(currentHistoryMap[ASSISTANT])
	userHistory := ReverseSlice(currentHistoryMap[USER])
	var messages []Message
	if currentHistoryMap == nil || firstTalk {
		return messages
	}
	if len(assistantHistory) < 2 && len(userHistory) < 2 {

		messages = append(messages, Message{USER, userHistory[0].(string)})
		messages = append(messages, Message{ASSISTANT, assistantHistory[0].(string)})
		return messages
	}

	if len(userHistory) < len(assistantHistory) {
		for index, history := range userHistory[:2] {
			messages = append(messages, Message{USER, history.(string)}, Message{ASSISTANT, assistantHistory[index+1].(string)})
		}
	}

	if len(userHistory) == len(assistantHistory) {
		for index, history := range userHistory[:2] {
			messages = append(messages, Message{USER, history.(string)}, Message{ASSISTANT, assistantHistory[index].(string)})
		}
	}

	return messages
}
