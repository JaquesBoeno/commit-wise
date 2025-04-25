package prompts

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/internal/questions"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type newTextInputData struct {
	placeholder string
	max         int
	min         int
}

func textBindings(key string, m *Model) {
	if data, ok := m.CurrentQuestion.Data.(questions.TextQuestionData); ok {
		switch key {
		case "enter":
			if isValidateInputLength(data.Min, data.Max, m.TextInput.Value()) {
				nextPrompt(m.TextInput.Value(), m)
			}
		}
	}
}

func textUpdate(msg tea.Msg, m *Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)
	return cmd
}

func textRender(m *Model) string {
	str := strings.Builder{}
	if data, ok := m.CurrentQuestion.Data.(questions.TextQuestionData); ok {
		str.WriteString(m.TextInput.View())
		str.WriteString("\n")

		validateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(m.Colors.Green))

		if !isValidateInputLength(data.Min, data.Max, m.TextInput.Value()) {
			validateStyle = validateStyle.Foreground(lipgloss.Color(m.Colors.Red))
		}
		validStr := ""
		if data.Max > 0 {
			validStr = fmt.Sprintf("(%d/%d)", len(m.TextInput.Value()), data.Max)
		} else {
			validStr = fmt.Sprintf("(%d)", len(m.TextInput.Value()))
		}

		str.WriteString(validateStyle.Render(validStr))
	}
	str.WriteString("\n")
	return str.String()
}

func newTextInput(data newTextInputData) textinput.Model {
	ti := textinput.New()
	if data.placeholder != "" {
		ti.Placeholder = data.placeholder
	} else {
		ti.Placeholder = "Write your answers here"
	}

	ti.CharLimit = data.max
	ti.Focus()
	ti.Prompt = ""
	ti.Width = 100

	return ti
}

func isValidateInputLength(min, max int, text string) bool {
	validateMax, validateMin := false, false
	if max <= 0 || len(text) <= max {
		validateMax = true
	}

	if len(text) >= min {
		validateMin = true
	}

	return validateMax && validateMin
}
