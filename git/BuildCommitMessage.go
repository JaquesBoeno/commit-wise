package utils

import (
	"fmt"
	"strings"
)

func BuildCommitMessage(template string, answers map[string]string, questionList *QuestionLinkedList) string {
	message := template
	for _, key := range questionList.getAllKeys() {
		if _, exists := answers[key]; !exists {
			answers[key] = ""
		}
	}

	for key, value := range answers {
		placeholder := fmt.Sprintf("<%s>", key)
		message = strings.ReplaceAll(message, placeholder, value)
	}

	message = normalizeNewlines(message)

	return message
}

func normalizeNewlines(s string) string {
	for strings.Contains(s, "\n\n\n") {
		s = strings.ReplaceAll(s, "\n\n\n", "\n\n")
	}

	return s
}
