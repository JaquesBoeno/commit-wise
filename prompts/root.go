package prompts

import (
	"github.com/JaquesBoeno/CommitWise/config"
	"github.com/JaquesBoeno/CommitWise/questions"
	"github.com/JaquesBoeno/CommitWise/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Model struct {
	Questions       questions.QuestionLinkedList
	Answers         map[string]string
	CurrentQuestion *questions.QuestionNode
	Colors          config.Colors
	isQuiting       bool
	quitNow         bool

	ShownAnswered string
	Cursor        int
	TextInput     textinput.Model
	width         int
}

// Setup styles
var questionMarkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
var labelStyle = lipgloss.NewStyle().Bold(true)
var valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

type InitData struct {
	Questions questions.QuestionLinkedList
	Colors    config.Colors
}

func InitialModel(initData InitData) Model {
	TextInput := newTextInput(newTextInputData{})
	questionMarkStyle = questionMarkStyle.Foreground(lipgloss.Color(initData.Colors.Secondary))
	valueStyle = valueStyle.Foreground(lipgloss.Color(initData.Colors.Primary))

	switch data := initData.Questions.Head.Data.(type) {
	case questions.TextQuestionData:
		TextInput = newTextInput(newTextInputData{
			placeholder: data.Placeholder,
			max:         data.Max,
			min:         data.Min,
		})
	}
	return Model{
		Questions:       initData.Questions,
		ShownAnswered:   "",
		CurrentQuestion: initData.Questions.Head,
		Cursor:          0,
		Answers:         map[string]string{},
		width:           1,
		TextInput:       TextInput,
		Colors:          initData.Colors,
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
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}

	if m.isQuiting {
		m.quitNow = true
	}
	if m.quitNow {
		return m, tea.Quit
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
	if m.isQuiting {
		return str.String()
	}

	// Render the question label
	str.WriteString(questionMarkStyle.Render("? "))
	str.WriteString(labelStyle.Render(m.CurrentQuestion.Label))
	str.WriteString("\n")

	// Render the Current Question
	switch m.CurrentQuestion.Data.(type) {
	case questions.SelectQuestionData:
		str.WriteString(selectRender(&m))
	case questions.TextQuestionData:
		str.WriteString(textRender(&m))
	}

	str.WriteString("\nPress ctrl+c to quit.\n")
	return utils.WrapText(str.String(), m.width)
}

func nextPrompt(value string, m *Model) {
	m.Answers[m.CurrentQuestion.Key] = value
	m.ShownAnswered = m.ShownAnswered + formatShownAnswered(m.CurrentQuestion.Label, value)

	if m.CurrentQuestion.SubQuestionCondition == value && m.CurrentQuestion.SubQuestionCondition != "" {
		m.Questions.InsertListAfterNode(m.CurrentQuestion, m.CurrentQuestion.SubQuestions)
	}

	if m.CurrentQuestion.NextQuest != nil {
		m.CurrentQuestion = m.CurrentQuestion.NextQuest
		switch data := m.CurrentQuestion.Data.(type) {
		case questions.TextQuestionData:
			m.TextInput = newTextInput(newTextInputData{
				placeholder: data.Placeholder,
				max:         data.Max,
				min:         data.Min,
			})
		}
	} else {
		m.isQuiting = true
	}
}

func formatShownAnswered(label, value string) string {
	str := strings.Builder{}

	str.WriteString(questionMarkStyle.Render("? "))
	str.WriteString(labelStyle.Render(label + ": "))
	str.WriteString(valueStyle.Render(value))
	str.WriteString("\n")

	return str.String()
}
