package day05

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type AocDay5 struct{}

const DIR = "day05/"

const DASH = "-"

type Range struct {
	start, end int64
}

func (d AocDay5) Puzzle1(useSample int) {

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
		parts              []string
		line               string
		rn                 Range
		ing                int64
		n, freshCount      int
		readingIngredients bool
	)

	ranges := make([]Range, 0)
	ingredients := make([]int64, 0)

	for scanner.Scan() {

		line = scanner.Text()

		if readingIngredients {

			ing, err = strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			ingredients = append(ingredients, ing)

		} else {

			if line == "" {
				readingIngredients = true
				continue
			}

			parts = strings.Split(line, DASH)
			rn = Range{}

			rn.start, err = strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			rn.end, err = strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			ranges = append(ranges, rn)

		}

	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.end, b.end)
	})

	// Territory, as a synonym for range
	territory := make([]int64, 2, len(ranges)/2)

	territory[0] = ranges[0].start
	territory[1] = ranges[0].end

	n = 1

	for i := 1; i < len(ranges); i++ {

		rn = ranges[i]

		if rn.start <= territory[n] {
			territory[n-1] = min(territory[n-1], rn.start)
			territory[n] = max(territory[n], rn.end)
		} else {
			territory = append(territory, rn.start, rn.end)
			n += 2
		}

	}

	for _, ing = range ingredients {
		n, _ = slices.BinarySearch(territory, ing)
		if n%2 == 1 {
			freshCount++
		}
		if useSample > 0 {
			fmt.Println(ing, " > ", n)
		}
	}

	fmt.Println("")
	fmt.Println("Fresh ingredients:", freshCount)

}

func (d AocDay5) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	switch useSample {
	case 1:
		datafile = DIR + "sample.txt"
	case 2:
		datafile = DIR + "sample2.txt"
	case 3:
		datafile = DIR + "sample3.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		parts                  []string
		line                   string
		start, end, sum, point int64
		sweep, i               int
	)

	starts := make(map[int64]int)
	ends := make(map[int64]int)
	points := make(map[int64]struct{})

	for scanner.Scan() {

		line = scanner.Text()

		if line == "" {
			break
		}

		parts = strings.Split(line, DASH)

		start, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		end, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		starts[start]++
		ends[end]++

		points[start] = struct{}{}
		points[end] = struct{}{}

	}

	sortedPoints := make([]int64, 0, len(points))
	for point = range points {
		sortedPoints = append(sortedPoints, point)
	}

	slices.Sort(sortedPoints)
	// start = sortedPoints[0]

	for i, point = range sortedPoints {

		if useSample > 0 {
			fmt.Println(i, ": ", point, " [", ends[point], ",", starts[point], "] ", sweep, " - ", sum)
		}

		if sweep > 0 {
			sum += point - sortedPoints[i-1]
		}

		sweep -= ends[point]
		sweep += starts[point]

		if sweep == 0 {
			sum++
		}

	}

	fmt.Println("")
	fmt.Println("Fresh ingredients:", sum)

}

// Almost, but no...
// In fact it was wrong, then I refactored and got it 1 lower, still wrong
// Aaaand the actual answer was just 1 more below... thanks Eric (!)
func (d AocDay5) PuzzlePoo(useSample int) {

	datafile := DIR + "data.txt"
	switch useSample {
	case 1:
		datafile = DIR + "sample.txt"
	case 2:
		datafile = DIR + "sample2.txt"
	case 3:
		datafile = DIR + "sample3.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		parts           []string
		line            string
		rn              Range
		freshCount, sum int64
		n, i, j         int
	)

	ranges := make([]Range, 0)

	for scanner.Scan() {

		line = scanner.Text()

		if line == "" {
			break
		}

		parts = strings.Split(line, DASH)
		rn = Range{}

		rn.start, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		rn.end, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		ranges = append(ranges, rn)

	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.end == b.end {
			return cmp.Compare(a.start, b.start)
		}
		return cmp.Compare(a.end, b.end)
	})

	// Territory, as a synonym for range
	territory := make([]int64, 2, len(ranges)/2)

	territory[0] = ranges[0].start
	territory[1] = ranges[0].end

	n = 1
	_ = j

	for i = 1; i < len(ranges); i++ {

		rn = ranges[i]
		if useSample > 0 {
			fmt.Println(rn)
			fmt.Println(territory)
			fmt.Println("")
		}

		if rn.start <= territory[n] {
			territory[n-1] = min(territory[n-1], rn.start)
			territory[n] = rn.end
		} else {
			territory = append(territory, rn.start, rn.end)
			n += 2
		}

		for j = n - 3; j >= 0 && territory[j] >= rn.start; j -= 2 {
			territory[j] = rn.start
			territory[j+1] = rn.end
			territory = territory[:j+2]
		}

	}

	for i = 0; i < len(territory); i += 2 {
		sum = territory[i+1] - territory[i] + 1
		freshCount += sum
		fmt.Println(territory[i], "-", territory[i+1], "  +", sum)
	}

	fmt.Println("")
	fmt.Println("Fresh ingredients:", freshCount)

}
