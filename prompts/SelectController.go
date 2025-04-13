package prompts

import (
	"fmt"
	"strings"
)

func selectBindings(key string, m *Model) {
	switch key {
	case "up", "k":
		if m.Cursor > 0 {
			m.Cursor--
		} else {
			m.Cursor = len(m.CurrentQuestion.Options) - 1
		}

	case "down", "j":
		if m.Cursor < len(m.CurrentQuestion.Options)-1 {
			m.Cursor++
		} else {
			m.Cursor = 0
		}

	case "enter":
		nextPrompt(m.CurrentQuestion.Options[m.Cursor].Name, m)
		//fmt.Println("teste")
	}

}

func selectRender(m *Model) string {
	str := strings.Builder{}
	choices := m.CurrentQuestion.Options

	windowSize := 7
	for i := range windowSize {
		offset := windowSize / 2
		choiceIndex := i - offset + m.Cursor

		if choiceIndex-m.Cursor == 0 {
			str.WriteString(fmt.Sprintf("\u001B[%sm❯ ", m.Colors.Primary))
		} else {
			str.WriteString("  ")
		}

		if choiceIndex >= 0 && choiceIndex <= len(choices)-1 {
			str.WriteString(choices[choiceIndex].Name)
		} else if choiceIndex < 0 {
			str.WriteString(choices[len(choices)+choiceIndex].Name)
		} else if choiceIndex >= len(choices) {
			str.WriteString(choices[choiceIndex-len(choices)].Name)
		}

		str.WriteString("\u001B[0m\n")
	}

	return str.String()
}
