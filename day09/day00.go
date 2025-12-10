package day09

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay9 struct{}

const DIR = "day09/"

type Point struct {
	x, y int
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

}
