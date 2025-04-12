package prompts

import (
	"fmt"
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
	str.WriteString(m.ShownAnswered + "\n\n")

	switch m.CurrentQuestion.Type {
	case "select":
		str.WriteString(selectRender(&m))
	}

	str.WriteString("\nPress q to quit.\n")
	return str.String()
}

func selectBindings(key string, m *Model) {
	switch key {
	case "up", "k":
		if m.Cursor > 0 {
			m.Cursor--
		}

	case "down", "j":
		if m.Cursor < len(m.CurrentQuestion.Options)-1 {
			m.Cursor++
		}
	}
}
func selectRender(m *Model) string {
	str := strings.Builder{}

	for i, choice := range m.CurrentQuestion.Options {
		cursor := " "
		if m.Cursor == i {
			cursor = "➜" // possible cursors "❯" "➜" ">"
		}

		checked := " "

		str.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name))
	}
	return str.String()
}
