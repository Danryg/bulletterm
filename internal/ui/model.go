package ui

import tea "github.com/charmbracelet/bubbletea"

type viewState int

const (
	listView viewState = iota
	helpView
)

type Model struct {
	state viewState
	list  listModel
	help  helpModel
}

func InitialModel() Model {
	return Model{
		state: listView,
		list:  initialListModel(),
		help:  initialHelpModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if key, ok := msg.(tea.KeyMsg); ok  && key.String() == "q" {
		return m, tea.Quit
	}

	switch m.state {

	case listView:
		newList, cmd := m.list.Update(msg)
		m.list = newList

		// example: switch to help screen
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "h" {
			m.state = helpView
		}

		return m, cmd

	case helpView:
		newHelp, cmd := m.help.Update(msg)
		m.help = newHelp

		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.state = listView
		}

		return m, cmd
	}




	return m, nil
}

func (m Model) View() string {

	switch m.state {

	case listView:
		return m.list.View()

	case helpView:
		return m.help.View()
	}

	return ""
}
