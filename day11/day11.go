package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type AocDay11 struct{}

const DIR = "day11/"

type deviceID [3]byte

func (d AocDay11) Puzzle1(useSample int) {

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

	devices := make(map[deviceID]int)
	edges := make([][]int, 0)
	i := 0

	var start, end int

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, " ")
		parts[0] = strings.Trim(parts[0], ":")

		ids := make([]int, 0, len(parts))

		for _, part := range parts {

			var id deviceID
			copy(id[:], part)

			if _, ok := devices[id]; !ok {
				devices[id] = i
				i++
			}

			switch part {
			case "you":
				start = devices[id]
			case "out":
				end = devices[id]
			}

			ids = append(ids, devices[id])

		}

		if len(edges) <= ids[0] {
			edges = append(edges, make([][]int, 1+ids[0]-len(edges))...)
		}

		for c := 1; c < len(ids); c++ {
			edges[ids[0]] = append(edges[ids[0]], ids[c])
		}

	}

	if useSample > 0 {

		for dname, did := range devices {
			fmt.Printf("%s: %d\n", dname, did)
		}

		fmt.Println("")

		for e, edge := range edges {
			if len(edge) == 0 {
				continue
			}
			fmt.Println(e, " > ", edge)
		}

		fmt.Println("")
		fmt.Println("Start:", start, " End:", end)
		fmt.Println("")

	}

	steps := make([]int, i)
	q := []int{start}
	var node int

	for len(q) > 0 {

		node, q = q[0], q[1:]

		for _, e := range edges[node] {
			steps[e]++
			if e != end {
				q = append(q, e)
			}
		}

	}

	if useSample > 0 {
		for x := range steps {
			fmt.Println(x, steps[x])
		}
	}

	fmt.Println("")
	fmt.Println("Paths:", steps[end])

}

func (d AocDay11) Puzzle2(useSample int) {

}
