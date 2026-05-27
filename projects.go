package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Project struct {
	name        string
	description string
	tools       []string
}

var projects = []Project{{
	name:        "Formula Genie\nSaaS Platform",
	description: "Problem: formula workflows were scattered. Built: a full-stack SaaS for generation, auth, payments, and usage tracking.",
	tools:       []string{"AI", "React.js", "backend APIs", "Razorpay"},
}, {
	name:        "PokéGuess\nPokémon Wordle",
	description: "Problem: make daily guessing feel fair, quick, and a little annoying. Built: a Wordle-style game with progressive hints and fuzzy matching.",
	tools:       []string{"React.js", "Python", "Flask"},
}, {
	name:        "FONO\nPhone Memory Game",
	description: "Problem: make short-term memory training less boring. Built: a browser game for recalling phone numbers digit by digit.",
	tools:       []string{"JavaScript", "HTML", "CSS"},
}, {
	name:        "Assero\nBlockchain Platform",
	description: "Problem: manage assets with clearer ownership records. Built: Ethereum smart contracts for asset workflows.",
	tools:       []string{"Java", "React.js"},
}}

func getProjectsView(m Model, width, height int) string {

	projectsStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	usableHeight := height - 2
	usableWidth := width - 6

	var projectSections []string

	var toolStyle = m.renderer.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).MarginRight(1)

	for _, p := range projects {
		var toolsText []string
		if len(p.tools) > 0 {
			for _, tool := range p.tools {
				toolsText = append(toolsText, toolStyle.Render(tool))
			}
		}

		projectTitle := m.renderer.NewStyle().Foreground(m.theme.secondary).Bold(true).Render(p.name)
		projectDesc := lipgloss.NewStyle().Width(usableWidth).Render(p.description)
		projectTools := wrapAndJoin(toolsText, usableWidth)

		projectSections = append(projectSections, projectTitle+"\n"+projectDesc+"\n"+projectTools+"\n")
	}

	projectsText := m.renderer.NewStyle().Foreground(m.theme.primary).Bold(true).MarginBottom(1).Render("Projects") + "\n" +
		lipgloss.JoinVertical(lipgloss.Left, projectSections...)

	projectsContent := m.renderer.NewStyle().
		Width(usableWidth).
		Height(usableHeight).
		Render(projectsText)

	totalLines := lipgloss.Height(projectsContent)
	scrollViewTotalLines = totalLines
	scrollViewUsableHeight = usableHeight
	if totalLines > usableHeight {
		projectsContentLines := strings.Split(projectsContent, "\n")
		startLine := m.scrollOffset
		endLine := min(int(startLine)+usableHeight, totalLines)
		projectsContent = strings.Join(projectsContentLines[startLine:endLine], "\n")
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

	return projectsStyle.Render(
		lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, top+"\n"+projectsContent+"\n"+bottom),
	)
}
