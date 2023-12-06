package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// for part 1
const TOTAL_RACES = 4

func part1(in string) string {
	splits := strings.Split(in, "\n")
	timeSplits := strings.Fields(splits[0])
	distSplits := strings.Fields(splits[1])

	totalCount := 1
	for i := 0; i < TOTAL_RACES; i += 1 {
		time, _ := strconv.Atoi(strings.TrimSpace(timeSplits[i+1]))
		dist, _ := strconv.Atoi(strings.TrimSpace(distSplits[i+1]))

		count := 0
		for t := 1; t < time; t += 1 {
			if t*(time-t) > dist {
				count += 1
			}
		}

		totalCount *= count
	}

	return fmt.Sprintf("%d", totalCount)
}

func part2(in string) string {
	// distance traveled = t * (time - t)
	// time :: constant
	// t :: variable
	// expand expression -t**2 + t*time
	// we want to count for how many discrete t does
	// the following hold: -t**2 + t*time > dist
	// dist :: constant
	// graphing out -t**2 + t*time - dist we see
	// that we want to find t1 and t2 (roots) and count how many
	// natural t are in range. I.e. size of { floor(t) : t1 < t < t2 }

	in = strings.ReplaceAll(in, " ", "")
	splits := strings.Split(in, "\n")
	time, _ := strconv.Atoi(splits[0][5:])
	dist, _ := strconv.Atoi(splits[1][9:])

	rad := math.Sqrt(float64(time*time - 4*dist))
	t1 := (-float64(time) + rad) / -2
	t2 := (-float64(time) - rad) / -2

	total := uint(t2) - uint(t1)

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
