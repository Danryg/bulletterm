package views

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Danryg/bulletterm/internal/notes"
	"github.com/Danryg/bulletterm/internal/ui/styles"
)

type OpenNoteMsg struct {
	Note notes.Note
}

type notesListMode int

const (
	modeList     notesListMode = iota
	modeCreating notesListMode = iota
)

type NotesListModel struct {
	store  *notes.Store
	cursor int
	mode   notesListMode
	input  string
	height int
}

func (m NotesListModel) WithHeight(h int) NotesListModel {
	m.height = h
	return m
}

func InitialNotesListModel(store *notes.Store) NotesListModel {
	return NotesListModel{store: store}
}

func (m NotesListModel) IsInputting() bool {
	return m.mode == modeCreating
}

func (m NotesListModel) Update(msg tea.Msg) (NotesListModel, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	if m.mode == modeCreating {
		switch key.String() {
		case "esc":
			m.mode = modeList
			m.input = ""
		case "enter":
			if strings.TrimSpace(m.input) != "" {
				m.store.Add(m.input, "")
			}
			m.mode = modeList
			m.input = ""
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if len(key.Runes) > 0 {
				m.input += string(key.Runes)
			}
		}
		return m, nil
	}

	allNotes := m.store.All()
	switch key.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(allNotes)-1 {
			m.cursor++
		}
	case "n":
		m.mode = modeCreating
		m.input = ""
	case "enter":
		if len(allNotes) > 0 {
			selected := allNotes[m.cursor]
			return m, func() tea.Msg { return OpenNoteMsg{Note: selected} }
		}
	}

	return m, nil
}

func countTodos(body string) (done, total int) {
	for _, line := range strings.Split(body, "\n") {
		if strings.HasPrefix(line, "[x] ") {
			done++
			total++
		} else if strings.HasPrefix(line, "[ ] ") {
			total++
		}
	}
	return
}

func formatDate(t time.Time) string {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	switch {
	case sameDay(t, today):
		return "Today"
	case sameDay(t, yesterday):
		return "Yesterday"
	default:
		return t.Format("Monday, January 2 2006")
	}
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}

func (m NotesListModel) View() string {
	help := styles.HelpStyle.Render("n: new note • ↑/↓ k/j: navigate • enter: open • q: quit")

	if m.mode == modeCreating {
		content := fmt.Sprintf("New note title:\n\n> %s_\n", m.input)
		return padToBottom(content, styles.HelpStyle.Render("enter: save • esc: cancel"), m.height)
	}

	allNotes := m.store.All()
	s := "Notes\n\n"

	if len(allNotes) == 0 {
		s += "No notes yet.\n"
	} else {
		lastDate := ""
		for i, n := range allNotes {
			date := formatDate(n.CreatedAt)
			if date != lastDate {
				s += fmt.Sprintf("── %s ──\n", date)
				lastDate = date
			}
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			done, total := countTodos(n.Body)
			todoInfo := ""
			if total > 0 {
				todoInfo = fmt.Sprintf(" (%d/%d)", done, total)
			}
			s += fmt.Sprintf("%s %s%s\n", cursor, n.Title, todoInfo)
		}
	}

	return padToBottom(s, help, m.height)
}

