package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AocDay7 struct{}

const DIR = "day07/"

const (
	CHAR_START           = 'S'
	CHAR_SPACE           = '.'
	CHAR_SPLITTER        = '^'
	CHAR_SPLITTER_ACTIVE = '#'
	CHAR_BEAM            = '|'

	// In theory saving on conversions later?
	STR_START    = "S"
	STR_SPACE    = "."
	STR_SPLITTER = "^"
	STR_BEAM     = "|"
)

// This colour palette was AI generated
var beamColours = []lipgloss.Style{
	lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")),
	lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF")),
	lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")),
	lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")),
	lipgloss.NewStyle().Foreground(lipgloss.Color("#BC13FE")),
	lipgloss.NewStyle().Foreground(lipgloss.Color("#0165FC")),
	lipgloss.NewStyle().Foreground(lipgloss.Color("#FF7124")),
	lipgloss.NewStyle().Foreground(lipgloss.Color("#FF028D")),
}

var beamChars = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

type Diagram struct {
	state               [][]byte
	timeline            [][]int // For puzzle 2
	width, height, row  int
	beamCount, lastBeam int
	activeSplitters     int
	timelineCount       int
}

func isBeam(c byte) bool {
	return c >= '0' && c <= '9'
}

func (d *Diagram) NewBeam() byte {
	d.beamCount++
	d.lastBeam = (d.lastBeam + 1) % len(beamColours)
	return beamChars[d.lastBeam]
}

func (d *Diagram) Update() {

	beamColours = append(beamColours[1:], beamColours[0])

	if d.row >= d.height {
		if len(d.timeline) > 0 && d.timelineCount == 0 {
			for _, tl := range d.timeline[d.height-1] {
				d.timelineCount += tl
			}
		}
		return
	}

	var c byte

	for x := range d.state[d.row] {

		c = d.state[d.row][x]

		if d.state[d.row][x] == CHAR_SPACE {

			c = CHAR_SPACE

			if isBeam(d.state[d.row-1][x]) {
				c = d.state[d.row-1][x]
				if len(d.timeline) > 0 {
					d.timeline[d.row][x] += d.timeline[d.row-1][x]
				}
			} else if d.state[d.row-1][x] == CHAR_START {
				if len(d.timeline) > 0 {
					d.timeline[d.row][x] = 1
				}
				c = d.NewBeam()
			}

		} else if d.state[d.row][x] == CHAR_SPLITTER {

			c = CHAR_SPLITTER

			if isBeam(d.state[d.row-1][x]) {
				c = CHAR_SPLITTER_ACTIVE
				d.activeSplitters++

				if x > 0 {
					if d.state[d.row][x-1] == CHAR_SPACE {
						d.state[d.row][x-1] = d.NewBeam()
					}
					if len(d.timeline) > 0 {
						d.timeline[d.row][x-1] += d.timeline[d.row-1][x]
					}
				}
				if x < d.width-1 {
					if d.state[d.row][x+1] == CHAR_SPACE {
						d.state[d.row][x+1] = d.NewBeam()
					}
					if len(d.timeline) > 0 {
						d.timeline[d.row][x+1] += d.timeline[d.row-1][x]
					}
				}
			}

		} else if isBeam(d.state[d.row][x]) {
			if len(d.timeline) > 0 {
				d.timeline[d.row][x] += d.timeline[d.row-1][x]
			}
		}

		d.state[d.row][x] = c

	}

	d.row++

}

func (d AocDay7) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	} else if useSample == 2 {
		datafile = DIR + "sample2.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	startingState := make([][]byte, 0)

	for scanner.Scan() {
		b := make([]byte, len(scanner.Bytes()))
		copy(b, scanner.Bytes())
		startingState = append(startingState, b)
	}

	diag := Diagram{
		state:     startingState,
		width:     len(startingState[0]),
		height:    len(startingState),
		row:       1, // Start at 2nd row to always read off row above
		beamCount: 0,
		lastBeam:  -1,
	}

	p := tea.NewProgram(model{
		sub:     make(chan struct{}),
		diagram: &diag,
		rawMode: false,
	}, tea.WithAltScreen())

	if _, err = p.Run(); err != nil {
		fmt.Println("Quitting...")
		// os.Exit(1)
	}

}

func (d AocDay7) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	startingState := make([][]byte, 0)

	for scanner.Scan() {
		b := make([]byte, len(scanner.Bytes()))
		copy(b, scanner.Bytes())
		startingState = append(startingState, b)
	}

	w, h := len(startingState[0]), len(startingState)
	tline := make([][]int, h)
	for t := range tline {
		tline[t] = make([]int, w)
	}

	diag := Diagram{
		state:     startingState,
		width:     w,
		height:    h,
		timeline:  tline,
		row:       1, // Start at 2nd row to always read off row above
		beamCount: 0,
		lastBeam:  -1,
	}

	p := tea.NewProgram(model{
		sub:     make(chan struct{}),
		diagram: &diag,
		rawMode: false,
	}, tea.WithAltScreen())

	if _, err = p.Run(); err != nil {
		fmt.Println("Quitting...")
		// os.Exit(1)
	}

}
