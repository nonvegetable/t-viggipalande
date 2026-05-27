package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type AboutItem struct {
	Title string
	Desc  string
}

var aboutItems = []AboutItem{
	{Title: "Engineering Background", Desc: "Built through coursework, internships, and shipping real projects instead of pretending specs are enough."},
	{Title: "Always Learning", Desc: "Learns by building. Keeps the cycle honest: try, break, fix, improve."},
	{Title: "Football Mindset", Desc: "Football taught discipline, pressure, and how to stay calm when the plan changes mid-match."},
	{Title: "Curious by Default", Desc: "Interested in geography, systems, and the patterns that show up before they become obvious."},
	{Title: "Pattern Thinking", Desc: "Likes finding structure in messy problems. Usually before coffee. Not always by much."},
	{Title: "Built to Improve", Desc: "Not trying to be loud. Just better than the last version."},
}

func getAboutView(m Model, width, height int) string {
	aboutStyle := m.renderer.NewStyle().Width(width).Height(height)
	usableHeight := height - 2
	usableWidth := width - 6

	subHeaderStyle := m.renderer.NewStyle().Foreground(m.theme.secondary).Bold(true)
	descStyle := m.renderer.NewStyle().Foreground(m.theme.foreground)

	var sections []string

	for _, item := range aboutItems {
		title := subHeaderStyle.Render(item.Title)
		desc := descStyle.Render(item.Desc)
		wrappedItem := lipgloss.NewStyle().Width(usableWidth).Render(title + "\n" + desc)
		sections = append(sections, wrappedItem+"\n")
	}

	aboutText := m.renderer.NewStyle().Foreground(m.theme.primary).Bold(true).MarginBottom(1).Render("About Me") + "\n" +
		lipgloss.JoinVertical(lipgloss.Left, sections...)

	aboutContent := m.renderer.NewStyle().Width(usableWidth).Height(usableHeight).Render(aboutText)

	totalLines := lipgloss.Height(aboutContent)
	scrollViewTotalLines = totalLines
	scrollViewUsableHeight = usableHeight
	if totalLines > usableHeight {
		aboutContentLines := strings.Split(aboutContent, "\n")
		startLine := m.scrollOffset
		endLine := min(int(startLine)+usableHeight, totalLines)
		aboutContent = strings.Join(aboutContentLines[startLine:endLine], "\n")
	}

	topStyle := m.renderer.NewStyle().Width(usableWidth).Height(1).Align(lipgloss.Center)
	bottomStyle := m.renderer.NewStyle().Width(usableWidth).Height(1).Align(lipgloss.Center)

	top := ""
	bottom := ""

	if m.scrollOffset != 0 {
		top = topStyle.Render("⌃")
	}
	if m.scrollOffset+usableHeight < totalLines {
		bottom = bottomStyle.Render("⌄")
	}
	return aboutStyle.Render(lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, top+"\n"+aboutContent+"\n"+bottom))
}
