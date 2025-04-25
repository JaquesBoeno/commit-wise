package prompts

import (
	"github.com/JaquesBoeno/CommitWise/internal/questions"
	"github.com/JaquesBoeno/CommitWise/internal/utils"
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
		highlightStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(m.Colors.Primary))
		maxLengthChoiceName := 0

		for _, choice := range choices {
			maxLengthChoiceName = max(maxLengthChoiceName, len(choice.Value))
		}

		slideIndex := slideWindowsOptions(len(choices), windowSize, m.Cursor)

		for _, idx := range slideIndex {
			curChoice := choices[idx]
			isSelect := idx-m.Cursor == 0
			var paddedValue string

			if curChoice.Desc != "" {
				paddedValue = utils.PadEnd(curChoice.Value+": ", maxLengthChoiceName+2, ' ')
			} else {
				paddedValue = utils.PadEnd(curChoice.Value, maxLengthChoiceName+2, ' ')
			}

			cursor := "  "
			if isSelect {
				cursor = "❯ "
			}

			line := strings.Join([]string{cursor, paddedValue, curChoice.Desc}, "")

			if isSelect {
				str.WriteString(highlightStyle.Render(line))
			} else {
				str.WriteString(line)
			}
			str.WriteString("\n")
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
