package commandchat

import (
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

func TestShouldReturnEmptyMessageWhenCurrentHistoryIsEmpty(t *testing.T) {
	messages := Convert2HistoryMessage(nil)
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
	messages := Convert2HistoryMessage(currentHistory)
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
	messages := Convert2HistoryMessage(currentHistory)
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
	messages := Convert2HistoryMessage(currentHistory)
	if len(messages) == 0 {
		t.Errorf("should not return empty message")
	}

	if !reflect.DeepEqual(messages, expectedMessages) {
		t.Errorf("should return expectedMessage %v, got %v", expectedMessages, messages)
	}
}
