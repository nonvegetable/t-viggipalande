package main

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func getContactView(m Model, width, height int) string {
	boldStyle := m.renderer.NewStyle().
		Foreground(m.theme.primary).
		Bold(true)

	content := boldStyle.Render("# Contact Me") + "\n\n"

	if m.contactForm.State == huh.StateCompleted {
		content += m.renderer.Place(width, 11, lipgloss.Center, lipgloss.Center, "Thank you for your message! I will get back to you soon.")
	} else {
		content += m.contactForm.WithWidth(width - 10).WithHeight(12).View()
	}
	content += "\n\nUse Tab to navigate."

	content += "\n" + createHyperlink("mailto:pvighnesh81203@gmail.com", "Email") + " | " +
		createHyperlink("https://github.com/nonvegetable", "GitHub") + " | " +
		createHyperlink("https://www.linkedin.com/in/vighnesh-palande/", "LinkedIn") + " | " +
		createHyperlink("https://drive.google.com/file/d/1Tzh-pc6sZeDywWG6QVOvirKNhLphOgwV/view?usp=sharing", "Resume")

	// Render the content with the theme and dimensions
	container := m.renderer.NewStyle().
		Width(width).
		Height(height).
		Padding(1, 3).
		Render(content)

	return m.renderer.NewStyle().
		Width(width).
		Height(height).
		Render(container)
}