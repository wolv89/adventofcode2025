package day07

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	styleDefault        = lipgloss.NewStyle().Foreground(lipgloss.Color("#447744"))
	styleActive         = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
	styleSplitter       = lipgloss.NewStyle().Foreground(lipgloss.Color("#aa9900"))
	styleSplitterActive = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff00"))
)

type updateMsg struct{}

func listenForUpdate(sub chan struct{}, diagram *Diagram) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * 50)
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
	sub          chan struct{}
	diagram      *Diagram
	rawMode      bool
	timelineMode bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForUpdate(m.sub, m.diagram),
		waitForUpdate(m.sub),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "c", "q":
			return m, tea.Quit
		case "m":
			m.rawMode = !m.rawMode
			return m, nil
		case "t":
			m.timelineMode = !m.timelineMode
			return m, nil
		}
	case updateMsg:
		return m, waitForUpdate(m.sub)
	}
	return m, nil
}

func (m model) View() string {

	head := fmt.Sprintf("Beams: %d\nSplitters Activated: %d", m.diagram.beamCount, m.diagram.activeSplitters)
	tail := ""
	if len(m.diagram.timeline) > 0 {
		tail = fmt.Sprintf("Timelines: %d", m.diagram.timelineCount)
	}

	var s strings.Builder
	s.Grow(m.diagram.height * (m.diagram.width + 1))

	if m.rawMode {
		for _, row := range m.diagram.state {
			s.WriteString(string(row) + "\n")
		}
		return head + "\n\n" + s.String() + "\n\n" + head + "\n" + tail
	} else if m.timelineMode {
		for _, row := range m.diagram.timeline {
			for _, t := range row {
				if t == 0 {
					s.WriteString(styleDefault.Render(". "))
				} else {
					s.WriteString(fmt.Sprintf("%d ", t))
				}
			}
			s.WriteString("\n")
		}
		return head + "\n\n" + s.String() + "\n\n" + head + "\n" + tail
	}

	var (
		spaceStyle lipgloss.Style
		row        []byte
		c          byte
		i          int
	)

	for i, row = range m.diagram.state {

		spaceStyle = styleDefault
		if i == m.diagram.row-1 {
			spaceStyle = styleActive
		}

		for _, c = range row {

			switch c {
			case CHAR_SPACE:
				s.WriteString(spaceStyle.Render(STR_SPACE))
			case CHAR_SPLITTER:
				s.WriteString(styleSplitter.Render(STR_SPLITTER))
			case CHAR_SPLITTER_ACTIVE:
				s.WriteString(styleSplitterActive.Render(STR_SPLITTER))
			case CHAR_START:
				s.WriteString(styleSplitterActive.Render(STR_START))
			default:
				s.WriteString(beamColours[int(c-'0')].Render(STR_BEAM))
			}

		}

		s.WriteString("\n")

	}

	return head + "\n\n" + s.String() + "\n\n" + head + "\n" + tail

}
