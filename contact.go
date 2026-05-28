package main

func getContactView(m Model, width, height int) string {

	content := m.renderer.NewStyle().Foreground(m.theme.primary).Bold(true).MarginBottom(1).Render("Contact") + "\n"

	content += "Feel free to reach out via any of the links below!\n\n"

	content += "\n" + createHyperlink("https://viggipalande.live", "Website") + " | " +
		createHyperlink("https://www.linkedin.com/in/vighnesh-palande/", "LinkedIn") + " | " +
		createHyperlink("https://github.com/nonvegetable", "GitHub") + " | " +
		createHyperlink("https://drive.google.com/file/d/1xgUsuoLY_VSygUWIel56EmhkebWpqCAE/view?usp=sharing", "Resume") + 
		"\n\nOr email me at: " + createHyperlink("mailto:pvighnesh81203@gmail.com", "pvighnesh81203@gmail.com")

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
