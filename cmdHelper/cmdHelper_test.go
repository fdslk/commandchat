package cmdHelper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

var setting = ChatSetting{"gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"}

func TestShouldReturnInvertedMap(t *testing.T) {
	originalSlice := []interface{}{"one", "two", "three"}
	expectedSlice := []interface{}{"three", "two", "one"}
	invertedSlice := ReverseSlice(originalSlice)

	if !reflect.DeepEqual(invertedSlice, expectedSlice) {
		t.Errorf("should get inverted map: %#v but got %#v", expectedSlice, invertedSlice)
	}
}

func TestShouldAddNewKeyWhenKeyNotExists(t *testing.T) {
	currentMap := map[string][]interface{}{"assistant": {"test20", "test21"}}
	result := UpdateMap("user", "test", currentMap)
	if !reflect.DeepEqual(currentMap, result) {
		t.Errorf("should get inverted map: %#v but got %#v", currentMap, result)
	}
}

func TestShouldUpdateExistKeyWhenKeyExists(t *testing.T) {
	currentMap := map[string][]interface{}{"assistant": {"test20", "test21"}}
	expectedMap := map[string][]interface{}{"assistant": {"test20", "test21", "test22"}}
	result := UpdateMap("assistant", "test22", currentMap)
	if !reflect.DeepEqual(expectedMap, result) {
		t.Errorf("should get inverted map: %#v but got %#v", expectedMap, result)
	}
}

func TestShouldReturnEmptyMessages(t *testing.T) {
	messages := Convert2HistoryMessage(nil, setting)
	if len(messages) != 0 {
		t.Errorf("should return empty message")
	}
}

func TestShouldReturnEmptyMessagesWhenModelIsNotTurbo(t *testing.T) {
	messages := Convert2HistoryMessage(nil,
		ChatSetting{"non-gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"})
	if len(messages) != 0 {
		t.Errorf("should return empty message")
	}
}

func TestShouldReturnEmptyMessageWhenCurrentHistoryIsEmpty(t *testing.T) {
	messages := Convert2HistoryMessage(nil, setting)
	if len(messages) != 0 {
		t.Errorf("should return empty message")
	}
}

func TestShouldReturnEmptyMessageWhenCurrentHistoryNotUserData(t *testing.T) {
	currentHistory := map[string][]interface{}{"assistant": {"test2"}}
	messages := Convert2HistoryMessage(currentHistory, setting)
	if len(messages) != 0 {
		t.Errorf("should return empty message")
	}
}

func TestShouldReturnConvertAllCurrentHistoryToMessageWhenUserMessageLessThanTwo(t *testing.T) {
	currentHistory := map[string][]interface{}{"user": {"test10"}}
	expectedMessages := []Message{
		{Role: "user", Content: "test10"},
	}
	messages := Convert2HistoryMessage(currentHistory, setting)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}

func TestShouldReturnConvertAllCurrentHistoryToMessageWhenUserMessageEqualToTwo(t *testing.T) {
	currentHistory := map[string][]interface{}{"user": {"test10", "test11"}}
	expectedMessages := []Message{
		{Role: "user", Content: "test10"},
		{Role: "user", Content: "test11"},
	}
	messages := Convert2HistoryMessage(currentHistory, setting)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}

func TestShouldReturnConvertAllCurrentHistoryToMessage(t *testing.T) {
	currentHistory := map[string][]interface{}{"user": {"test10", "test11", "test12"}, "assistant": {"test20", "test21"}}
	expectedMessages := []Message{
		{Role: "user", Content: "test11"},
		{Role: "user", Content: "test12"},
	}
	messages := Convert2HistoryMessage(currentHistory, setting)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}

func TestShouldCreateCompletionRequestWithoutPromptWhenModelIsgpt3Dot5turbo(t *testing.T) {
	var request CompletionsRequest
	setting := ChatSetting{"gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"}
	expectedCompletionsRequest := CompletionsRequest{
		Model:            "gpt-3.5-turbo",
		Messages:         []Message{{"user", "test"}},
		TopP:             1,
		Temperature:      0.9,
		MaxTokens:        1000,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.7,
	}
	completionsReques, _ := CreateCompletionsRequest("test", []Message{}, setting)
	json.Unmarshal(completionsReques, &request)

	if !reflect.DeepEqual(request, expectedCompletionsRequest) {
		t.Errorf("should return expected request %v, got %v", expectedCompletionsRequest, request)
	}
}

func TestShouldCreateCompletionRequestWithoutMessageWhenModelIsNotgpt3Dot5turbo(t *testing.T) {
	var request CompletionsRequest
	setting := ChatSetting{"non-gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"}
	expectedCompletionsRequest := CompletionsRequest{
		Model:            "non-gpt-3.5-turbo",
		TopP:             1,
		Temperature:      0.9,
		Prompt:           "test",
		MaxTokens:        1000,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.7,
	}
	completionsReques, _ := CreateCompletionsRequest("test", []Message{}, setting)
	json.Unmarshal(completionsReques, &request)

	if !reflect.DeepEqual(request, expectedCompletionsRequest) {
		t.Errorf("should return expected request %v, got %v", expectedCompletionsRequest, request)
	}
}

func TestShouldCreateCompletionsResponse(t *testing.T) {
	rawResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString("{\"id\": \"testId\"}")),
		Header:     make(http.Header),
	}

	completionsResponse, _ := CreateCompletionsResponse(*rawResponse)

	if completionsResponse.ID != "testId" {
		t.Errorf("should return expected Id testId, got %s", completionsResponse.ID)
	}
}

func TestShouldReturnChatSetting(t *testing.T) {
	setting, _ := ReadFile("../Configuration/setting-test.json")
	expectedSetting := ChatSetting{"gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"}

	if !reflect.DeepEqual(setting, expectedSetting) {
		t.Errorf("should get: %#v but got %#v", expectedSetting, setting)
	}
}

func TestReturnErrorWhenFileNotExist(t *testing.T) {
	_, err := ReadFile("../setting-non-test.json")

	if err.Error() != "open ../setting-non-test.json: no such file or directory" {
		t.Errorf("should get: %s but got %s", "open ../setting-non-test.json: no such file or directory", err.Error())
	}
}

func TestShouldReturnErrorWhenFileIsNotJsonFormat(t *testing.T) {
	_, err := ReadFile("../Configuration/setting-test.xml")

	if err.Error() != "json: cannot unmarshal string into Go value of type cmdHelper.ChatSetting" {
		t.Errorf("should get: %s but got %s", "json: cannot unmarshal string into Go value of type cmdHelper.ChatSetting", err.Error())
	}
}

func TestShouldSaveFile(t *testing.T) {
	expectedSetting := ChatSetting{"gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"}
	err := SaveFile(expectedSetting, "../", "setting-test.json")

	if err != nil {
		t.Errorf("should save file, but get err %s", err.Error())
	}
}

func TestIsChatModel(t *testing.T) {
	expectedSetting := ChatSetting{"gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"}
	if !expectedSetting.IsChatModel() {
		t.Errorf("should be gpt-3.5-turbo")
	}
}

func TestIsNotChatModel(t *testing.T) {
	expectedSetting := ChatSetting{"non-gpt-3.5-turbo", "https://api.openai.com/v1/chat/completions"}
	if expectedSetting.IsChatModel() {
		t.Errorf("should be gpt-3.5-turbo")
	}
}
