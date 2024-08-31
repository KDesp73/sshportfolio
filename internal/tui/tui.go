package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	INDEX = 0
	PROJECTS = 1
	ABOUT = 2
	CONTACT = 3
)

type Model struct {
	title string
	keys keyMap
	width int
	height int
	pages []string
	currentPage int
	quitting bool
	enterPressed bool
	ready    bool

	help help.Model
	table table.Model
}



func NewModel() Model {
	pages := []string{
		"Index",
		"Projects",
		"About",
		"Contact",
	}

	return Model{
		title: "SSH Portfolio",
		keys: keys,
		currentPage: 0,
		pages: pages,
		ready: false,
		quitting: false,
		help: help.New(),
		table: newTable(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			m.enterPressed = true
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.NextPage):
			m.currentPage = euMod(m.currentPage+1, len(m.pages))
		case key.Matches(msg, m.keys.PrevPage):
			m.currentPage = euMod(m.currentPage-1, len(m.pages))
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder


	b.WriteString(_navbar(m))
	b.WriteString("\n\n")
	b.WriteString(page(m))
	b.WriteString("\n\n")
	b.WriteString(m.help.View(m.keys))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		b.String(),
	)
}
