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
	{Title: "Technical Consultant Intern\nAxe Finance", Meta: "Jan 2026 - Present", Description: "Migrated client document templates to an internal engine. Developed and tested stored procedures for business validation rules."},
	{Title: "Software Developer Intern\nHindalco Industries", Meta: "May 2025 - Jul 2025", Description: "Automated Scrap Receiving using Oracle PL/SQL. Contributed to workflows for lot ticket generation. Optimized multiple PL/SQL procedures."},
	{Title: "Winner, Best in Category\nCodenovate 2024 Hackathon", Meta: "Nov 2024", Description: "Developed AIML software designed to determine genetic diseases based on analysis of a person’s DNA samples."},
	{Title: "Frontend Developer Intern\nKJSCE", Meta: "Jul 2024 - Dec 2024", Description: "Built Minors-Honors and Open-Elective allotment software. Engineered 6 responsive pages with 20+ reusable components using React.js."},
	{Title: "Top 6%\nNavonmesh 2.0 Hackathon", Meta: "Dec 2024", Description: "Optimized the DNA-based genetic disease detection software to improve accuracy and processing speed."},
	{Title: "Finalist\nParul-Hackverse Hackathon", Meta: "Jan 2025", Description: "Built an IoT system integrating 3+ real-time sensors for automated farm water management."},
	{Title: "Digital Marketing and Content Writer\nKJSCE", Meta: "Jun 2023 - Jul 2024", Description: "Authored and published 120+ posts across official college social media channels (Instagram, Facebook, LinkedIn)."},
}

func getExperienceView(m Model, width, height int) string {
	experienceStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	boldStyle := m.renderer.NewStyle().
		Foreground(m.theme.secondary).
		Bold(true)
	metaStyle := m.renderer.NewStyle().
		Foreground(m.theme.foregroundMuted)

	usableHeight := height - 2
	usableWidth := width - 6

	var expSections []string
	for _, exp := range experienceList {
		expTitle := boldStyle.Render(exp.Title)
		expMeta := metaStyle.Render(exp.Meta)

		expDesc := lipgloss.NewStyle().Width(usableWidth).Render(exp.Description)

		expSections = append(expSections, expTitle+"\n"+expMeta+"\n"+expDesc+"\n")
	}

	experienceText := m.renderer.NewStyle().Foreground(m.theme.primary).Bold(true).MarginBottom(1).Render("Experience") + "\n" +
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
