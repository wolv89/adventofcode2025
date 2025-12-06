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

type Problem struct {
	nums int
	op   byte
}

func (d AocDay6) Puzzle2(useSample int) {

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
		line                string
		val, calc, total    int64
		x, y, z, size, left int
		operator            byte
	)

	data := make([]string, 0)
	problems := make([]Problem, 0)

	for scanner.Scan() {

		line = scanner.Text()

		if line[0] == OP_ADD || line[0] == OP_MUL {

			for x = range line {
				if line[x] != ' ' {
					if size > 0 {
						problems = append(problems, Problem{
							nums: size,
							op:   operator,
						})
					}
					// Start at 0, don't count the operator as a space, just those in between
					size = 0
					operator = line[x]
				} else {
					size++
				}
			}

			if size > 0 {
				// No extra column here so add 1 or we're short
				problems = append(problems, Problem{
					nums: size + 1,
					op:   operator,
				})
			}

		} else {
			data = append(data, line)
		}

	}

	width := len(data) // Width, but top to bottom - 3 rows of data means each number is up to 3 digits wide
	left = 0

	for _, p := range problems {

		if useSample > 0 {
			fmt.Println("")
			fmt.Println("PROBLEM: ", string(p.op))
		}

		values := make([][]byte, p.nums)
		for y = range values {
			values[y] = make([]byte, width)
		}

		for x, z = left, 0; x < left+p.nums; x, z = x+1, z+1 {
			for y = 0; y < width; y++ {
				values[z][y] = data[y][x]
			}
		}

		calc = 0

		for _, v := range values {

			if useSample > 0 {
				fmt.Println(string(v))
			}

			val, err = strconv.ParseInt(strings.TrimSpace(string(v)), 10, 64)
			if err != nil {
				log.Fatal(err.Error())
			}

			if calc == 0 {
				calc = val
			} else {
				switch p.op {
				case OP_ADD:
					calc += val
				case OP_MUL:
					calc *= val
				}
			}

		}

		total += calc

		left += p.nums + 1

	}

	fmt.Println("")
	fmt.Println("Total:", total)

}
