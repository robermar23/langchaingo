package memory

import "github.com/robermar23/langchaingo/schema"

// ChatMessageHistoryOption is a function for creating new chat message history
// with other than the default values.
type ChatMessageHistoryOption func(m *ChatMessageHistory)

// WithPreviousMessages is an option for NewChatMessageHistory for adding
// previous messages to the history.
func WithPreviousMessages(previousMessages []schema.ChatMessage) ChatMessageHistoryOption {
	return func(m *ChatMessageHistory) {
		m.messages = append(m.messages, previousMessages...)
	}
}

func applyChatOptions(options ...ChatMessageHistoryOption) *ChatMessageHistory {
	h := &ChatMessageHistory{
		messages: make([]schema.ChatMessage, 0),
	}

	for _, option := range options {
		option(h)
	}

	return h
}
