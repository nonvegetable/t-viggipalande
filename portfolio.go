package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type styleRenderer interface {
	NewStyle() lipgloss.Style
}

type portfolioSection struct {
	key         string
	title       string
	summary     string
	description string
	items       []portfolioItem
}

type portfolioItem struct {
	title       string
	meta        string
	description string
	linkText    string
	linkURL     string
}

type Model struct {
	renderer styleRenderer
	width    int
	height   int
	selected int
	sections []portfolioSection
}

func NewModel(renderer styleRenderer) Model {
	return Model{
		renderer: renderer,
		sections: portfolioSections(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "up", "h", "k", "shift+tab":
			m.selected = (m.selected - 1 + len(m.sections)) % len(m.sections)
		case "right", "down", "l", "j", "tab":
			m.selected = (m.selected + 1) % len(m.sections)
		case "1", "2", "3", "4", "5", "6", "7", "8":
			index := int(msg.Runes[0] - '1')
			if index >= 0 && index < len(m.sections) {
				m.selected = index
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	width := m.width
	height := m.height
	if width <= 0 {
		width = 110
	}
	if height <= 0 {
		height = 42
	}

	base := m.style()
	background := base.
		Background(lipgloss.Color("#09101D")).
		Foreground(lipgloss.Color("#E5EEF8")).
		Width(width).
		Height(height)

	header := getHomeView(m, width-4, 13)
	content := m.renderContent(width - 4)
	footer := m.renderFooter(width - 4)

	stack := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		strings.Repeat("\n", 1),
		content,
		strings.Repeat("\n", 1),
		footer,
	)

	return background.Render(lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		stack,
	))
}

func (m Model) style() lipgloss.Style {
	if m.renderer != nil {
		return m.renderer.NewStyle()
	}
	return lipgloss.NewStyle()
}

func (m Model) currentSection() portfolioSection {
	if len(m.sections) == 0 {
		return portfolioSection{}
	}
	if m.selected < 0 || m.selected >= len(m.sections) {
		return m.sections[0]
	}
	return m.sections[m.selected]
}

func (m Model) renderContent(width int) string {
	section := m.currentSection()
	sidebarWidth := 29
	stacked := width < 92
	contentWidth := width - sidebarWidth - 4
	if stacked {
		contentWidth = width
	}
	if contentWidth < 48 {
		contentWidth = width
	}

	sectionList := m.renderSectionList(sidebarWidth)
	sectionCard := m.renderSectionCard(section, contentWidth)

	if stacked {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			sectionList,
			strings.Repeat("\n", 1),
			sectionCard,
		)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, sectionList, sectionCard)
}

func (m Model) renderSectionList(width int) string {
	base := m.style()
	labelStyle := base.
		Bold(true).
		Foreground(lipgloss.Color("#7DD3FC"))
	chipStyle := base.
		Padding(0, 1).
		MarginRight(1).
		Foreground(lipgloss.Color("#0F172A")).
		Background(lipgloss.Color("#38BDF8"))
	selectedChipStyle := chipStyle.Background(lipgloss.Color("#F59E0B"))
	panel := base.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#334155")).
		Padding(1, 2).
		Width(width)

	var b strings.Builder
	b.WriteString(labelStyle.Render("Sections"))
	b.WriteString("\n\n")
	for index, section := range m.sections {
		chip := fmt.Sprintf("%d", index+1)
		if index == m.selected {
			chip = selectedChipStyle.Render(chip)
		} else {
			chip = chipStyle.Render(chip)
		}
		line := fmt.Sprintf("%s %s", chip, section.title)
		if index == m.selected {
			line = base.
				Foreground(lipgloss.Color("#F8FAFC")).
				Bold(true).
				Render(line)
		}
		b.WriteString(line)
		b.WriteString("\n")
		b.WriteString(base.Foreground(lipgloss.Color("#94A3B8")).Render(section.summary))
		b.WriteString("\n\n")
	}

	b.WriteString(base.Foreground(lipgloss.Color("#94A3B8")).Render("Use h/j/k/l or arrow keys to move. Press q to quit."))

	return panel.Render(b.String())
}

func (m Model) renderSectionCard(section portfolioSection, width int) string {
	base := m.style()
	titleStyle := base.
		Bold(true).
		Foreground(lipgloss.Color("#F8FAFC")).
		MarginBottom(1)
	metaStyle := base.Foreground(lipgloss.Color("#94A3B8"))
	bodyStyle := base.Foreground(lipgloss.Color("#CBD5E1"))
	linkStyle := base.Foreground(lipgloss.Color("#38BDF8"))
	itemCard := base.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#334155")).
		Padding(1, 2)

	sectionTitle := titleStyle.Render(section.title)
	sectionDescription := metaStyle.Render(section.description)

	cardWidth := width - 4
	if cardWidth < 48 {
		cardWidth = width
	}
	itemCard = itemCard.Width(cardWidth)

	var cards []string
	for _, item := range section.items {
		lines := []string{base.Bold(true).Foreground(lipgloss.Color("#F8FAFC")).Render(item.title)}
		if item.meta != "" {
			lines = append(lines, metaStyle.Render(item.meta))
		}
		lines = append(lines, bodyStyle.Render(item.description))
		if item.linkURL != "" && item.linkText != "" {
			lines = append(lines, linkStyle.Render(createHyperlink(item.linkURL, item.linkText)))
		}
		cards = append(cards, itemCard.Render(strings.Join(lines, "\n\n")))
	}

	if len(cards) == 0 {
		cards = append(cards, itemCard.Render(metaStyle.Render("No content available yet.")))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, cards...)
	panel := base.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#334155")).
		Padding(1, 2).
		Width(width)

	return panel.Render(strings.Join([]string{sectionTitle, sectionDescription, "", content}, "\n"))
}

func (m Model) renderFooter(width int) string {
	base := m.style()
	footerStyle := base.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#334155")).
		Padding(0, 2).
		Width(width).
		Foreground(lipgloss.Color("#CBD5E1"))

	return footerStyle.Render(strings.Join([]string{
		"If the problem is interesting, I am usually in.",
		"Open to freelance work and collaborations.",
		createHyperlink("mailto:pvighnesh81203@gmail.com", "Email"),
		createHyperlink("https://github.com/nonvegetable", "GitHub"),
		createHyperlink("https://www.linkedin.com/in/vighnesh-palande/", "LinkedIn"),
		createHyperlink("https://drive.google.com/file/d/1Tzh-pc6sZeDywWG6QVOvirKNhLphOgwV/view?usp=sharing", "Resume"),
	}, "  •  "))
}

func portfolioSections() []portfolioSection {
	return []portfolioSection{
		aboutSection(),
		experienceSection(),
		projectsSection(),
		skillsSection(),
		hackathonSection(),
		otherSection(),
		contactSection(),
	}
}

func createHyperlink(url, text string) string {
	return "\x1b]8;;" + url + "\x1b\\" + text + "\x1b]8;;\x1b\\"
}
