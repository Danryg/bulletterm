package views

import tea "github.com/charmbracelet/bubbletea"

type HelpModel struct{}

func InitialHelpModel() HelpModel {
	return HelpModel{}
}

func (m HelpModel) Update(msg tea.Msg) (HelpModel, tea.Cmd) {
	return m, nil
}

func (m HelpModel) View() string {
	return `
HELP

↑/↓  move
space select
h     help
esc   back
`
}
