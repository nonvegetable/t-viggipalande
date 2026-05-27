package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Experience struct {
	Title       string
	Meta        string
	Description string
}

var experienceList = []Experience{
	{Title: "Technical Consultant Intern • Axe Finance", Meta: "Jan 2026 - Present", Description: "Worked on enterprise migration flows, API validation, ERP integration, and documentation. Kept legacy systems moving without drama."},
	{Title: "Software Developer Intern • Hindalco Industries", Meta: "May 2025 - Jul 2025", Description: "Built Oracle PL/SQL procedures and supported blockchain integration. Tightened backend logic and data handling."},
	{Title: "Frontend Developer Intern • KJSCE", Meta: "Jul 2024 - Dec 2024", Description: "Built the Minors-Honors and Open-Elective portals. Improved navigation, reduced friction, and shipped cleaner flows."},
	{Title: "Digital Marketing Intern • KJSCE", Meta: "Jun 2023 - Jul 2024", Description: "Created 120+ posts across social platforms. Learned consistency, iteration, and how small improvements add up."},
}

func getExperienceView(m Model, width, height int) string {
	experienceStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	boldStyle := m.renderer.NewStyle().
		Foreground(m.theme.primary).
		Bold(true)
	metaStyle := m.renderer.NewStyle().
		Foreground(m.theme.secondary)

	usableHeight := height - 2
	usableWidth := width - 6

	var expSections []string
	for _, exp := range experienceList {
		expTitle := boldStyle.Render(exp.Title)
		expMeta := metaStyle.Render(exp.Meta)
		
		expDesc := lipgloss.NewStyle().Width(usableWidth).Render(exp.Description)
		
		expSections = append(expSections, expTitle+"\n"+expMeta+"\n"+expDesc+"\n")
	}

	experienceText := boldStyle.Render("# Experience") + "\n\n" +
		lipgloss.JoinVertical(lipgloss.Left, expSections...)

	experienceContent := m.renderer.NewStyle().
		Width(usableWidth).
		Height(usableHeight).
		Render(experienceText)

	totalLines := lipgloss.Height(experienceContent)
	scrollViewTotalLines = totalLines
	scrollViewUsableHeight = usableHeight
	if totalLines > usableHeight {
		experienceContentLines := strings.Split(experienceContent, "\n")
		startLine := m.scrollOffset
		endLine := min(int(startLine)+usableHeight, totalLines)
		experienceContent = strings.Join(experienceContentLines[startLine:endLine], "\n")
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
	return experienceStyle.Render(lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		top+"\n"+experienceContent+"\n"+bottom,
	))
}
