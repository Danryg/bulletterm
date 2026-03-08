package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Danryg/bulletterm/internal/ui/styles"
)

type NotesListModel struct {
	choices []string
	cursor  int
}

func InitialNotesListModel() NotesListModel {
	return NotesListModel{
		choices: []string{
			"Note 1",
			"Note 2",
			"Note 3",
		},
	}
}

func (m NotesListModel) Update(msg tea.Msg) (NotesListModel, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
		}
	}

	return m, nil
}

func (m NotesListModel) View() string {

	s := "Which Note do you want to view?\n\n"

	// Create a list where each item is within a box
	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		itemBox := "--------------------------------------\n"
		itemBox += fmt.Sprintf("%s %s\n", cursor, choice)

		s += itemBox
	}

	s += "--------------------------------------\n"

	return s + styles.HelpStyle.Render("space: select • h: help • q: quit\n")
}
