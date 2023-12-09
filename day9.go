package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diffArray(arr []int) {
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i] - arr[i-1]
	}
}

func diffArrayPart2(arr []int) {
	for i := len(arr); i > 1; i-- {
		arr[i-1] = arr[i-1] - arr[i-2]
	}
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	total := 0
	for s.Scan() {
		line := s.Text()

		arr := []int{}
		for _, w := range strings.Split(line, " ") {
			x, err := strconv.Atoi(w)
			if err != nil {
				panic("how could it?!")
			}
			arr = append(arr, x)
		}

		allZero := func(a []int) bool {
			for _, x := range a {
				if x != 0 {
					return false
				}
			}
			return true
		}

		level := 0
		for !allZero(arr[:len(arr)-level]) && level < len(arr)-1 {
			diffArray(arr[:len(arr)-level])
			level += 1
		}

		for level > 0 {
			arr[len(arr)-level] += arr[len(arr)-1-level]
			level -= 1
		}
		total += arr[len(arr)-1]
	}

	return fmt.Sprintf("%d", total)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	total := 0
	for s.Scan() {
		line := s.Text()

		arr := []int{}
		for _, w := range strings.Split(line, " ") {
			x, err := strconv.Atoi(w)
			if err != nil {
				panic("how could it?!")
			}
			arr = append(arr, x)
		}

		allZero := func(a []int) bool {
			for _, x := range a {
				if x != 0 {
					return false
				}
			}
			return true
		}

		level := 0
		for !allZero(arr[level:]) && level < len(arr)-1 {
			diffArrayPart2(arr[level:])
			level += 1
		}

		for level > 0 {
			arr[level-1] -= arr[level]
			level -= 1
		}
		total += arr[0]
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
