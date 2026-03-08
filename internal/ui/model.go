package ui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Danryg/bulletterm/internal/notes"
	"github.com/Danryg/bulletterm/internal/ui/views"
)


type viewState int

const (
	listView  viewState = iota
	notesList viewState = iota
	helpView
)

type Model struct {
	store     *notes.Store
	state     viewState
	notesList views.NotesListModel
	list      views.ListModel
	help      views.HelpModel
}

func InitialModel(store *notes.Store) Model {
	return Model{
		store:     store,
		state:     notesList,
		notesList: views.InitialNotesListModel(store),
		list:      views.InitialListModel(),
		help:      views.InitialHelpModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if openMsg, ok := msg.(views.OpenNoteMsg); ok {
		m.list = m.list.WithNote(openMsg.Note, m.store)
		m.state = listView
		return m, nil
	}

	switch m.state {

	case notesList:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "q" && !m.notesList.IsInputting() {
			return m, tea.Quit
		}
		newNotesList, cmd := m.notesList.Update(msg)
		m.notesList = newNotesList
		return m, cmd

	case listView:
		if key, ok := msg.(tea.KeyMsg); ok && (key.String() == "esc" || key.String() == "q") && !m.list.IsInputting() {
			m.state = notesList
			return m, nil
		}
		newList, cmd := m.list.Update(msg)
		m.list = newList
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
