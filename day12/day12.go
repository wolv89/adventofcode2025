package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay12 struct{}

const DIR = "day12/"

/* Flatly represents a 3x3 grid */
type Shape [9]bool

/* Pre-calculated moves to rotate 90 degrees */
var rotate90 = [9]int{2, 4, 6, -2, 0, 2, -6, -4, -2}

// Rotate 90 degrees clockwise
func (s Shape) Rotate() Shape {

	shape := Shape{}

	for i := range 9 {
		shape[i+rotate90[i]] = s[i]
	}

	return shape

}

// Swap left and right columns
func (s Shape) Flip() Shape {

	shape := s

	shape[0], shape[2] = shape[2], shape[0]
	shape[3], shape[5] = shape[5], shape[3]
	shape[6], shape[8] = shape[8], shape[6]

	return shape

}

func (s Shape) Render() string {
	var b strings.Builder
	b.Grow(11)
	for i := range 9 {
		if i > 0 && i%3 == 0 {
			b.WriteString("\n")
		}
		if s[i] {
			b.WriteByte('#')
		} else {
			b.WriteByte('.')
		}
	}
	return b.String()
}

type Region struct {
	shapes        []int
	width, height int
}

func (d AocDay12) Puzzle1(useSample int) {

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
		sh           Shape
		line         string
		c            rune
		x, y, n      int
		readingShape bool
	)

	rawShapes := make([]Shape, 0)
	regions := make([]Region, 0)

	for scanner.Scan() {

		line = strings.TrimSuffix(scanner.Text(), "\n")

		if readingShape {

			if line == "" {
				readingShape = false
				rawShapes = append(rawShapes, sh)
				continue
			}

			for _, c = range line {
				if c == '#' {
					sh[y] = true
				}
				y++
			}

		} else {

			_, err = fmt.Sscanf(line, "%d:", &x)

			if err == nil {
				readingShape = true
				sh = Shape{}
				y = 0
				continue
			}

			parts := strings.Split(line, ":")

			fmt.Sscanf(parts[0], "%dx%d", &x, &y)
			shapeids := make([]int, 0)

			counts := strings.SplitSeq(strings.TrimSpace(parts[1]), " ")
			for ct := range counts {
				n, _ = strconv.Atoi(ct)
				shapeids = append(shapeids, n)
			}

			regions = append(regions, Region{
				shapes: shapeids,
				width:  x,
				height: y,
			})

		}

	}

	shapes := make([][]Shape, 0, len(rawShapes))

	for _, sh := range rawShapes {

		set := make(map[Shape]struct{})

		set[sh] = struct{}{}

		for range 3 {
			sh = sh.Rotate()
			set[sh] = struct{}{}
		}

		sh = sh.Flip()
		set[sh] = struct{}{}

		for range 3 {
			sh = sh.Rotate()
			set[sh] = struct{}{}
		}

		variations := make([]Shape, 0, len(set))
		for sh = range set {
			variations = append(variations, sh)
		}

		shapes = append(shapes, variations)

	}

	/*
		fmt.Println("")

		for _, vars := range shapes {
			for _, sh := range vars {
				fmt.Println(sh.Render())
				fmt.Println("")
			}
			fmt.Println("+---------------------------+")
			fmt.Println("")
		}
	*/

	fmt.Println("")

	fits := 0

	for _, r := range regions {

		count := 0
		for _, ct := range r.shapes {
			count += ct
		}

		count *= 9

		easy := r.width * r.height

		fmt.Println(r, " | ", count, " | ", easy)
		if count <= easy {
			fits++
		}

		// if lim > 5 {
		// 	break
		// }

	}

	fmt.Println("")
	fmt.Println("Easy fits:", fits)

}

func (d AocDay12) Puzzle2(useSample int) {

}
