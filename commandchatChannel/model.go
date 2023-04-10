package commandchat

const (
	USER      = "user"
	ASSISTANT = "assistant"
)

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
		Message      Message     `json:"message,omitempty"`
	} `json:"choices,omitempty"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens,omitempty"`
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`

	Error Error `json:"error,omitempty"`
}

// 最下面的是最新的
// Input your question (type `quit` to exit): 都可以
// map[assistant:[你好，有什么可以帮您的吗？ 您好，请问您想写一篇什么主题的文章呢？] user:[你好 我想写一篇文章]]
// {"model":"gpt-3.5-turbo","messages":[{"role":"user","content":"都可以"},{"role":"assistant","content":"您好，请问您想写一篇什么主题的文章呢？"},{"role":"user","content":"我想写一篇文章"},{"role":"assistant","content":"你好，有什么可以帮您的吗？"},{"role":"user","content":"你好"}],"max_tokens":1000,"temperature":0.9,"top_p":1,"stop":["Human","AI"],"presence_penalty":0.7}
// Response status: 200 OK
// 您好，请问有什么需要帮助的吗？
func Convert2HistoryMessage(currentHistoryMap map[string][]interface{}, setting ChatSetting) []Message {
	assistantHistory := ReverseSlice(currentHistoryMap[ASSISTANT])
	userHistory := ReverseSlice(currentHistoryMap[USER])
	var messages []Message
	if currentHistoryMap == nil || setting.ModelName != "gpt-3.5-turbo" || len(userHistory) == 0 || len(assistantHistory) == 0 {
		return messages
	}

	if len(assistantHistory) == 1 && len(userHistory) == 1 {
		messages = append(messages, Message{ASSISTANT, assistantHistory[0].(string)})
		messages = append(messages, Message{USER, userHistory[0].(string)})
		return messages
	}

	if len(userHistory) == len(assistantHistory) {
		for index, history := range userHistory[:2] {
			messages = append(messages, Message{ASSISTANT, assistantHistory[index].(string)}, Message{USER, history.(string)})
		}
	}

	return messages
}
