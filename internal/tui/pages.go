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

	b.WriteString("  " + accentForegroundStyle.Bold(true).Render("Welcome") + " to my portfolio!\n\n")
	b.WriteString("  Press `tab` to switch pages. For more information on the controls of this app press `?`")

	return b.String()
}

func projects(m Model) string {
	var b strings.Builder
	
	b.WriteString(_table(m))
	b.WriteString("\n\n\n")
	
	pool, _ := proj.LoadProjects()
	project := pool.Items[pool.TitleMap[m.table.SelectedRow()[0]]]

	maxWidth := 60

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
	var b strings.Builder
	
	b.WriteString("  " + titleStyle.Render("Hello there!") + "\n\n")

	b.WriteString(lipgloss.JoinVertical(lipgloss.Center,
		wrapString("My name is Konstantinos Despoinidis and I am 21 years old", 60),
		wrapString("I am currently studying Information and Electronic Engineering at the International Hellenic University", 60),
	))

	return b.String()
}

func contact(m Model) string {
	var b strings.Builder

	b.WriteString("  " + titleStyle.Bold(true).Render("Reach me @")+"\n\n\n")

	b.WriteString(lipgloss.JoinVertical(lipgloss.Left,
		"  Github: https://github.com/KDesp73\n",
		"  Email: despoinidisk@gmail.com\n",
	))

	b.WriteString("\n\n")


	button := &blurredButton
	if m.emailFocusIndex == len(m.emailInputs)+1 {
		button = &focusedButton
	}

	b.WriteString(inputBorderStyle.PaddingLeft(1).PaddingRight(1).Render(lipgloss.JoinVertical(lipgloss.Center,
		titleStyle.Render("Send me an email!") + "\n",
		m.emailInputs[0].View() + "\n",
		m.emailInputs[1].View() + "\n",
		m.emailContent.View() + "\n",
		fmt.Sprintf("%s", *button),
	)))

	if m.emailSubmitPressed {
		b.WriteString("\n\n")
		if m.EmailError == nil {
			b.WriteString(successStyle.Render("Email sent"))
		} else {
			b.WriteString(failureStyle.Render(m.EmailError.Error()))
		}
	}

	return b.String()
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
