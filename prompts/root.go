package prompts

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Model struct {
	Questions       utils.QuestionLinkedList
	Colors          utils.Colors
	Answers         map[string]string
	CurrentQuestion *utils.QuestionNode
	isQuiting       bool
	quitNow         bool

	ShownAnswered string
	Cursor        int
	TextInput     textinput.Model
	ConfirmInput  string
}

type InitData struct {
	Questions utils.QuestionLinkedList
	Colors    utils.Colors
}

func InitialModel(initData InitData) Model {
	return Model{
		Questions:       initData.Questions,
		Colors:          initData.Colors,
		ShownAnswered:   "",
		CurrentQuestion: initData.Questions.Head,
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
		key := msg.String()
		switch key {

		// default bindings
		case "ctrl+c":
			return m, tea.Quit
		}

		switch m.CurrentQuestion.Type {
		case "select":
			selectBindings(key, &m)
		case "text":
			textBindings(key, &m)
		case "confirm":
			confirmBindings(key, &m)
		}
	}
	switch m.CurrentQuestion.Type {
	case "text":
		return m, textUpdate(msg, &m)
	}
	if m.isQuiting {
		m.quitNow = true
	}
	if m.quitNow {
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) View() string {
	str := strings.Builder{}
	str.WriteString(m.ShownAnswered)
	if m.isQuiting {
		return str.String()
	}
	str.WriteString(fmt.Sprintf("\033[%sm? \033[0m", m.Colors.Green))
	str.WriteString(fmt.Sprintf("\033[1m%s:\033[0m ", m.CurrentQuestion.Label))
	str.WriteString("\n")

	switch m.CurrentQuestion.Type {
	case "select":
		str.WriteString(selectRender(&m))
	case "text":
		str.WriteString(textRender(&m))
	case "confirm":
		str.WriteString(confirmRender(&m))
	}

	str.WriteString("\nPress ctrl+c to quit.\n")
	return str.String()
}

func nextPrompt(value string, m *Model) {
	str := strings.Builder{}
	m.Answers[m.CurrentQuestion.Id] = value
	str.WriteString(fmt.Sprintf("\033[%sm? \033[0m", m.Colors.Green))
	str.WriteString(fmt.Sprintf("\033[1m%s:\033[0m ", m.CurrentQuestion.Label))
	str.WriteString(fmt.Sprintf("\033[%sm%s\033[0m\n", m.Colors.Primary, value))
	m.ShownAnswered = m.ShownAnswered + str.String()

	if m.CurrentQuestion.NextQuest != nil {
		m.CurrentQuestion = m.CurrentQuestion.NextQuest
		switch m.CurrentQuestion.Type {
		case "text":
			m.TextInput = newTextInput(newTextInputData{})
		case "confirm":
			m.ConfirmInput = ""
		}
	} else {
		m.isQuiting = true
	}
}
