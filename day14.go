package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func part1(in string) string {
	rows := strings.Split(in, "\n")
	cols := make([]int, len(rows[0]))

	afterRoll := make([][]byte, 0, len(rows))
	total := 0
	for j, r := range rows {
		afterRoll = append(afterRoll, bytes.Repeat([]byte("."), len(r)))
		for i, c := range r {
			if c == '#' {
				cols[i] = j + 1
				afterRoll[j][i] = '#'
			}
			if c == 'O' {
				afterRoll[cols[i]][i] = 'O'
				total += len(rows) - cols[i] - 1
				cols[i]++
			}
		}
	}

	return fmt.Sprintf("%d", total)
}

type Rock struct {
	x, y  int
	rolls bool
}

func part2(in string) string {
	rocks := []Rock{}
	var w, h int
	for y, r := range strings.Split(in, "\n") {
		for x, b := range r {
			if b == '#' {
				rocks = append(rocks, Rock{x, y, false})
			}
			if b == 'O' {
				rocks = append(rocks, Rock{x, y, true})
			}
			h = x
		}
		w = y
	}

	w++
	h++

	fmt.Print("[")
	for k := 0; k < 1000000000; k++ {
		// north
		{
			sort.Slice(rocks, func(i, j int) bool {
				return rocks[i].y < rocks[j].y
			})
			cols := make([]int, w)
			for i, r := range rocks {
				if r.rolls {
					rocks[i].y = cols[r.x]
					cols[r.x]++
				} else {
					cols[r.x] = r.y + 1
				}
			}
		}

		// // west
		{
			sort.Slice(rocks, func(i, j int) bool {
				return rocks[i].x < rocks[j].x
			})
			rows := make([]int, h)
			for i, r := range rocks {
				if r.rolls {
					rocks[i].x = rows[r.y]
					rows[r.y]++
				} else {
					rows[r.y] = r.x + 1
				}
			}
		}

		// south
		{
			sort.Slice(rocks, func(i, j int) bool {
				return rocks[i].y > rocks[j].y
			})
			cols := make([]int, w)
			for i := range cols {
				cols[i] = h - 1
			}
			for i, r := range rocks {
				if r.rolls {
					rocks[i].y = cols[r.x]
					cols[r.x]--
				} else {
					cols[r.x] = r.y - 1
				}
			}
		}

		// east
		{
			sort.Slice(rocks, func(i, j int) bool {
				return rocks[i].x > rocks[j].x
			})
			rows := make([]int, h)
			for i := range rows {
				rows[i] = w - 1
			}
			for i, r := range rocks {
				if r.rolls {
					rocks[i].x = rows[r.y]
					rows[r.y]--
				} else {
					rows[r.y] = r.x - 1
				}
			}
		}

		load := 0
		crr := make([]byte, w)
		for _, r := range rocks {
			if r.rolls {
				load += h - r.y
				crr[r.x]++
			}
		}

		fmt.Print(load, ",")

		if k == 1000 {
			break
		}
	}
	fmt.Println("]")

	return `Now analyze the previous list. Possibly plotting it with python.
You will notice that it is periodic. You can calculate the period
length, and from there calculate the load for after 1.000.000.000
cycles. Assume 1-index counting of list. Choose any i-th element
for which you are sure it falls withing some repeating interval
and not at the beginning where it does not repeat. Now the load
after the 1.000.000.000 cycles will be same as the load for after
i + j cycles where j = ((1e9 - i) % period.`
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
