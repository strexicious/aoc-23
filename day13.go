package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Set map[int]bool

func findVertAxes(row []byte) Set {
	res := Set{}
	for i := 0; i < len(row)-1; i++ {
		r := i + 1
		if r > len(row)-r {
			r = len(row) - r
		}

		isRefl := true
		for j := 0; j < r; j++ {
			if row[i-j] != row[i+j+1] {
				isRefl = false
				break
			}
		}

		if isRefl {
			res[i] = true
		}
	}

	return res
}

func transpose(grid [][]byte) [][]byte {
	rowLen := len(grid)
	resGrid := make([][]byte, 0, len(grid[0]))

	for i := 0; i < len(grid[0]); i++ {
		resGrid = append(resGrid, make([]byte, rowLen))
		for j := 0; j < rowLen; j++ {
			resGrid[i][j] = grid[j][i]
		}
	}

	return resGrid
}

func intersect(s Set, sets ...Set) Set {
	res := s
	for _, ss := range sets {
		tmp := Set{}
		for e, _ := range ss {
			if _, ok := res[e]; ok {
				tmp[e] = true
			}
		}
		res = tmp
	}
	return res
}

func part1(in string) string {
	groups := strings.Split(in, "\n\n")
	total := 0
	for _, g := range groups {
		rdr := strings.NewReader(g)
		s := bufio.NewScanner(rdr)

		grid := [][]byte{}
		for s.Scan() {
			grid = append(grid, bytes.Clone(s.Bytes()))
		}

		allVAxes := []Set{}
		for _, l := range grid {
			allVAxes = append(allVAxes, findVertAxes(l))
		}

		allHAxes := []Set{}
		grid = transpose(grid)
		for _, l := range grid {
			allHAxes = append(allHAxes, findVertAxes(l))
		}

		// WARNING: see remark about part 1 in part 2 which
		// makes use of singleOutLine function. Essentially you don't
		// really need to intersect the sets per say, just count for all elements
		// in all sets how many times they appear, and then the one that
		// has count len(sets) is our answer
		// now it's uniqueness can be assumed because otherwise it implies
		// multiple reflection axis which the problem doesn't really specify
		intV := intersect(allVAxes[0], allVAxes[1:]...)
		intH := intersect(allHAxes[0], allHAxes[1:]...)

		// WARNING: BAD CODE AHEAD
		// i mean intV and intH are expected to only have 1 element but w/e
		for k, _ := range intV {
			total += k + 1
		}
		for k, _ := range intH {
			total += (k + 1) * 100
		}
	}

	return fmt.Sprintf("%d", total)
}

func singleOutLine(sets ...Set) int {
	counts := map[int]int{}
	for _, ss := range sets {
		for e, _ := range ss {
			if _, ok := counts[e]; !ok {
				counts[e] = 0
			}

			counts[e]++
		}
	}

	for k, c := range counts {
		if c == len(sets)-1 {
			return k
		}
	}

	return -1
}

func part2(in string) string {
	groups := strings.Split(in, "\n\n")
	total := 0
	for _, g := range groups {
		rdr := strings.NewReader(g)
		s := bufio.NewScanner(rdr)

		grid := [][]byte{}
		for s.Scan() {
			grid = append(grid, bytes.Clone(s.Bytes()))
		}

		allVAxes := []Set{}
		for _, l := range grid {
			allVAxes = append(allVAxes, findVertAxes(l))
			// fmt.Println(allVAxes[len(allVAxes)-1])
		}

		allHAxes := []Set{}
		grid = transpose(grid)
		for _, l := range grid {
			allHAxes = append(allHAxes, findVertAxes(l))
			// fmt.Println(allHAxes[len(allHAxes)-1])
		}

		// really part 1 could make use of a logic like this
		// you would just need to single out a line that
		// appears len(sets) times rather than len(sets)-1
		if v := singleOutLine(allVAxes...); v != -1 {
			total += v + 1
		} else {
			total += (singleOutLine(allHAxes...) + 1) * 100
		}
	}

	return fmt.Sprintf("%d", total)
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
