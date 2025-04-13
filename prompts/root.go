package prompts

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Model struct {
	Questions []utils.Question
	Answers   map[string]string

	ShownAnswered        string
	CurrentQuestion      utils.Question
	CurrentQuestionIndex int
	Cursor               int
	TextInput            textinput.Model
}

func InitialModel(config utils.Settings) Model {
	return Model{
		Questions:       config.Questions,
		ShownAnswered:   "",
		CurrentQuestion: config.Questions[0],
		Cursor:          0,
		Answers:         map[string]string{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		// default bindings
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		switch m.CurrentQuestion.Type {
		case "select":
			selectBindings(msg.String(), &m)
		case "text":
			textBindings(msg.String(), &m)
		}
	}
	switch m.CurrentQuestion.Type {
	case "text":
		return m, textUpdate(msg, &m)
	}

	return m, nil
}

func (m Model) View() string {
	str := strings.Builder{}
	str.WriteString(m.ShownAnswered)

	switch m.CurrentQuestion.Type {
	case "select":
		str.WriteString(selectRender(&m))
	case "text":
		str.WriteString(textRender(&m))
	}

	str.WriteString("\nPress q to quit.\n")
	return str.String()
}

func nextPrompt(value string, m *Model) {
	m.Answers[m.CurrentQuestion.Id] = value
	m.ShownAnswered = m.ShownAnswered + fmt.Sprintf("%s: %s\n", m.CurrentQuestion.Label, value)
	m.CurrentQuestionIndex++
	m.CurrentQuestion = m.Questions[m.CurrentQuestionIndex]

	if m.CurrentQuestion.Type == "text" {
		m.TextInput = newTextInput(newTextInputData{})
	}
}
