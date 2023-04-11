package commandchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func CreateCompletionsRequest(question string, messages []Message, setting ChatSetting) ([]byte, error) {

	newRequest := CompletionsRequest{
		Model:       setting.ModelName,
		TopP:        1,
		Temperature: 0.9,
		Stop: StrArray{
			"Human", "AI",
		},
		MaxTokens:        1000,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.7,
	}

	if setting.ModelName == "gpt-3.5-turbo" {
		newRequest.Messages = append(messages, Message{USER, question})
	} else {
		newRequest.Prompt = question
	}

	newRequestBytes, err := json.Marshal(newRequest)

	if err != nil {
		fmt.Println("error occurred:", err)
		return newRequestBytes, err
	}

	return newRequestBytes, nil
}

func CreateCompletionsResponse(rawResponse http.Response) (response CompletionsResponse, err error) {
	fmt.Println("Response status:", rawResponse.Status)

	body, err := io.ReadAll(rawResponse.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return response, err
	}

	defer rawResponse.Body.Close()

	err = json.Unmarshal(body, &response)

	return response, err
}

func Chat(requestBytes []byte, setting ChatSetting) (http.Response, error) {
	apiKey := os.Getenv("API_KEY")
	request, err := http.NewRequest(http.MethodPost, setting.ApiUrl, bytes.NewBuffer(requestBytes))

	if err != nil {
		fmt.Println("Error sending request:", err)
		return http.Response{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	//send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return *response, err
	}

	return *response, err
}
