package ui

import tea "github.com/charmbracelet/bubbletea"

type helpModel struct{}

func initialHelpModel() helpModel {
	return helpModel{}
}

func (m helpModel) Update(msg tea.Msg) (helpModel, tea.Cmd) {
	return m, nil
}

func (m helpModel) View() string {
	return `
HELP

↑/↓  move
space select
h     help
esc   back
`
}
