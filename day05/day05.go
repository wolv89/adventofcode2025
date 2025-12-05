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

}
