package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type AocDay1 struct{}

const DIR = "day01/"

func (d AocDay1) Puzzle1(useSample int) {

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
		line                []byte
		dial                int = 50
		move, dir, password int
	)

	for scanner.Scan() {

		line = scanner.Bytes()

		if line[0] == 'R' {
			dir = 1
		} else {
			dir = -1
		}

		move, err = strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatal(err.Error())
		}

		dial += move * dir
		dial %= 100
		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			password++
		}

		if useSample > 0 {
			fmt.Println(string(line), dial)
		}

	}

	fmt.Println("")
	fmt.Println("Password: ", password)

}

func (d AocDay1) Puzzle2(useSample int) {

}
