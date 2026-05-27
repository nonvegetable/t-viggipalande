package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func getAboutView(m Model, width, height int) string {
	aboutStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	boldStyle := m.renderer.NewStyle().
		Foreground(m.theme.primary).
		Bold(true)

	usableHeight := height - 2
	usableWidth := width - 6

	aboutText := boldStyle.Render("# About Me") + `

- Engineering Background: Built through coursework, internships, and shipping real projects instead of pretending specs are enough.
- Always Learning: Learns by building. Keeps the cycle honest: try, break, fix, improve.
- Football Mindset: Football taught discipline, pressure, and how to stay calm when the plan changes mid-match.
- Curious by Default: Interested in geography, systems, and the patterns that show up before they become obvious.
- Pattern Thinking: Likes finding structure in messy problems. Usually before coffee. Not always by much.
- Built to Improve: Not trying to be loud. Just better than the last version.`

	aboutContent := m.renderer.NewStyle().
		Width(usableWidth).
		Height(usableHeight).
		Render(aboutText)

	totalLines := lipgloss.Height(aboutContent)
	scrollViewTotalLines = totalLines
	scrollViewUsableHeight = usableHeight
	if totalLines > usableHeight {
		aboutContentLines := strings.Split(aboutContent, "\n")
		startLine := m.scrollOffset
		endLine := min(int(startLine)+usableHeight, totalLines)
		aboutContent = strings.Join(aboutContentLines[startLine:endLine], "\n")
	}

	topStyle := m.renderer.NewStyle().
		Width(usableWidth).
		Height(1).
		Align(lipgloss.Center)

	bottomStyle := m.renderer.NewStyle().
		Width(usableWidth).
		Height(1).
		Align(lipgloss.Center)

	top := ""
	bottom := ""

	if m.scrollOffset != 0 {
		top = topStyle.Render("⌃")
	}
	if m.scrollOffset+usableHeight < totalLines {
		bottom = bottomStyle.Render("⌄")
	}
	return aboutStyle.Render(lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		top+"\n"+aboutContent+"\n"+bottom,
	))
}
