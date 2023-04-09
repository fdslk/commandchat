package commandchat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

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
	messages := Convert2HistoryMessage(nil, true)
	if len(messages) != 0 {
		t.Errorf("should return empty message")
	}
}

func TestShouldReturnEmptyMessageWhenCurrentHistoryIsEmpty(t *testing.T) {
	messages := Convert2HistoryMessage(nil, false)
	if len(messages) != 0 {
		t.Errorf("should return empty message")
	}
}

func TestShouldReturnConvertAllCurrentHistoryToMessageWhenCurrentHistorySizeIsLessThenTwo(t *testing.T) {
	currentHistory := map[string][]interface{}{"user": {"test1"}, "assistant": {"test2"}}
	expectedMessages := []Message{
		{Role: "user", Content: "test1"},
		{Role: "assistant", Content: "test2"},
	}
	messages := Convert2HistoryMessage(currentHistory, false)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}

func TestShouldReturnConvertAllCurrentHistoryToMessageWhenUserCurrentHistorySizeIsEqualToAssistant(t *testing.T) {
	currentHistory := map[string][]interface{}{"user": {"test10", "test11"}, "assistant": {"test20", "test21"}}
	expectedMessages := []Message{
		{Role: "user", Content: "test11"},
		{Role: "assistant", Content: "test21"},
		{Role: "user", Content: "test10"},
		{Role: "assistant", Content: "test20"},
	}
	messages := Convert2HistoryMessage(currentHistory, false)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}

func TestShouldReturnConvertAllCurrentHistoryToMessageWhenUserCurrentHistorySizeIsLessThanAssistant(t *testing.T) {
	currentHistory := map[string][]interface{}{"user": {"test10", "test11"}, "assistant": {"test20", "test21", "test22"}}
	expectedMessages := []Message{
		{Role: "user", Content: "test11"},
		{Role: "assistant", Content: "test21"},
		{Role: "user", Content: "test10"},
		{Role: "assistant", Content: "test20"},
	}
	messages := Convert2HistoryMessage(currentHistory, false)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}

func TestShouldReturnTop2MessageWhenUserCurrentHistorySizeIsGreaterThanTwo(t *testing.T) {
	currentHistory := map[string][]interface{}{"user": {"test10", "test11", "test12"}, "assistant": {"test20", "test21", "test22", "test23"}}
	expectedMessages := []Message{
		{Role: "user", Content: "test12"},
		{Role: "assistant", Content: "test22"},
		{Role: "user", Content: "test11"},
		{Role: "assistant", Content: "test21"},
	}
	messages := Convert2HistoryMessage(currentHistory, false)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}

func TestShouldCreateCompletionRequest(t *testing.T) {
	emptyMessage := []Message{}
	var request CompletionsRequest
	expectedCompletionsRequest := CompletionsRequest{
		Model:       "text-davinci-003",
		Messages:    emptyMessage,
		TopP:        1,
		Temperature: 0.9,
		Stop: StrArray{
			"Human", "AI",
		},
		Prompt:           "test",
		MaxTokens:        1000,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.7,
	}
	completionsReques, _ := CreateCompletionsRequest("test", emptyMessage)
	json.Unmarshal(completionsReques, &request)

	if reflect.DeepEqual(request, expectedCompletionsRequest) {
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
