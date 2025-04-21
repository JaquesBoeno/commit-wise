package prompts

import (
	"github.com/JaquesBoeno/CommitWise/questions"
	"github.com/JaquesBoeno/CommitWise/utils"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func selectBindings(key string, m *Model) {
	if data, ok := m.CurrentQuestion.Data.(questions.SelectQuestionData); ok {
		switch key {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			} else {
				m.Cursor = len(data.Options) - 1
			}

		case "down", "j":
			if m.Cursor < len(data.Options)-1 {
				m.Cursor++
			} else {
				m.Cursor = 0
			}

		case "enter":
			nextPrompt(data.Options[m.Cursor].Value, m)
		}
	}
}

func selectRender(m *Model) string {
	str := strings.Builder{}
	if data, ok := m.CurrentQuestion.Data.(questions.SelectQuestionData); ok {
		choices := data.Options
		windowSize := 7

		maxLengthChoiceName := 0
		for _, choice := range choices {
			maxLengthChoiceName = max(maxLengthChoiceName, len(choice.Value))
		}

		for _, choiceIndex := range slideWindowsOptions(len(choices), windowSize, m.Cursor) {
			curChoice := choices[choiceIndex]
			highlightStyle := lipgloss.NewStyle()
			var cursorStr string

			if choiceIndex-m.Cursor == 0 {
				cursorStr = "❯ "
				highlightStyle = highlightStyle.Foreground(lipgloss.Color("32"))
			} else {
				cursorStr = "  "
			}

			str.WriteString(highlightStyle.Render(cursorStr))
			str.WriteString(highlightStyle.Render(utils.PadEnd(curChoice.Value+": ", maxLengthChoiceName+2, ' ')))
			str.WriteString(highlightStyle.Render(curChoice.Desc))
			str.WriteString("\u001B[0m\n")
		}
	}
	return str.String()
}

func slideWindowsOptions(length, windowSize, currentIndex int) []int {
	if length > windowSize {
		indexes := make([]int, windowSize)
		for i := range windowSize {
			offset := windowSize / 2
			choiceIndex := currentIndex + i - offset

			indexes[i] = utils.ArithmeticMod(choiceIndex, length)
		}
		return indexes
	} else {
		indexes := make([]int, length)
		for i := range length {
			indexes[i] = i
		}
		return indexes
	}
}
