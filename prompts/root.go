package prompts

import (
	"fmt"
	"github.com/JaquesBoeno/CommitWise/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Model struct {
	Questions utils.QuestionLinkedList
	Colors    utils.Colors
	Answers   map[string]string

	ShownAnswered   string
	CurrentQuestion *utils.QuestionNode
	Cursor          int
	TextInput       textinput.Model
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
		switch msg.String() {

		// default bindings
		case "ctrl+c":
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

	str.WriteString(fmt.Sprintf("\033[%sm? \033[0m", m.Colors.Green))
	str.WriteString(fmt.Sprintf("\033[1m%s:\033[0m ", m.CurrentQuestion.Label))
	str.WriteString("\n")

	switch m.CurrentQuestion.Type {
	case "select":
		str.WriteString(selectRender(&m))
	case "text":
		str.WriteString(textRender(&m))
	}

	str.WriteString("\nPress ctrl+c to quit.\n")
	return str.String()
}

func nextPrompt(value string, m *Model) {
	str := strings.Builder{}
	m.Answers[m.CurrentQuestion.Id] = value
	str.WriteString(fmt.Sprintf("\033[%sm? \033[0m", m.Colors.Green))
	str.WriteString(fmt.Sprintf("\033[1m%s:\033[0m ", m.CurrentQuestion.Label))
	str.WriteString(fmt.Sprintf("%s\n", value))
	m.ShownAnswered = m.ShownAnswered + str.String()

	m.CurrentQuestion = m.CurrentQuestion.NextQuest

	if m.CurrentQuestion.Type == "text" {
		m.TextInput = newTextInput(newTextInputData{})
	}
}
