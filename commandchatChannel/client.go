package commandchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func CreateCompletionsRequest(question string) ([]byte, error) {

	newRequest := CompletionsRequest{
		Model:       "text-davinci-003",
		TopP:        1,
		Temperature: 0.9,
		Stop: StrArray{
			"Human", "AI",
		},
		Prompt:           question,
		MaxTokens:        1000,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.7,
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

func Chat(requestBytes []byte) (http.Response, error) {
	apiKey := os.Getenv("API_KEY")
	request, err := http.NewRequest(http.MethodPost, COMPLETIONS_URL, bytes.NewBuffer(requestBytes))

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
