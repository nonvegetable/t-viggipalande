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
	name:        createHyperlink("https://developers.reddit.com/apps/mod-notes-memo","ModNotes Memo") + "\nMod Decision Documentation Tool",
	description: "Production-ready Devvit app for Reddit moderators to attach private notes. Built full-stack architecture using Hono, Redis, and React + TypeScript.",
	tools:       []string{"React", "TypeScript", "Redis", "Hono", "Devvit"},
}, {
	name:        "Blockchain Asset Transfer Platform\nCustom Blockchain",
	description: "Full-stack asset transfer platform using Spring Boot and React. Implemented a custom blockchain in Java with designed transactions and ledger validation.",
	tools:       []string{"Java", "Spring Boot", "React.js"},
}, {
	name:        createHyperlink("https://formula-genie-main.vercel.app/", "AI-Powered Formula Extraction") + "\nMath Formula Extraction Tool",
	description: "Full-stack application using React, TypeScript, Supabase, and Gemini API to extract mathematical formulas. Integrated Razorpay.",
	tools:       []string{"React", "TypeScript", "Supabase", "Gemini API", "Razorpay"},
}, {
	name:        createHyperlink("https://pokemon-wordle-murex.vercel.app/", "Pokémon Wordle") + "\nPokemon-Themed Wordle Game",
	description: "Wordle-style guessing game with a Python Flask backend and React frontend. Implemented fuzzy matching to handle user input.",
	tools:       []string{"React.js", "Python", "Flask", "REST APIs"},
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
