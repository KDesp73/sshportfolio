package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	HOME = 0
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
	ready    bool

	help help.Model
	table table.Model
	gameInput textinput.Model

	emailInputs []textinput.Model
	emailFocusIndex int
	emailContent textarea.Model
	EmailError error
	emailSubmitPressed bool
}



func NewModel() Model {
	model := Model {
		pages: []string{
			"Home",
			"Projects",
			"About",
			"Contact",
		},
		title: "SSH Portfolio",
		keys: keys,
		currentPage: 0,
		ready: false,
		quitting: false,
		help: help.New(),
		table: newTable(),
		emailInputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range model.emailInputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.CharLimit = 64
			t.Prompt = ""
		case 1:
			t.Placeholder = "Email"
			t.CharLimit = 64
			t.Prompt = ""
		}

		model.emailInputs[i] = t
	}

	model.emailContent = textarea.New()
	model.emailContent.Placeholder = "Body"

	return model
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
		case key.Matches(msg, m.keys.Up) || key.Matches(msg, m.keys.Down):
			if m.currentPage != CONTACT {
				break
			}
			s := msg.String()

			// Cycle indexes
			if s == "up" {
				m.emailFocusIndex--
			} else {
				m.emailFocusIndex++
			}

			if m.emailFocusIndex > len(m.emailInputs)+1 {
				m.emailFocusIndex = 0
			} else if m.emailFocusIndex < 0 {
				m.emailFocusIndex = len(m.emailInputs)+1
			}

			inputCmds := make([]tea.Cmd, len(m.emailInputs)+1)
			for i := 0; i <= len(m.emailInputs); i++ {
				if i == m.emailFocusIndex {
					if i == len(m.emailInputs) {
						inputCmds[i] = m.emailContent.Focus()
					} else {
						inputCmds[i] = m.emailInputs[i].Focus()
						m.emailInputs[i].PromptStyle = focusedStyle
						m.emailInputs[i].TextStyle = focusedStyle

					}
					continue
				}
				// Remove focused state
				if i == len(m.emailInputs) {
					m.emailContent.Blur()
				} else {
					m.emailInputs[i].Blur()
					m.emailInputs[i].PromptStyle = noStyle
					m.emailInputs[i].TextStyle = noStyle
				}
			}
			cmds = append(cmds, tea.Batch(inputCmds...))
		case key.Matches(msg, m.keys.Enter):
			m.EmailError = nil
			if m.emailFocusIndex == len(m.emailInputs)+1 && m.currentPage == CONTACT {
				m.EmailError = Mail(m)

				if m.EmailError == nil {
					m.emailInputs[0].SetValue("")
					m.emailInputs[1].SetValue("")
					m.emailContent.SetValue("")
				}

				m.emailSubmitPressed = true
			}
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
	if m.currentPage == CONTACT {
		m.emailContent, cmd = m.emailContent.Update(msg)
		cmds = append(cmds, cmd)
		cmds = append(cmds, m.updateInputs(msg))
	}
	
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
