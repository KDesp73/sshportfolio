package tui

import (
	"fmt"
	"os"
	"strings"

	proj "sshportfolio/internal/projects"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func euMod(a, b int) int {
    if b == 0 {
        fmt.Fprintf(os.Stderr, "Error: Division by zero is undefined.\n");
        return 0;
    }
    
    var r = a % b
    if (r < 0) {
		if b > 0 {
			r += b
		} else {
			r += -b
		}
    }
    return r;
}

func newPaginator(len, perPage int ) paginator.Model {
	p := paginator.New()
	p.KeyMap.NextPage = key.NewBinding(
		key.WithKeys("tab"))
	p.KeyMap.PrevPage = key.NewBinding(
		key.WithKeys("shift+tab"))
	p.Type = paginator.Dots
	p.PerPage = perPage
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len)

	return p
}

func tux() string {
	return "" +
"                           _nnnn_                      \n" +
"                          dGGGGMMb     ,\"\"\"\"\"\"\"\"\"\"\"\"\"\".\n" +
"                         @p~qp~~qMb    | Hello World! |\n" +
"                         M|@||@) M|   _;..............'\n" +
"        @,----.JM| -'\n" +
"      JS^\\__/  qKL\n" +
"     dZP        qKRb\n" +
"    dZP          qKKb\n" +
"   hjm            SMMb\n" +
"   HZM            MMMM\n" +
"   FqM            MMMM\n" +
" __| \".        |\\dS\"qML\n" +
" |    `.       | `' \\Zq\n" +
"_)      \\.___.,|     .'\n" +
"\\____   )MMMMMM|   .'\n" +
"     `-'       `--'\n"
}

func newTable() table.Model {
	columns := []table.Column{
		{Title: "Name", Width: 15},
		{Title: "Language", Width: 10},
		{Title: "Description", Width: 35},
	}

	pool, _ := proj.LoadProjects()

	var rows []table.Row

	for _, project := range pool.Items {
		rows = append(rows, table.Row{
			project.Title, project.Language, project.Description,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

func wrapString(s string, width int) string {
    var wrapped []string
    words := strings.Fields(s) // Split the string into words
    currentLine := ""

    for _, word := range words {
        // Check if adding the next word would exceed the width
        if len(currentLine)+len(word)+1 > width {
            wrapped = append(wrapped, currentLine) // Add the current line to the wrapped lines
            currentLine = word // Start a new line with the current word
        } else {
            if currentLine != "" {
                currentLine += " " // Add a space before the next word
            }
            currentLine += word // Add the word to the current line
        }
    }

    // Add the last line if it's not empty
    if currentLine != "" {
        wrapped = append(wrapped, currentLine)
    }

    return strings.Join(wrapped, "\n") // Join the wrapped lines with newlines
}

