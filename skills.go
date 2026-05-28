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
			{Name: "Go (Golang)", URL: ""},
			{Name: "Java", URL: ""},
			{Name: "Python", URL: ""},
			{Name: "SQL", URL: ""},
			{Name: "PL/SQL", URL: ""},
			{Name: "JavaScript", URL: ""},
			{Name: "HTML/CSS", URL: ""},
		},
	},
	{
		Title: "Frameworks & Technologies",
		Tools: []Tool{
			{Name: "React.js", URL: ""},
			{Name: "Flask", URL: ""},
			{Name: "Docker", URL: ""},
			{Name: "Azure", URL: ""},
			{Name: "Oracle E-Business Suite", URL: ""},
		},
	},
	{
		Title: "Databases",
		Tools: []Tool{
			{Name: "Oracle Database", URL: ""},
			{Name: "PostgreSQL", URL: ""},
			{Name: "MySQL", URL: ""},
			{Name: "MongoDB", URL: ""},
			{Name: "MS SQL Server", URL: ""},
		},
	},
	{
		Title: "Developer Tools",
		Tools: []Tool{
			{Name: "Git", URL: ""},
			{Name: "GitHub", URL: ""},
			{Name: "VS Code", URL: ""},
			{Name: "Linux", URL: ""},
			{Name: "MS SQL Server Studio", URL: ""},
		},
	},
}

func getSkillsView(m Model, width, height int) string {
	skillsStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

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
		toolsSections = append(toolsSections, m.renderer.NewStyle().Foreground(m.theme.secondary).Bold(true).Render(section.Title)+"\n"+wrapAndJoin(tools, usableWidth)+"\n")
	}

	skillsText := m.renderer.NewStyle().Foreground(m.theme.primary).Bold(true).MarginBottom(1).Render("Tools & Technologies") + "\n" +
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
