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

	var names []string

	if useSample > 0 {

		names = make([]string, i)
		for dname, did := range devices {
			names[did] = fmt.Sprintf("%s", dname)
		}

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

	indStart := make([]int, i)

	for _, nodes := range edges {
		for _, node := range nodes {
			indStart[node]++
		}
	}

	/*
	 * This reads funny, but we always walk the full path from the starting node (which is svr, not the "start" param)
	 * We only start counting paths once we reach the "start" param, but need to walk from the start to drop the in-degree counts
	 * Perhaps there's a better way?
	 * Ok a bit more reading and yes indeed there are much better ways, like doing the walk just once and generating a proper sorting of the nodes
	 * to then count the paths between them more quickly (although this runs pretty dang fast as is)
	 */
	countPaths := func(start, end int) int64 {

		ind := make([]int, i)
		copy(ind, indStart)

		paths := make([]int64, i)
		var node, e int

		q := []int{svr}
		paths[svr]++

		started := false

		for len(q) > 0 {

			node, q = q[0], q[1:]
			if node >= len(edges) {
				continue
			}

			if node == start {
				started = true
			}

			for _, e = range edges[node] {

				if started {
					if node == start {
						paths[e] = 1
					} else {
						paths[e] += paths[node]
					}
				}

				ind[e]--
				if ind[e] == 0 {
					q = append(q, e)
				}

			}

		}

		return paths[end]

	}

	s2d := countPaths(svr, dac)
	d2f := countPaths(dac, fft)
	f2o := countPaths(fft, out)

	s2f := countPaths(svr, fft)
	f2d := countPaths(fft, dac)
	d2o := countPaths(dac, out)

	sum1 := s2d * d2f * f2o
	sum2 := s2f * f2d * d2o

	fmt.Printf("%d * %d * %d = %d\n", s2d, d2f, f2o, sum1)
	fmt.Printf("%d * %d * %d = %d\n", s2f, f2d, d2o, sum2)

	fmt.Println("")
	fmt.Println("Total paths:", sum1+sum2)

}

/*
func countPaths(adj map[string][]string, start, end string) int {
	inDegree := make(map[string]int)
	for u := range adj {
		for _, v := range adj[u] {
			inDegree[v]++
		}
	}

	// paths[node] stores number of ways to reach 'node' from 'start'
	paths := make(map[string]int)
	paths[start] = 1

	queue := []string{}
	// In a true Topo Sort, you'd add all 0-in-degree nodes.
	// For path counting from a specific start, we start there.
	for node := range adj {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range adj[u] {
			paths[v] += paths[u]
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return paths[end]
}
*/
