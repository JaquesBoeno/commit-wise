package prompts

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Question struct {
	Id        string
	Type      string
	Label     string
	Min       int8
	Max       int8
	Options   []Option
	Questions []Question
}
type Option struct {
	Name string
	Desc string
}

type Answer struct {
	Id    string
	Value string
}

type Model struct {
	Questions []Question
	Answers   []Answer

	ShownAnswered        string
	CurrentQuestion      Question
	CurrentQuestionIndex int
	Cursor               int
}

func InitialModel() Model {
	return Model{
		Questions: []Question{{
			Id:        "type",
			Type:      "select",
			Label:     "type of change",
			Min:       0,
			Max:       0,
			Options:   []Option{{Name: "fix", Desc: "fix a bug"}, {Name: "style", Desc: "css style"}},
			Questions: nil,
		}},
		ShownAnswered: "",
		CurrentQuestion: Question{
			Id:        "type",
			Type:      "select",
			Label:     "type of change",
			Min:       0,
			Max:       0,
			Options:   []Option{{Name: "fix", Desc: "fix a bug"}, {Name: "style", Desc: "css style"}},
			Questions: nil,
		},
		Cursor: 0,
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

		str.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice))
	}
	return str.String()
}
