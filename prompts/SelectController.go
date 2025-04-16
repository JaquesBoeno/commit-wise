package prompts

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/utils"
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
	maxLengthChoiceName := 0
	for _, choice := range choices {
		maxLengthChoiceName = max(maxLengthChoiceName, len(choice.Name))
	}

	windowSize := 7
	if len(choices) > windowSize {
		for i := range windowSize {
			offset := windowSize / 2
			choiceIndex := i - offset + m.Cursor
			var choice utils.Option

			if choiceIndex-m.Cursor == 0 {
				str.WriteString(fmt.Sprintf("\u001B[%sm❯ ", m.Colors.Primary))
			} else {
				str.WriteString("  ")
			}

			if choiceIndex >= 0 && choiceIndex <= len(choices)-1 {
				choice = choices[choiceIndex]
			} else if choiceIndex < 0 {
				choice = choices[len(choices)+choiceIndex]
			} else if choiceIndex >= len(choices) {
				choice = choices[choiceIndex-len(choices)]
			}
			str.WriteString(padEnd(choice.Name+": ", maxLengthChoiceName+2, ' '))
			str.WriteString(choice.Desc)
			str.WriteString("\u001B[0m\n")
		}
	} else {
		for i, choice := range choices {
			if m.Cursor == i {
				str.WriteString(fmt.Sprintf("\u001B[%sm❯ ", m.Colors.Primary))
			} else {
				str.WriteString("  ")
			}
			str.WriteString(choice.Name)
			str.WriteString("\u001B[0m\n")

		}
	}

	return str.String()
}

func padEnd(str string, length int, pad rune) string {
	for len(str) < length {
		str += string(pad)
	}

	return str
}
