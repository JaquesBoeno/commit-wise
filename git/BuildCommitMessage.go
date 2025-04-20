package git

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/questions"
	"github.com/JaquesBoeno/CommitWise/utils"
	"strings"
)

func BuildCommitMessage(template string, answers map[string]string, questionList *questions.QuestionLinkedList) string {
	message := template
	for _, key := range questionList.GetAllKeys() {
		if _, exists := answers[key]; !exists {
			answers[key] = ""
		}
	}

	for key, value := range answers {
		placeholder := fmt.Sprintf("<%s>", key)
		message = strings.ReplaceAll(message, placeholder, value)
	}

	message = utils.NormalizeNewlines(message)

	return message
}
