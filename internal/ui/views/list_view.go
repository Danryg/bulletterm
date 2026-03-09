package views

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Danryg/bulletterm/internal/notes"
	"github.com/Danryg/bulletterm/internal/ui/styles"
)

type listMode int

const (
	listModeNormal   listMode = iota
	listModeCreating listMode = iota
)

type todoItem struct {
	text    string
	checked bool
}

type ListModel struct {
	store  *notes.Store
	note   notes.Note
	items  []todoItem
	cursor int
	mode   listMode
	input  string
	height int
}

func (m ListModel) WithHeight(h int) ListModel {
	m.height = h
	return m
}

func InitialListModel() ListModel {
	return ListModel{}
}

func (m ListModel) IsInputting() bool {
	return m.mode == listModeCreating
}

func (m ListModel) WithNote(n notes.Note, store *notes.Store) ListModel {
	items := []todoItem{}
	for _, line := range strings.Split(n.Body, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.HasPrefix(line, "[x] ") {
			items = append(items, todoItem{text: line[4:], checked: true})
		} else if strings.HasPrefix(line, "[ ] ") {
			items = append(items, todoItem{text: line[4:], checked: false})
		} else {
			items = append(items, todoItem{text: line, checked: false})
		}
	}
	return ListModel{
		store:  store,
		note:   n,
		items:  items,
		cursor: 0,
	}
}

func (m *ListModel) saveItems() {
	lines := make([]string, len(m.items))
	for i, item := range m.items {
		if item.checked {
			lines[i] = "[x] " + item.text
		} else {
			lines[i] = "[ ] " + item.text
		}
	}
	m.note.Body = strings.Join(lines, "\n")
	m.store.Update(m.note.ID, m.note.Title, m.note.Body)
}

func (m ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	if m.mode == listModeCreating {
		switch key.String() {
		case "esc":
			m.mode = listModeNormal
			m.input = ""
		case "enter":
			if strings.TrimSpace(m.input) != "" {
				m.items = append(m.items, todoItem{text: m.input, checked: false})
				m.saveItems()
			}
			m.mode = listModeNormal
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

	switch key.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.items)-1 {
			m.cursor++
		}
	case "enter", " ":
		if len(m.items) > 0 {
			m.items[m.cursor].checked = !m.items[m.cursor].checked
			m.saveItems()
		}
	case "n":
		m.mode = listModeCreating
		m.input = ""
	}

	return m, nil
}

func (m ListModel) View() string {
	if m.mode == listModeCreating {
		content := fmt.Sprintf("%s\n\nNew item:\n\n> %s_\n", m.note.Title, m.input)
		return padToBottom(content, styles.HelpStyle.Render("enter: save • esc: cancel"), m.height)
	}

	s := fmt.Sprintf("%s\n\n", m.note.Title)

	if len(m.items) == 0 {
		s += "No items yet.\n"
	} else {
		for i, item := range m.items {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			checked := " "
			if item.checked {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item.text)
		}
	}

	return padToBottom(s, styles.HelpStyle.Render("n: new item • space/enter: check • ↑/↓ k/j: navigate • esc: back • q: quit"), m.height)
}

func padToBottom(content, help string, height int) string {
	if height == 0 {
		return content + "\n" + help
	}
	contentLines := strings.Count(content, "\n")
	padding := height - contentLines - 1
	if padding < 1 {
		padding = 1
	}
	return content + strings.Repeat("\n", padding) + help
}
