package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Enter key.Binding
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Help  key.Binding
	Quit  key.Binding
	NextPage  key.Binding
	PrevPage  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right, k.NextPage, k.PrevPage}, // first column
		{k.Help, k.Quit},                // second column
	}
}

var keys = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
	),
	NextPage: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next page"),
	),
	PrevPage: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "prev spage"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	// Left: key.NewBinding(
	// 	key.WithKeys("left", "h"),
	// 	key.WithHelp("←/h", "move left"),
	// ),
	// Right: key.NewBinding(
	// 	key.WithKeys("right", "l"),
	// 	key.WithHelp("→/l", "move right"),
	// ),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func _navbar(m Model) string {
	var b strings.Builder
	b.WriteString("\n┌")
	b.WriteString(strings.Repeat("─", len(m.title) + 2)) // title
	for range m.pages {
		b.WriteString("┬")
		b.WriteString(strings.Repeat("─", navbarItemWidth+2))
	}
	b.WriteString("┐\n")

	b.WriteString(fmt.Sprintf(" │ %s │ ", accentForegroundStyle.Render(m.title)))
	for i, page := range m.pages {
		if m.currentPage == i {
			b.WriteString(activeStyle.Render(page) + " │ ")
		} else {
			b.WriteString(inactiveStyle.Render(page) + " │ ")
		}
	}
	b.WriteString("\n└")
	b.WriteString(strings.Repeat("─", len(m.title) + 2)) // title
	for range m.pages {
		b.WriteString("┴")
		b.WriteString(strings.Repeat("─", navbarItemWidth+2))
	}
	b.WriteString("┘")

	return b.String()
}

func _table(m Model) string {
	return m.table.View()
}

