package day09

import (
	"bufio"
	"cmp"
	"fmt"
	"iter"
	"log"
	"os"
)

type AocDay9 struct{}

const DIR = "day09/"

type Point struct {
	x, y int
}

// @Credit for this function: some assistance from ChatGTP
func WalkBetween(start, end Point) iter.Seq[Point] {
	return func(yield func(Point) bool) {

		dx := cmp.Compare(end.x, start.x)
		dy := cmp.Compare(end.y, start.y)

		// Offsets so we range BETWEEN the points, but not include them
		x, y := start.x+dx, start.y+dy

		end.x -= dx
		end.y -= dy

		// Loop until we've reached the end point
		for {
			if !yield(Point{x, y}) {
				return
			}
			if x == end.x && y == end.y {
				return
			}
			x += dx
			y += dy
		}
	}
}

func (d AocDay9) Puzzle1(useSample int) {

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
		p1, p2     Point
		line       string
		x, y, area int
	)

	points := make([]Point, 0)

	for scanner.Scan() {

		line = scanner.Text()

		fmt.Sscanf(line, "%d,%d", &x, &y)

		points = append(points, Point{x, y})

	}

	n := len(points)

	for x = 0; x < n-1; x++ {
		for y = x + 1; y < n; y++ {
			p1, p2 = points[x], points[y]
			area = max(area, (abs(p1.y-p2.y)+1)*(abs(p1.x-p2.x)+1))
		}
	}

	fmt.Println("")
	fmt.Println("Max area:", area)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d AocDay9) Puzzle2(useSample int) {

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
		line          string
		x, y          int
		width, height int
	)

	// Hard coded from generating an SVG render of the shape
	// and figuring out which point would be the most likely anchor
	// and working off that (limited by where it would start leaving the bounds)
	// Holy heck...
	anchor := Point{94880, 50218}
	lim := 69776

	points := make([]Point, 0)

	for scanner.Scan() {

		line = scanner.Text()

		fmt.Sscanf(line, "%d,%d", &x, &y)

		width = max(width, x)
		height = max(height, y)

		p := Point{x, y}

		if p == anchor {
			break
		}

		points = append(points, p)

	}

	area := 0

	for i := len(points) - 1; i >= 0; i-- {

		p := points[i]
		if p.y > lim {
			break
		}

		area = max(area, (abs(anchor.y-p.y)+1)*(abs(anchor.x-p.x)+1))
		// fmt.Println(i, area)

	}

	fmt.Println("")
	fmt.Println("Max area:", area)

}
