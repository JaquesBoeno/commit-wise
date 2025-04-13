package prompts

import (
	"github.com/JaquesBoeno/CommitWise/utils"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Answer struct {
	Id    string
	Value string
}

type Model struct {
	Questions []utils.Question
	Answers   []Answer

	ShownAnswered        string
	CurrentQuestion      utils.Question
	CurrentQuestionIndex int
	Cursor               int
}

func InitialModel(config utils.Settings) Model {
	return Model{
		Questions:       config.Questions,
		ShownAnswered:   "",
		CurrentQuestion: config.Questions[0],
		Cursor:          0,
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
		}
	}

	return m, nil
}

func (m Model) View() string {
	str := strings.Builder{}
	str.WriteString(m.ShownAnswered)

	switch m.CurrentQuestion.Type {
	case "select":
		str.WriteString(selectRender(&m))
	}

	str.WriteString("\nPress q to quit.\n")
	return str.String()
}
