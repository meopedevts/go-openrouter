package openrouter_test

import (
	"encoding/json"
	"testing"

	openrouter "github.com/revrost/go-openrouter"
)

// ChatContent json.Marshal tests
//
// Tests the case where MultiContent is not empty
func TestChatContentMarshalJSON_MultiContent(t *testing.T) {
	parts := []openrouter.ChatMessagePart{
		{
			Type: openrouter.ChatMessagePartTypeText,
			Text: "What is in this image?",
		},
		{
			Type: openrouter.ChatMessagePartTypeImageURL,
			ImageURL: &openrouter.ChatMessageImageURL{
				URL: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg",
			},
		},
	}
	message := openrouter.ChatCompletionMessage{
		Role: openrouter.ChatMessageRoleUser,
		// Content: "This is a simple content",
		ContentValue: openrouter.NewMultiContent(parts),
	}

	expected := `{"role":"user","content":[{"type":"text","text":"What is in this image?"},{"type":"image_url","image_url":{"url":"https://upload.wikimedia.org/wikipedia/commons/thumb/d/dd/Gfp-wisconsin-madison-the-nature-boardwalk.jpg/2560px-Gfp-wisconsin-madison-the-nature-boardwalk.jpg"}}]}`
	marshalAndValidate(t, message, expected)
}

// Tests the case where Content is used (MultiContent is empty)
func TestChatContentMarshalJSON_Content(t *testing.T) {
	message := openrouter.ChatCompletionMessage{
		Role:         openrouter.ChatMessageRoleUser,
		ContentValue: openrouter.NewContent("This is a simple content"),
	}

	expected := `{"role":"user","content":"This is a simple content"}`
	marshalAndValidate(t, message, expected)
}

// Tests the case where Deprecated Content is used (ContentValue is empty)
func TestChatContentMarshalJSON_DeprecatedContent(t *testing.T) {
	message := openrouter.ChatCompletionMessage{
		Role:    openrouter.ChatMessageRoleUser,
		Content: "This is a simple content",
	}

	expected := `{"role":"user","content":"This is a simple content"}`
	marshalAndValidate(t, message, expected)
}

// Tests the case where Content and ContentValue are both used
func TestChatContentMarshalJSON_ContentAndContentValue(t *testing.T) {
	message := openrouter.ChatCompletionMessage{
		Role:         openrouter.ChatMessageRoleUser,
		Content:      "This is the deprecated content",
		ContentValue: openrouter.NewContent("This is the new content"),
	}

	expected := `{"role":"user","content":"This is the deprecated content"}`
	marshalAndValidate(t, message, expected)
}

func marshalAndValidate(t *testing.T, message openrouter.ChatCompletionMessage, expected string) {
	// Calls MarshalJSON
	result, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validates the resulting JSON
	if string(result) != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
