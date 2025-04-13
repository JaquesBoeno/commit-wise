package prompts

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type newTextInputData struct {
	placeholder string
	charLimit   int
}

func textBindings(key string, m *Model) {
	switch key {
	case "enter":
		nextPrompt(m.TextInput.Value(), m)
	}
}

func textUpdate(msg tea.Msg, m *Model) tea.Cmd {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)
	return cmd
}

func textRender(m *Model) string {
	str := strings.Builder{}
	str.WriteString(m.TextInput.View())
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

	ti.CharLimit = 0
	ti.Focus()
	ti.Prompt = ""
	ti.Width = 100

	return ti
}
