package day02

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type AocDay2 struct{}

const DIR = "day02/"

const (
	COMMA = ','
	DASH  = '-'
)

func (d AocDay2) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	ranges := bytes.Split(data, []byte{COMMA})

	var (
		upperStr, lowerStr, upperHalfStr, lowerHalfStr           string
		upper, lower, i, upperHalf, lowerHalf, check, invalidSum int64
		sep, invalidCount                                        int
	)

	for _, rng := range ranges {

		if useSample > 0 {
			fmt.Println("")
			fmt.Println("RANGE:", string(rng))
			fmt.Println("---------------------")
		}

		sep = bytes.IndexByte(rng, DASH)

		lowerStr = string(rng[0:sep])
		upperStr = string(rng[sep+1:])

		if len(lowerStr)%2 != 0 && len(upperStr)%2 != 0 {
			if useSample > 0 {
				fmt.Println(" - Odd length ranges, skipping...")
			}
			continue
		}

		// Skipping error handling for trusted input
		// (Don't do this at home - except I am...)
		lower, _ = strconv.ParseInt(lowerStr, 10, 64)
		upper, _ = strconv.ParseInt(upperStr, 10, 64)

		lowerHalfStr = string(rng[0 : sep/2])
		upperHalfStr = string(rng[sep+1 : sep+1+(len(rng)-sep)/2])

		lowerHalf, _ = strconv.ParseInt(lowerHalfStr, 10, 64)
		upperHalf, _ = strconv.ParseInt(upperHalfStr, 10, 64)

		for i = lowerHalf; i <= upperHalf; i++ {

			for check = 1; check <= i; check *= 10 {
			}
			check = check*i + i

			if useSample > 0 {
				fmt.Println("Trying:", check)
			}

			if check < lower {
				if useSample > 0 {
					fmt.Println(" - Too low")
				}
				continue
			}

			if check > upper {
				if useSample > 0 {
					fmt.Println(" - Too high")
				}
				break
			}

			if useSample > 0 {
				fmt.Println(" - Bingo")
			}
			invalidCount++
			invalidSum += check

		}

	}

	fmt.Println("")
	fmt.Println("Invalid IDs:", invalidSum, " (", invalidCount, ")")

}

func (d AocDay2) Puzzle2(useSample int) {

}
