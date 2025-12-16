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

type Edge struct {
	a, b Point
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
		p1, p2        Point
		line          string
		x, y          int
		width, height int
	)

	points := make([]Point, 0)
	pointMap := make(map[Point]struct{})

	for scanner.Scan() {

		line = scanner.Text()

		fmt.Sscanf(line, "%d,%d", &x, &y)

		width = max(width, x)
		height = max(height, y)

		p1 = Point{x, y}
		points = append(points, p1)
		pointMap[p1] = struct{}{}

	}

	points = append(points, points[0])
	n := len(points)

	ver, hor := make([]Edge, 0, n/2), make([]Edge, 0, n/2)
	for x = 1; x < n; x++ {
		p1, p2 = points[x-1], points[x]
		if p1.x == p2.x {
			if p1.y > p2.y {
				p1, p2 = p2, p1
			}
			ver = append(ver, Edge{p1, p2})
		} else {
			if p1.x > p2.x {
				p1, p2 = p2, p1
			}
			hor = append(hor, Edge{p1, p2})
		}
	}

	slices.SortFunc(ver, func(e, f Edge) int {
		return cmp.Compare(e.a.x, f.a.x)
	})
	slices.SortFunc(hor, func(e, f Edge) int {
		return cmp.Compare(e.a.y, f.a.y)
	})

	if useSample > 0 {
		fmt.Println(ver)
		fmt.Println(hor)
		fmt.Println("")
	}

	n--
	points = points[:n]

	getValidArea := func(p1, p2 Point) int {

		xmin, xmax := min(p1.x, p2.x), max(p1.x, p2.x)
		ymin, ymax := min(p1.y, p2.y), max(p1.y, p2.y)

		corners := [4]Point{
			{xmin, ymin},
			{xmin, ymax},
			{xmax, ymax},
			{xmax, ymin},
		}

		if useSample > 0 {
			fmt.Println("")
			fmt.Println("\t", corners)
		}

		// rect := [4]Edge{
		// 	{corners[0], corners[1]},
		// 	{corners[1], corners[2]},
		// 	{corners[2], corners[3]},
		// 	{corners[3], corners[0]},
		// }

	cornercheck:
		for _, c := range corners {

			if _, ok := pointMap[c]; ok {
				if useSample > 0 {
					fmt.Println("\t\t", c, " is an exact point")
				}
				continue
			}

			pointEdge := Edge{c, c}

			// Search vertical edges, to see if point sits on any of them
			i, found := slices.BinarySearchFunc(ver, pointEdge, func(e, f Edge) int {
				return cmp.Compare(e.a.x, f.a.x)
			})
			if found {
				for ; i < len(ver); i++ {
					if ver[i].a.x != c.x {
						break
					}
					if ver[i].a.y <= c.y && ver[i].b.y >= c.y {
						if useSample > 0 {
							fmt.Println("\t\t", c, "on edge", ver[i])
						}
						continue cornercheck
					}
				}
			}

			// If not, try horizontal edges
			i, found = slices.BinarySearchFunc(hor, pointEdge, func(e, f Edge) int {
				return cmp.Compare(e.a.y, f.a.y)
			})
			if found {
				for ; i < len(hor); i++ {
					if hor[i].a.y != c.y {
						break
					}
					if hor[i].a.x <= c.x && hor[i].b.x >= c.x {
						if useSample > 0 {
							fmt.Println("\t\t", c, "on edge", hor[i])
						}
						continue cornercheck
					}
				}
			}

			if useSample > 0 {
				fmt.Println("\t\t sweeping")
			}

			// If not on an edge, then we sweep across lines until we hit the point, toggling our status from outside to in
			inside := false

			for _, e := range ver {

				if e.a.y <= c.y && e.b.y > c.y {
					inside = !inside
				}
				if c.x >= e.a.x {
					break
				}

			}

			if !inside {
				return 0
			}

		}

		return (abs(p1.y-p2.y) + 1) * (abs(p1.x-p2.x) + 1)

	}

	var (
		pa, pb Point
		area   int
	)

	for x = 0; x < n-1; x++ {

		pa = points[x]

		for y = x + 1; y < n; y++ {

			pb = points[y]

			if pa.x == pb.x || pa.y == pb.y {
				continue
			}

			area = max(area, getValidArea(pa, pb))

		}

	}

	fmt.Println("")
	fmt.Println("Max area:", area)

}
