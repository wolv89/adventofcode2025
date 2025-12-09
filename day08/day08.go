package day08

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type AocDay8 struct{}

const DIR = "day08/"

/*
 * Bit of a lazy way to carry this around the package
 */
var HeapLimit int

type Point struct {
	x, y, z int
}

type Pair struct {
	dist   int64
	p1, p2 Point
}

func (d AocDay8) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	HeapLimit = 1000

	if useSample == 1 {
		datafile = DIR + "sample.txt"
		HeapLimit = 10
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		pair           Pair
		line           string
		dist           int64
		x, y, z, c, c2 int
	)

	// Simple slice of circuits, where the circuit is identified by its index, and the value is the count of boxes within
	// Starting with length one to leave the 0th index empty
	circuits := make([]int, 1)

	// Lookup the index of the circuit a given point is in
	circuitPoint := make(map[Point]int)

	// Need a way to look up all the points in a circuit when they get joined/re-assigned
	pointLists := make([][]Point, 1)

	// Each box (point) is going to start in its own circuit
	c = 1

	points := make([]Point, 0)

	for scanner.Scan() {

		line = scanner.Text()
		coords := strings.Split(line, ",")

		x, err = strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err.Error())
		}

		y, err = strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		z, err = strconv.Atoi(coords[2])
		if err != nil {
			log.Fatal(err.Error())
		}

		p := Point{x, y, z}

		points = append(points, p)
		circuits = append(circuits, 1)
		pointLists = append(pointLists, []Point{p})
		circuitPoint[p] = c
		c++

	}

	pairs := make(PairHeap, 0, HeapLimit)

	heap.Init(&pairs)

	for x = 0; x < len(points)-1; x++ {
		for y = x + 1; y < len(points); y++ {

			p1, p2 := points[x], points[y]

			dist = int64(math.Pow(float64(p1.x-p2.x), 2)) + int64(math.Pow(float64(p1.y-p2.y), 2)) + int64(math.Pow(float64(p1.z-p2.z), 2))

			pair = Pair{
				dist: dist,
				p1:   p1,
				p2:   p2,
			}

			heap.Push(&pairs, pair)

		}
	}

	for l := 0; pairs.Len() > 0 && l < HeapLimit; l++ {

		pair = heap.Pop(&pairs).(Pair)

		if useSample > 0 {
			fmt.Println("")
			fmt.Println(pair.p1, pair.p2, " | ", pair.dist)
			fmt.Println("@", circuitPoint[pair.p1], " | @", circuitPoint[pair.p2])
		}

		c = circuitPoint[pair.p1]
		c2 = circuitPoint[pair.p2]

		if c != c2 {
			circuits[c] += circuits[c2]
			circuits[c2] = 0

			// This worked, whoop!
			// Using a slice, not a set means I could end up with dupes of points within a given circuit, so this is not a perfect design
			// But worked well enough in this case, and I guess my separate counter (the circuits slice) is less fallible to this duplication
			for _, p := range pointLists[c2] {
				circuitPoint[p] = c
				pointLists[c] = append(pointLists[c], p)
			}
		}

	}

	slices.Sort(circuits)

	if useSample > 0 {
		fmt.Println("")
		for i, circ := range circuits {
			fmt.Println(i, circ)
		}
	}

	n := len(circuits) - 1
	res := circuits[n] * circuits[n-1] * circuits[n-2]

	fmt.Println("")
	fmt.Println("Result:", res)

	// Attempt 1: 7106 (Too low?)

}

func (d AocDay8) Puzzle2(useSample int) {

}
