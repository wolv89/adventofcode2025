package day10

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type AocDay10 struct{}

const DIR = "day10/"

type Machine struct {
	buttons       []uint16
	size, presses int
	lights        uint16
}

func (m *Machine) PressButtons(b, p int, state uint16) {

	if state == m.lights {
		m.presses = min(m.presses, p)
		return
	}

	for i := b; i < len(m.buttons); i++ {

		state ^= m.buttons[i]
		m.PressButtons(i+1, p+1, state)

		state ^= m.buttons[i]
		m.PressButtons(i+1, p, state)

	}

}

func (d AocDay10) Puzzle1(useSample int) {

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

	machines := make([]Machine, 0)

	for scanner.Scan() {

		line := scanner.Text()

		break1 := strings.Index(line, " ")
		break2 := strings.LastIndex(line, " ")

		l1 := line[0:break1]
		l2 := line[break1+1 : break2]
		// l3 := line[break2+1:] -- Not needed for part 1

		buttonList := strings.Split(l2, " ")

		btns := make([]uint16, 0, len(buttonList))
		for _, button := range buttonList {
			btns = append(btns, GetButton(button))
		}

		machines = append(machines, Machine{
			buttons: btns,
			size:    len(l1) - 2, // Don't count the square brackets
			lights:  GetLights(l1),
		})

	}

	var totalPresses int

	for m, mc := range machines {

		if useSample > 0 {
			fmt.Println("MACHINE", m+1)
			fmt.Println(Brender(mc.lights, mc.size))
			for _, b := range mc.buttons {
				fmt.Println("\t", Brender(b, mc.size))
			}
		}

		mc.presses = math.MaxInt
		mc.PressButtons(0, 0, 0)
		totalPresses += mc.presses

		if useSample > 0 {
			fmt.Println("Presses:", mc.presses)
			fmt.Println("")
		}

	}

	fmt.Println("")
	fmt.Println("Total Presses:", totalPresses)

}

func GetLights(line string) uint16 {

	var light uint16

	for i := range line {
		if line[i] == '#' {
			light += 1 << (i - 1)
		}
	}

	return light

}

func GetButton(btn string) uint16 {

	var button uint16

	for i := range btn {
		if btn[i] >= '0' && btn[i] <= '9' {
			button += 1 << int(btn[i]-'0')
		}
	}

	return button

}

// Binary Render
func Brender(x uint16, l int) string {

	b := fmt.Sprintf("%016b", x)

	var s strings.Builder
	s.Grow(l)

	for i := 15; i > 15-l; i-- {
		s.WriteByte(b[i])
	}

	return s.String()

}

/* +----------------------------+ PART 2 +----------------------------+ */

type Button [10]uint16

type Machine2 struct {
	buttons []Button
	target  Button
	size    int

	memo  map[Button]int
	costs map[Button]map[Button]int
}

/*
 * Big thanks to @tenthmascot on Reddit for their very clever python solution
 * Adapted to Golang below, with some help from Gemini
 *
 * Source: https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
 */
func (d AocDay10) Puzzle2(useSample int) {

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

	machines := make([]Machine2, 0)

	for scanner.Scan() {

		line := scanner.Text()

		break1 := strings.Index(line, " ")
		break2 := strings.LastIndex(line, " ")

		// l1 := line[0:break1]
		l2 := line[break1+1 : break2]
		l3 := line[break2+1:]

		buttonList := strings.Split(l2, " ")
		btns := make([]Button, 0, len(buttonList))

		for _, l := range buttonList {
			btns = append(btns, GetButton2(l))
		}

		tg, sz := GetTarget(l3)

		machines = append(machines, Machine2{
			buttons: btns,
			target:  tg,
			size:    sz,
			memo:    make(map[Button]int),
		})

	}

	total := 0

	for mi, mc := range machines {

		fmt.Println("CALCULATING MACHINE:", mi+1)

		mc.GeneratePatterns()

		solved := mc.Solve(mc.target)
		fmt.Println("Solved:", solved)

		total += solved
		fmt.Println("")

	}

	fmt.Println("")
	fmt.Println("Total:", total)

}

func GetButton2(line string) Button {

	var btn Button

	for _, v := range line {
		if v >= '0' && v <= '9' {
			btn[int(v-'0')] = 1
		}
	}

	return btn

}

func GetTarget(line string) (Button, int) {

	line = strings.Trim(line, "{}")
	vals := strings.Split(line, ",")

	var btn Button

	for v := range vals {
		n, _ := strconv.ParseUint(vals[v], 10, 64)
		btn[v] = uint16(n)
	}

	return btn, len(vals)

}

func (m *Machine2) GeneratePatterns() {

	m.costs = make(map[Button]map[Button]int)
	n := len(m.buttons)

	for i := 0; i < (1 << n); i++ {

		var pattern Button
		pressed := 0

		for j := range n {
			if (i>>j)&1 == 1 {
				pressed++
				for k := range m.size {
					pattern[k] += m.buttons[j][k]
				}
			}
		}

		var parity Button
		for k := range m.size {
			parity[k] = pattern[k] % 2
		}

		if _, ok := m.costs[parity]; !ok {
			m.costs[parity] = make(map[Button]int)
		}

		if val, ok := m.costs[parity][pattern]; !ok || pressed < val {
			m.costs[parity][pattern] = pressed
		}

	}

}

func (m *Machine2) Solve(goal Button) int {

	isZero := true
	var parity Button

	for i := range m.size {
		if goal[i] != 0 {
			isZero = false
		}
		parity[i] = goal[i] % 2
	}

	if isZero {
		return 0
	}

	if val, ok := m.memo[goal]; ok {
		return val
	}

	ans := 1_000_000

	if patterns, ok := m.costs[parity]; ok {
		for vec, cost := range patterns {
			canSubtract := true
			var nextGoal Button
			for i := range m.size {
				if vec[i] > goal[i] {
					canSubtract = false
					break
				}
				nextGoal[i] = (goal[i] - vec[i]) / 2
			}

			if canSubtract {
				ans = min(ans, cost+2*m.Solve(nextGoal))
			}
		}
	}

	m.memo[goal] = ans
	return ans

}
