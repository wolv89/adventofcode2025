package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type AocDay7 struct{}

const DIR = "day07/"

const (
	CHAR_START           = 'S'
	CHAR_SPACE           = '.'
	CHAR_SPLITTER        = '^'
	CHAR_SPLITTER_ACTIVE = '#'
	CHAR_BEAM            = '|'
)

type Diagram struct {
	state              [][]byte
	width, height, row int
}

func isBeam(c byte) bool {
	return c == CHAR_BEAM
}

func (d *Diagram) Update() {

	if d.row >= d.height {
		return
	}

	var c byte

	for x := range d.state[d.row] {

		c = d.state[d.row][x]

		if d.state[d.row][x] == CHAR_SPACE {

			c = CHAR_SPACE

			if isBeam(d.state[d.row-1][x]) || d.state[d.row-1][x] == CHAR_START {
				c = CHAR_BEAM
			}

		} else if d.state[d.row][x] == CHAR_SPLITTER {

			c = CHAR_SPLITTER

			if isBeam(d.state[d.row-1][x]) {
				c = CHAR_SPLITTER_ACTIVE

				if x > 0 && d.state[d.row][x-1] == CHAR_SPACE {
					d.state[d.row][x-1] = CHAR_BEAM
				}
				if x < d.width-1 && d.state[d.row][x+1] == CHAR_SPACE {
					d.state[d.row][x+1] = CHAR_BEAM
				}
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
		startingState = append(startingState, scanner.Bytes())
	}

	diag := Diagram{
		state:  startingState,
		width:  len(startingState[0]),
		height: len(startingState),
		row:    1, // Start at 2nd row to always read off row above
	}

	p := tea.NewProgram(model{
		sub:     make(chan struct{}),
		diagram: &diag,
	}, tea.WithAltScreen())

	if _, err = p.Run(); err != nil {
		fmt.Println("Quitting...")
		os.Exit(1)
	}

}

func (d AocDay7) Puzzle2(useSample int) {

}
