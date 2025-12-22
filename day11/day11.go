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
			case "svr":
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

type seenEdge [2]int

func (d AocDay11) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample2.txt" // Sample 1 irrelevant here
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

	var svr, fft, dac, out int

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
			case "svr":
				svr = devices[id]
			case "fft":
				fft = devices[id]
			case "dac":
				dac = devices[id]
			case "out":
				out = devices[id]
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

	names := make([]string, i)
	for dname, did := range devices {
		names[did] = fmt.Sprintf("%s", dname)
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

	}

	fmt.Println("Nodes:", i)
	fmt.Println("")
	fmt.Println("svr:", svr, " out:", out)
	fmt.Println("fft:", fft, " dac:", dac)
	fmt.Println("")

	stepsFrom := func(start int, log ...bool) []int {

		seen := make(map[seenEdge]struct{})

		steps := make([]int, i)
		q := []int{start}

		var (
			qued []int
			node int
			ok   bool
		)

		for len(q) > 0 {

			node, q = q[0], q[1:]

			if len(log) > 0 && log[0] {
				qued = make([]int, 0)
			}

			for _, e := range edges[node] {
				steps[e]++
				if _, ok = seen[seenEdge{node, e}]; !ok {
					q = append(q, e)
					seen[seenEdge{node, e}] = struct{}{}
					if len(log) > 0 && log[0] {
						qued = append(qued, e)
					}
				}
			}

			if len(log) > 0 && log[0] {
				fmt.Printf("Node: %s (%d)\n", names[node], node)
				if len(qued) > 0 {
					fmt.Print("\t")
					for _, qn := range qued {
						fmt.Printf("%s (%d) ", names[qn], qn)
					}
					fmt.Print("\n")
				}
			}

		}

		return steps

	}

	validPaths := 0

	svr2dac := stepsFrom(svr)       // map[int]struct{}{fft: {}, dac: {}, out: {}}
	dac2fft := stepsFrom(dac, true) // map[int]struct{}{fft: {}, out: {}}
	fft2out := stepsFrom(fft)       // map[int]struct{}{out: {}, dac: {}}

	if useSample > 0 {
		fmt.Println(svr2dac, svr2dac[dac])
		fmt.Println(dac2fft, dac2fft[fft])
		fmt.Println(fft2out, fft2out[out])
		fmt.Println("")
	} else {
		fmt.Println(svr2dac[dac])
		fmt.Println(dac2fft[fft])
		fmt.Println(fft2out[out])
		fmt.Println("")
	}

	validPaths += svr2dac[dac] * dac2fft[fft] * fft2out[out]

	// This now seems to just be repeated work...
	svr2fft := stepsFrom(svr)
	fft2dac := stepsFrom(fft)
	dac2out := stepsFrom(dac)

	if useSample > 0 {
		fmt.Println(svr2fft, svr2fft[fft])
		fmt.Println(fft2dac, fft2dac[dac])
		fmt.Println(dac2out, dac2out[out])
		fmt.Println("")
	} else {
		fmt.Println(svr2fft[fft])
		fmt.Println(fft2dac[dac])
		fmt.Println(dac2out[out])
		fmt.Println("")
	}

	validPaths += svr2fft[fft] * fft2dac[dac] * dac2out[out]

	fmt.Println("")
	fmt.Println("Valid paths:", validPaths)

}
