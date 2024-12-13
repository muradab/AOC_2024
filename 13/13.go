package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Button struct {
	x, y int
}

type Destination struct {
	x, y int
}

type Machine struct {
	a, b Button
	d    Destination
}

func main() {
	file, err := os.ReadFile("13/day13.txt")

	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(file), "\n")

	machines := make([]Machine, 0)

	for i := 0; i < len(lines)/4; i++ {
		var a, b Button
		var d Destination
		for j := 0; j < 4; j++ {

			if j%4 == 0 {
				a = parseButton(lines[(i*4)+j])
			}
			if j%4 == 1 {
				b = parseButton(lines[(i*4)+j])
			}
			if j%4 == 2 {
				d = parsePrize(lines[(i*4)+j])

			}
		}

		machines = append(machines, Machine{a, b, d})

	}

	total := 0
	part2 := 0

	for m := range machines {
		temp := canReach(machines[m])
		if temp != math.MaxInt {
			total += temp
		}
		another := part2Reach(Machine{machines[m].a, machines[m].b, Destination{machines[m].d.x + 10000000000000, machines[m].d.y + 10000000000000}})
		part2 += another
	}

	fmt.Println(total)
	fmt.Println(part2)

}

func parseButton(s string) Button {
	s = strings.ReplaceAll(s, ",", "")
	w := strings.Split(s, " ")
	// fmt.Println(w)
	x := strings.Split(w[2], "+")
	y := strings.Split(w[3], "+")
	// fmt.Println(x, y)
	// s = strings.ReplaceAll(s, "Y=", "")
	// ss := strings.Split(s, ",")
	j, _ := strconv.Atoi(x[1])
	k, _ := strconv.Atoi(y[1])
	return Button{j, k}
}
func parsePrize(s string) Destination {
	s = strings.ReplaceAll(s, ",", "")
	w := strings.Split(s, " ")
	// fmt.Println(w)
	x := strings.Split(w[1], "=")
	y := strings.Split(w[2], "=")
	// fmt.Println(x, y)
	// s = trings.ReplaceAll(s, "Y=", "")
	// ss := strings.Split(s, ",")
	j, _ := strconv.Atoi(x[1])
	// fmt.Println(x[1], y[1])
	k, _ := strconv.Atoi(y[1])
	return Destination{j, k}
}

func canReach(m Machine) int {
	b1 := m.a
	b2 := m.b
	d := m.d
	ans := math.MaxInt
	for i := 0; i <= 100; i++ {
		left_x := d.x - i*b1.x
		left_y := d.y - i*b1.y

		if left_x < 0 || left_y < 0 {
			return ans
		}
		if left_x%b2.x != 0 && left_y%b2.y != 0 {
			continue
		}

		count_b2x := left_x / b2.x
		count_b2y := left_y / b2.y
		if count_b2x != count_b2y {
			continue
		}
		if count_b2x > 100 {
			continue
		}
		ans = min(ans, (i*3)+count_b2x)

	}
	return ans
}

func part2Reach(m Machine) int {

	// calculate a and b

	b := (m.d.y*m.a.x - m.d.x*m.a.y) / (m.b.y*m.a.x - m.b.x*m.a.y)
	a := (m.d.x - b*m.b.x) / m.a.x

	if a*m.a.x+m.b.x*b != m.d.x || a*m.a.y+m.b.y*b != m.d.y {
		return 0
	}

	return (3 * a) + b

}
