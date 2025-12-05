package day04

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type AocDay4 struct{}

const DIR = "day04/"

func (d AocDay4) Puzzle1(useSample int) {

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

	var (
		line                           string
		c                              rune
		width, height, x, y, freeRolls int
		adjRolls                       uint8
	)

	// Preloop to find full size of grid

	for scanner.Scan() {

		if width == 0 {
			line = scanner.Text()
			width = len(line)
		}

		height++

	}

	// Build grid with gutter

	grid := make([][]uint8, height+2)

	for x = range grid {
		grid[x] = make([]uint8, width+2)
	}

	// Read again to fill grid

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner2 := bufio.NewScanner(f)
	scanner2.Split(bufio.ScanLines)

	x = 1

	for scanner2.Scan() {

		line = scanner2.Text()
		y = 1

		for _, c = range line {
			if c == '@' {
				grid[x][y] = 1
			}
			y++
		}

		x++

	}

	if useSample > 0 {
		for x = 0; x < height+2; x++ {
			fmt.Println(grid[x])
		}
	}

	// Count adjacent rolls

	for x = 1; x < height+1; x++ {

		for y = 1; y < width+1; y++ {

			if grid[x][y] != 1 {
				continue
			}

			adjRolls = grid[x-1][y-1] + grid[x-1][y] + grid[x-1][y+1] + grid[x][y-1] + grid[x][y+1] + grid[x+1][y-1] + grid[x+1][y] + grid[x+1][y+1]

			if adjRolls < 4 {
				freeRolls++
			}

		}

	}

	fmt.Println("")
	fmt.Println("Free rolls:", freeRolls)

}

func (d AocDay4) Puzzle2(useSample int) {

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

	var (
		line                                     string
		c                                        rune
		width, height, x, y, freeRolls, lastFree int
		adjRolls                                 uint8
	)

	// Preloop to find full size of grid

	for scanner.Scan() {

		if width == 0 {
			line = scanner.Text()
			width = len(line)
		}

		height++

	}

	// Build grid with gutter

	grid := make([][]uint8, height+2)

	for x = range grid {
		grid[x] = make([]uint8, width+2)
	}

	// Read again to fill grid

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner2 := bufio.NewScanner(f)
	scanner2.Split(bufio.ScanLines)

	x = 1

	for scanner2.Scan() {

		line = scanner2.Text()
		y = 1

		for _, c = range line {
			if c == '@' {
				grid[x][y] = 1
			}
			y++
		}

		x++

	}

	// Count adjacent rolls

	lastFree = -1

	for lastFree != freeRolls {

		lastFree = freeRolls

		for x = 1; x < height+1; x++ {

			for y = 1; y < width+1; y++ {

				if grid[x][y] != 1 {
					continue
				}

				adjRolls = grid[x-1][y-1] + grid[x-1][y] + grid[x-1][y+1] + grid[x][y-1] + grid[x][y+1] + grid[x+1][y-1] + grid[x+1][y] + grid[x+1][y+1]

				if adjRolls < 4 {
					freeRolls++

					// Free me
					grid[x][y] = 0
				}

			}

		}

	}

	fmt.Println("")
	fmt.Println("Free rolls:", freeRolls)

}
