package day07

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	baubleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e73007"))
	treeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#16df16"))
)

type updateMsg struct{}

func listenForUpdate(sub chan struct{}, diagram *Diagram) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * 500)
			diagram.Update()
			sub <- struct{}{}
		}
	}
}

func waitForUpdate(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return updateMsg(<-sub)
	}
}

type model struct {
	sub     chan struct{}
	diagram *Diagram
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForUpdate(m.sub, m.diagram),
		waitForUpdate(m.sub),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case updateMsg:
		return m, waitForUpdate(m.sub)
	default:
		return m, nil
	}
}

func (m model) View() string {

	var s strings.Builder
	s.Grow(m.diagram.height * (m.diagram.width + 1))

	for i, row := range m.diagram.state {
		if i%2 == 0 {
			s.WriteString(baubleStyle.Render(string(row)) + "\n")
		} else {
			s.WriteString(treeStyle.Render(string(row)) + "\n")
		}
	}

	return s.String()

}
