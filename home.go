package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func getHomeView(m Model, width, height int) string {
    homeStyle := m.style().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#334155")).
        Padding(1, 2).
        Width(width).
        Height(height)

    heroTitle := lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#F8FAFC")).
        Render(`██╗   ██╗██╗ ██████╗  ██████╗ ██╗
██║   ██║██║██╔════╝ ██╔════╝ ██║
██║   ██║██║██║  ███╗██║  ███╗██║
╚██╗ ██╔╝██║██║   ██║██║   ██║██║
 ╚████╔╝ ██║╚██████╔╝╚██████╔╝██║
  ╚═══╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝`)

    heroSubtitle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#7DD3FC")).
        Render("Full-stack Developer • Builder • Curious by default")
    heroBody := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#CBD5E1")).
        Render("Engineering background, internships, and shipped projects. I like practical systems, strong UX, and shipping under pressure without the drama.")
    links := strings.Join([]string{
        createHyperlink("https://terminal.shop", "terminal.shop"),
        createHyperlink("https://sa1.dev", "sa1.dev"),
        createHyperlink("https://viggipalande.live", "viggipalande.live"),
    }, "  •  ")
    linkLine := lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Render("Inspired by " + links + "  •  Web version")

    heroText := strings.Join([]string{heroTitle, heroSubtitle, "", heroBody, "", linkLine}, "\n")

    homeContent := m.style().
        Width(width - 6).
        MaxWidth(width - 6).
        Align(lipgloss.Center).
        Render(heroText)

    return homeStyle.Render(lipgloss.Place(
        width, height,
        lipgloss.Center, lipgloss.Center,
        homeContent,
    ))
}
