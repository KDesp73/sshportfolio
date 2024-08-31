package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	navbarItemWidth = 10
	accentColor = lipgloss.Color("57")

	successStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ff00"))
	failureStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff0000"))


	accentForegroundStyle = lipgloss.NewStyle().
		Foreground(accentColor).
		Blink(true)
	
	titleStyle = lipgloss.NewStyle().
		Background(accentColor).
		PaddingLeft(1).
		PaddingRight(1)

	activeStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ffffff")).
		Width(navbarItemWidth)

	inactiveStyle = lipgloss.NewStyle().
		Bold(false).
		Foreground(lipgloss.Color("#585858")).
		Width(navbarItemWidth)

	boldStyle = lipgloss.NewStyle().
		Bold(true)

	underlineStyle = lipgloss.NewStyle().
		Underline(true)

	focusedStyle        = lipgloss.NewStyle().Foreground(accentColor)

	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	cursorStyle         = focusedStyle

	noStyle             = lipgloss.NewStyle()

	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))	

	focusedButton = titleStyle.Render("Submit")

	blurredButton = fmt.Sprintf(" %s ", blurredStyle.Render("Submit"))

	inputBorderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
)
