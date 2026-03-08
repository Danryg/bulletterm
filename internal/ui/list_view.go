package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type listModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialListModel() listModel {
	return listModel{
		choices: []string{
			"Buy carrots",
			"Buy celery",
			"Buy kohlrabi",
		},
		selected: make(map[int]struct{}),
	}
}

func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {

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
			_, ok := m.selected[m.cursor]

			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m listModel) View() string {

	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s + helpStyle.Render("space: select • h: help • q: quit\n")
}
