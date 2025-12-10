package day09

import (
	"bufio"
	"cmp"
	"fmt"
	"iter"
	"log"
	"os"
	"slices"
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

var DIRS = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type Rect struct {
	p1, p2 Point
	area   int
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
		rc                Rect
		p, p1, p2, p3, p4 Point
		line              string
		x, y, dx, dy      int
		width, height     int
	)

	points := make([]Point, 0)

	for scanner.Scan() {

		line = scanner.Text()

		fmt.Sscanf(line, "%d,%d", &x, &y)

		width = max(width, x)
		height = max(height, y)

		points = append(points, Point{x, y})

	}

	// Capture before adding the repeated point
	n := len(points)

	fmt.Println("1. Parsed points from TXT file", n)

	// Repeat point 1, to loop cleanly
	points = append(points, points[0])

	fmt.Println("Tack")

	height += 2
	width += 2

	fmt.Println("Expand", width, height)

	grid := make([][]byte, height)
	for x = 0; x < height; x++ {
		grid[x] = make([]byte, width)
		for y = 0; y < width; y++ {
			grid[x][y] = '.'
		}
	}

	fmt.Println("(1.5. Created blank grid)")

	for x = 1; x < len(points); x++ {

		fmt.Printf("\tDrawing from point %d to %d\n", x-1, x)

		for p = range WalkBetween(points[x-1], points[x]) {
			grid[p.y][p.x] = 'X'
		}

		grid[points[x-1].y][points[x-1].x] = '#'
		grid[points[x].y][points[x].x] = '#'

	}

	fmt.Println("2. Drawn polygon border")

	checked := make([][]bool, height)
	for x = range checked {
		checked[x] = make([]bool, width)
	}
	next := []Point{{0, 0}}

	for len(next) > 0 {

		x, y = next[0].x, next[0].y
		next = next[1:]

		if checked[x][y] {
			continue
		}

		checked[x][y] = true
		grid[x][y] = '~'

		for _, dir := range DIRS {

			dx, dy = x+dir[0], y+dir[1]

			if dx < 0 || dy < 0 || dx >= height || dy >= width || checked[dx][dy] || grid[dx][dy] != '.' {
				continue
			}

			next = append(next, Point{dx, dy})

		}

	}

	for x = range grid {
		for y = range grid[x] {
			switch grid[x][y] {
			case '.':
				grid[x][y] = 'X'
			case '~':
				grid[x][y] = '.'
			}
		}
	}

	fmt.Println("3. Flood fill complete")

	// Bit arbitrary, pre allocate space for half of the n by n-1 possibilities
	rects := make([]Rect, 0, (n*(n-1))/2)

	for x = 0; x < n-1; x++ {
		for y = x + 1; y < n; y++ {
			p1, p2 = points[x], points[y]
			if p1.x == p2.x || p1.y == p2.y {
				continue
			}
			rects = append(rects, Rect{
				p1:   p1,
				p2:   p2,
				area: (abs(p1.y-p2.y) + 1) * (abs(p1.x-p2.x) + 1),
			})
		}
	}

	slices.SortFunc(rects, func(a, b Rect) int {
		return cmp.Compare(b.area, a.area)
	})

	fmt.Println("4. Calculated possible rectangles", len(rects))
	var a int

outer:
	for a, rc = range rects {

		if a > 100 {
			break
		}

		// p1 and p2 on the rect are opposite corners
		p1, p3 = rc.p1, rc.p2
		p2 = Point{p1.x, p3.y}
		p4 = Point{p3.x, p1.y}

		for p = range WalkBetween(p1, p2) {
			if grid[p.y][p.x] == '.' {
				continue outer
			}
		}

		for p = range WalkBetween(p2, p3) {
			if grid[p.y][p.x] == '.' {
				continue outer
			}
		}

		for p = range WalkBetween(p3, p4) {
			if grid[p.y][p.x] == '.' {
				continue outer
			}
		}

		for p = range WalkBetween(p4, p1) {
			if grid[p.y][p.x] == '.' {
				continue outer
			}
		}

		break

	}

	if useSample > 0 {
		fmt.Println("")
		for x = range grid {
			fmt.Println(string(grid[x]))
		}
	}

	fmt.Println("")
	fmt.Println("Rect:", rc)

}
