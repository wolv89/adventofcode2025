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

type Line struct {
	start, end, anchor int
}

type Rect struct {
	c1, c2 Point
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
		p1, p2                 Point
		l1, l2, target         Line
		line                   string
		x, y, b, area          int
		width, height          int
		xmin, xmax, ymin, ymax int
	)

	points := make([]Point, 0)

	for scanner.Scan() {

		line = scanner.Text()

		fmt.Sscanf(line, "%d,%d", &x, &y)

		width = max(width, x)
		height = max(height, y)

		points = append(points, Point{x, y})

	}

	// Add first point to list at the end to fully wrap around
	points = append(points, points[0])
	n := len(points)

	verts, horis := make([]Line, 0, n/2), make([]Line, 0, n/2)

	for x = 1; x < n; x++ {

		p1, p2 = points[x-1], points[x]

		// Vertical
		if p1.x == p2.x {

			verts = append(verts, Line{
				start:  min(p1.y, p2.y),
				end:    max(p1.y, p2.y),
				anchor: p1.x,
			})

		} else {

			horis = append(horis, Line{
				start:  min(p1.x, p2.x),
				end:    max(p1.x, p2.x),
				anchor: p1.y,
			})

		}

	}

	lineCmp := func(a, b Line) int {
		return cmp.Compare(a.anchor, b.anchor)
	}
	slices.SortFunc(verts, lineCmp)
	slices.SortFunc(horis, lineCmp)

	vn, hn := len(verts), len(horis)

	// Remove the duplicated point from earlier
	points = points[:n-1]
	n--

	for x = 0; x < n-1; x++ {
	control: // Naming this as I just saw the trailer for Control Resonant (hyyyppee)
		for y = x + 1; y < n; y++ {

			p1, p2 = points[x], points[y]

			if p1.x == p2.x || p1.y == p2.y {
				continue
			}

			fmt.Println("")
			fmt.Println("RECTANGLE ", p1, p2)

			xmin = min(p1.x, p2.x)
			xmax = max(p1.x, p2.x)
			ymin = min(p1.y, p2.y)
			ymax = max(p1.y, p2.y)

			// Top, Bottom
			l1 = Line{start: xmin, end: xmax, anchor: ymin}
			l2 = Line{start: xmin, end: xmax, anchor: ymax}

			fmt.Println("\t", l1, l2)

			target = Line{0, 0, l1.start}
			b, _ = slices.BinarySearchFunc(verts, target, lineCmp)
			for ; b < vn; b++ {
				fmt.Println("\t\t", verts[b])
				if verts[b].anchor >= l1.end {
					break
				}
				if (verts[b].start < l1.anchor && verts[b].end > l1.anchor) || (verts[b].start < l2.anchor && verts[b].end > l2.anchor) {
					break control
				}
			}

			// Left, Right
			l1 = Line{start: ymin, end: ymax, anchor: xmin}
			l2 = Line{start: ymin, end: ymax, anchor: xmax}

			fmt.Println("\t", l1, l2)

			target = Line{0, 0, l1.start}
			b, _ = slices.BinarySearchFunc(horis, target, lineCmp)
			for ; b < hn; b++ {
				fmt.Println("\t\t", horis[b])
				if horis[b].anchor >= l1.end {
					break
				}
				if (horis[b].start < l1.anchor && horis[b].end > l1.anchor) || (horis[b].start < l2.anchor && horis[b].end > l2.anchor) {
					break control
				}
			}

			fmt.Println("\t>>> Valid - ", (abs(p1.y-p2.y)+1)*(abs(p1.x-p2.x)+1))
			area = max(area, (abs(p1.y-p2.y)+1)*(abs(p1.x-p2.x)+1))

		}
	}

	fmt.Println("")
	fmt.Println("Max area:", area)

}
