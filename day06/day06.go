package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type AocDay6 struct{}

const DIR = "day06/"

const (
	OP_ADD = '+'
	OP_MUL = '*'

	OP_ADD_STR = "+"
	OP_ADD_MUL = "*"
)

func (d AocDay6) Puzzle1(useSample int) {

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
		ops        []byte
		line, item string
		val, total int64
		r, n, x    int
	)

	data := make([][]int64, 0)

	for scanner.Scan() {

		line = scanner.Text()

		fields := strings.Fields(line)
		n = len(fields) // Assuming this will be equal on every row - famous last words...?

		if fields[0] == OP_ADD_STR || fields[0] == OP_ADD_MUL {
			ops = make([]byte, 0, n)
			for _, item = range fields {
				ops = append(ops, item[0])
			}
			continue
		}

		data = append(data, make([]int64, 0, n))

		for _, item = range fields {

			val, err = strconv.ParseInt(item, 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			data[r] = append(data[r], val)

		}

		r++

	}

	if useSample > 0 {
		for _, row := range data {
			fmt.Println(row)
		}
		fmt.Print("[")
		// OOP ;)
		for o, op := range ops {
			if o > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%s", string(op))
		}
		fmt.Print("]\n\n")
	}

	calc := make([]int64, n)

	// Load first row
	for x = 0; x < n; x++ {
		calc[x] = data[0][x]
	}

	// Operate on remaining rows
	for r = 1; r < len(data); r++ {
		for x = 0; x < n; x++ {
			switch ops[x] {
			case OP_ADD:
				calc[x] += data[r][x]
			case OP_MUL:
				calc[x] *= data[r][x]
			}
		}
	}

	// Sum the total
	for _, val = range calc {
		total += val
	}

	fmt.Println("")
	fmt.Println("Total:", total)

}

func (d AocDay6) Puzzle2(useSample int) {

}
