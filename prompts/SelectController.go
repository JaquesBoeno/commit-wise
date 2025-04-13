package prompts

import "strings"

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
	}
}

func selectRender(m *Model) string {
	str := strings.Builder{}

	str.WriteString(m.CurrentQuestion.Label)
	str.WriteString("\n\n")
	choices := m.CurrentQuestion.Options

	windowSize := 5
	for i := range windowSize {
		offset := windowSize / 2
		choiceIndex := i - offset + m.Cursor

		if choiceIndex-m.Cursor == 0 {
			str.WriteString("❯ ")
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

		str.WriteString("\n")
	}

	str.WriteString("\n")
	str.WriteString(choices[m.Cursor].Name)
	str.WriteString("\n")

	return str.String()
}
