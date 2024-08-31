package tui

import (
	"fmt"
	proj "sshportfolio/internal/projects"
	"strings"

	"github.com/charmbracelet/lipgloss"
)



func home(m Model) string {
	var b strings.Builder
	
	b.WriteString(tux())
	b.WriteString("\n\n")

	b.WriteString("  " + lipgloss.NewStyle().Foreground(lipgloss.Color("57")).Bold(true).Render("Welcome") + " to my portfolio!\n\n")
	b.WriteString("  Press `tab` to switch pages. For more information on the controls of this app press `?`")

	return b.String()
}

func projects(m Model) string {
	var (
		boldStyle = lipgloss.NewStyle().Bold(true)
		underlineStyle = lipgloss.NewStyle().Underline(true)
	)
	
	var b strings.Builder
	
	b.WriteString(_table(m))
	b.WriteString("\n\n\n")
	
	pool, _ := proj.LoadProjects()
	project := pool.Items[pool.TitleMap[m.table.SelectedRow()[0]]]

	maxWidth := 60

	titleStyle := lipgloss.NewStyle().Background(lipgloss.Color("57")).PaddingLeft(1).PaddingRight(1)
	b.WriteString(fmt.Sprintf("%s\n\n%s\n\n", titleStyle.Render(project.Title), wrapString(project.Description, maxWidth)))

	b.WriteString(lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("Language: %s", boldStyle.Render(project.Language)),
		fmt.Sprintf("License: %s", boldStyle.Render(project.License)),
		fmt.Sprintf("Link: %s", underlineStyle.Render(project.Link)),
	))

	if project.Content != "" {
		b.WriteString("\n\n" + wrapString(project.Content, maxWidth))
	}

	b.WriteString("\n\n")

	return b.String()
}

func about(m Model) string {
	return m.pages[ABOUT]
}

func contact(m Model) string {
	return m.pages[CONTACT]
}

func page(m Model) string {
	switch m.currentPage {
	case HOME:
		return home(m)
	case PROJECTS:
		return projects(m)
	case ABOUT:
		return about(m)
	case CONTACT:
		return contact(m)
	}
	return "404 Not Found"
}
