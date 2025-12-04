package day03

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type AocDay3 struct{}

const DIR = "day03/"

func (d AocDay3) Puzzle1(useSample int) {

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
		line                  string
		bite                  Bite
		x, left, joltage, sum int
	)

	for scanner.Scan() {

		line = scanner.Text()
		if useSample > 0 {
			fmt.Println("")
			fmt.Println(line)
		}

		h := &BiteHeap{}
		heap.Init(h)

		// Add all except the final digit to a max heap
		for x = 0; x < len(line)-1; x++ {

			heap.Push(h, Bite{
				i: x,
				b: line[x],
			})

		}

		// Pop the highest value digit
		bite = heap.Pop(h).(Bite)

		left = bite.i
		joltage = int(bite.b-'0') * 10

		if useSample > 0 {
			fmt.Println(bite)
		}

		// Push the final digit, that we left before
		heap.Push(h, Bite{
			i: x,
			b: line[x],
		})

		// Pop digits until we find the largest
		// that has an index AFTER the one we selected above
		for h.Len() > 0 {
			bite = heap.Pop(h).(Bite)
			if bite.i < left {
				continue
			}
			if useSample > 0 {
				fmt.Println(bite)
			}
			joltage += int(bite.b - '0')
			break
		}

		sum += joltage
		if useSample > 0 {
			fmt.Println("Added", joltage, "to", sum)
		}

	}

	fmt.Println("")
	fmt.Println("Total Joltage:", sum)

}

func (d AocDay3) Puzzle2(useSample int) {

}
