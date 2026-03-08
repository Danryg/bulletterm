package ui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Danryg/bulletterm/internal/ui/views"
)

type viewState int

const (
	listView  viewState = iota
	notesList viewState = iota
	helpView
)

type Model struct {
	state     viewState
	notesList views.NotesListModel
	list      views.ListModel
	help      views.HelpModel
}

func InitialModel() Model {
	return Model{
		state:     notesList,
		notesList: views.InitialNotesListModel(),
		list:      views.InitialListModel(),
		help:      views.InitialHelpModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "q" {
		return m, tea.Quit
	}

	switch m.state {

	case notesList:
		newNotesList, cmd := m.notesList.Update(msg)
		m.notesList = newNotesList

		// example: switch to list view
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "l" {
			m.state = listView
		}

		return m, cmd

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
	case notesList:
		return m.notesList.View()

	case listView:
		return m.list.View()

	case helpView:
		return m.help.View()
	}

	return ""
}
