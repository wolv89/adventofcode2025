package day10

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

func (d AocDay10) Puzzle2(useSample int) {

}
