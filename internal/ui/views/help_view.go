package views

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type HelpModel struct {
	height int
}

func InitialHelpModel() HelpModel {
	return HelpModel{}
}

func (m HelpModel) WithHeight(h int) HelpModel {
	m.height = h
	return m
}

func (m HelpModel) Update(msg tea.Msg) (HelpModel, tea.Cmd) {
	return m, nil
}

func (m HelpModel) View() string {
	content := "HELP\n\n↑/↓  move\nspace select\nh     help\nesc   back\n"
	if m.height == 0 {
		return content
	}
	contentLines := strings.Count(content, "\n")
	padding := m.height - contentLines
	if padding < 0 {
		padding = 0
	}
	return content + strings.Repeat("\n", padding)
}
