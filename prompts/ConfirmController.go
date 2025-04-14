package prompts

import (
	"fmt"
	"strings"
)

func confirmBindings(key string, m *Model) {
	switch key {
	case "y", "Y":
		m.ConfirmInput = "Yes"
	case "n", "N":
		m.ConfirmInput = "No"
	case "enter":
		if m.ConfirmInput == "" {
			m.ConfirmInput = "No"
		}

		if m.ConfirmInput == "Yes" {
			m.Questions.InsertListAfterNode(m.CurrentQuestion, m.CurrentQuestion.QuestionLinkedList)
		}
		nextPrompt(m.ConfirmInput, m)
	}
}

func confirmRender(m *Model) string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("  %s", m.ConfirmInput))
	return str.String()
}
