package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const SPLASH_SCREEN_FRAME_COUNT = 80

var scrollViewTotalLines = 0
var scrollViewUsableHeight = 0

const (
	splashView uint = iota
	homeView
	aboutView
	experienceView
	projectsView
	skillsView
	contactView
)

type Theme struct {
	foreground      lipgloss.Color
	foregroundMuted lipgloss.Color
	primary         lipgloss.Color
	secondary       lipgloss.Color
}

type Model struct {
	renderer *lipgloss.Renderer

	scrollOffset int

	currentView uint
	height      int
	width       int

	frame int

	theme Theme
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func NewModel(renderer *lipgloss.Renderer) Model {
	return Model{
		renderer:    renderer,
		currentView: splashView,
		theme: Theme{
			foreground:      lipgloss.Color("255"),
			foregroundMuted: lipgloss.Color("244"),
			primary:         lipgloss.Color("#89CFF0"),
			secondary:       lipgloss.Color("45"),
		},
	}
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tickMsg:
		m.frame++
		if m.frame > 100 {
			m.frame = 0
		}
		if m.currentView == splashView && m.frame > SPLASH_SCREEN_FRAME_COUNT {
			m.currentView = homeView
		}
		return m, tick()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.currentView != contactView {
				return m, tea.Quit
			}
		case "right":
			if m.currentView < contactView {
				m.currentView++
				m.scrollOffset = 0
			}
		case "left":
			if m.currentView > homeView {
				m.currentView--
				m.scrollOffset = 0
			}
		case "up":
			if (m.currentView == aboutView || m.currentView == projectsView || m.currentView == skillsView || m.currentView == experienceView) && m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "down":
			if (m.currentView == aboutView || m.currentView == projectsView || m.currentView == skillsView || m.currentView == experienceView) && m.scrollOffset < scrollViewTotalLines-scrollViewUsableHeight {
				m.scrollOffset++
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, tea.Batch(cmds...)
}
