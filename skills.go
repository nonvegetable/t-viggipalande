package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Tool struct {
	Name string
	URL  string
}

type Section struct {
	Title string
	Tools []Tool
}

var toolsList = []Section{
	{
		Title: "Programming Languages",
		Tools: []Tool{
			{Name: "Java", URL: ""},
			{Name: "Python", URL: ""},
			{Name: "JavaScript", URL: ""},
			{Name: "TypeScript", URL: ""},
			{Name: "C", URL: ""},
			{Name: "HTML", URL: ""},
			{Name: "CSS", URL: ""},
			{Name: "PL/SQL", URL: ""},
		},
	},
	{
		Title: "Frontend Frameworks",
		Tools: []Tool{
			{Name: "React.js", URL: ""},
		},
	},
	{
		Title: "Backend Technologies",
		Tools: []Tool{
			{Name: "Node.js", URL: ""},
			{Name: "Express.js", URL: ""},
			{Name: "Flask", URL: ""},
		},
	},
	{
		Title: "Database Systems",
		Tools: []Tool{
			{Name: "MySQL", URL: ""},
			{Name: "MongoDB", URL: ""},
			{Name: "PostgreSQL", URL: ""},
			{Name: "Oracle", URL: ""},
		},
	},
	{
		Title: "Development Tools",
		Tools: []Tool{
			{Name: "Git", URL: ""},
			{Name: "GitHub", URL: ""},
			{Name: "VS Code", URL: ""},
			{Name: "PyCharm", URL: ""},
			{Name: "Docker", URL: ""},
		},
	},
	{
		Title: "Integration & APIs",
		Tools: []Tool{
			{Name: "REST APIs", URL: ""},
			{Name: "Razorpay", URL: ""},
			{Name: "Google APIs", URL: ""},
			{Name: "Supabase", URL: ""},
		},
	},
}

func getSkillsView(m Model, width, height int) string {
	skillsStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	boldStyle := m.renderer.NewStyle().
		Foreground(m.theme.primary).
		Bold(true)

	usableHeight := height - 2
	usableWidth := width - 6

	var toolsSections []string

	var toolStyle = m.renderer.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).MarginRight(1)

	for _, section := range toolsList {
		var tools []string
		for _, tool := range section.Tools {
			if tool.URL != "" {
				tools = append(tools, toolStyle.Render(createHyperlink(tool.URL, tool.Name)))
			} else {
				tools = append(tools, toolStyle.Render(tool.Name))
			}
		}
		toolsSections = append(toolsSections, "## "+section.Title+"\n"+wrapAndJoin(tools, usableWidth)+"\n")
	}

	skillsText := boldStyle.Render("# Tools & Technologies") + "\n\n" +
		lipgloss.JoinVertical(lipgloss.Left, toolsSections...)

	skillsContent := m.renderer.NewStyle().
		Width(usableWidth).
		Height(usableHeight).
		Render(skillsText)

	totalLines := lipgloss.Height(skillsContent)
	scrollViewTotalLines = totalLines
	scrollViewUsableHeight = usableHeight
	if totalLines > usableHeight {
		skillsContentLines := strings.Split(skillsContent, "\n")
		startLine := m.scrollOffset
		endLine := min(int(startLine)+usableHeight, totalLines)
		skillsContent = strings.Join(skillsContentLines[startLine:endLine], "\n")
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
	return skillsStyle.Render(lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		top+"\n"+skillsContent+"\n"+bottom,
	))
}
