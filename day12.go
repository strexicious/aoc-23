package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	row    string
	groups []int
}

func canPlace(expected string, have string) bool {
	if len(expected) != len(have) {
		panic("expected same lengths")
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != '?' && expected[i] != have[i] {
			return false
		}
	}

	return true
}

func arrgCount(cache map[string]int, generated []byte, validator string, taken int, gsize ...int) int {
	gsizeStr := strconv.Itoa(taken) + ":"
	for _, gs := range gsize {
		gsizeStr += strconv.Itoa(gs) + ","
	}

	// poor man's dynamic programming with poor tuples (a.k.a. gsizeStr)
	if res, ok := cache[gsizeStr]; ok {
		return res
	}

	if len(gsize) == 1 {
		count := 0
		for i := taken; len(validator)-i >= gsize[0]; i++ {
			generated = generated[:taken]
			for j := taken; j < i; j++ {
				generated = append(generated, '.')
			}
			for j := 0; j < gsize[0]; j++ {
				generated = append(generated, '#')
			}
			for j := i + gsize[0]; j < len(validator); j++ {
				generated = append(generated, '.')
			}
			if canPlace(validator, string(generated)) {
				count++
			}
		}
		cache[gsizeStr] = count
		return count
	}

	needed := len(gsize) - 1
	for _, gs := range gsize {
		needed += gs
	}

	if len(validator)-taken < needed {
		return 0
	}

	count := 0
	for i := taken; len(validator)-i >= needed; i++ {
		generated = generated[:taken]
		for j := taken; j < i; j++ {
			generated = append(generated, '.')
		}
		for j := 0; j < gsize[0]; j++ {
			generated = append(generated, '#')
		}
		generated = append(generated, '.')

		if !canPlace(validator[:i+gsize[0]+1], string(generated)) {
			continue
		}

		count += arrgCount(cache, generated, validator, i+gsize[0]+1, gsize[1:]...)
	}

	cache[gsizeStr] = count

	return count
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	total := 0
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		rcrd := Record{row: line[0]}

		for _, gcs := range strings.Split(line[1], ",") {
			gc, _ := strconv.Atoi(gcs)
			rcrd.groups = append(rcrd.groups, gc)
		}

		total += arrgCount(map[string]int{}, make([]byte, 0, len(rcrd.row)), rcrd.row, 0, rcrd.groups...)
	}

	return fmt.Sprintf("%d", total)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	total := 0
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		rcrd := Record{row: line[0]}

		for _, gcs := range strings.Split(line[1], ",") {
			gc, _ := strconv.Atoi(gcs)
			rcrd.groups = append(rcrd.groups, gc)
		}

		// man it's hard to repeat and join slices and strings in go
		// the following works for part 2 so w/e
		rcrd.row = rcrd.row + "?" + rcrd.row + "?" + rcrd.row + "?" + rcrd.row + "?" + rcrd.row
		tmp := make([]int, 0, len(rcrd.groups)*5)
		for i := 0; i < 5; i++ {
			for _, e := range rcrd.groups {
				tmp = append(tmp, e)
			}
		}
		rcrd.groups = tmp

		total += arrgCount(map[string]int{}, make([]byte, 0, len(rcrd.row)), rcrd.row, 0, rcrd.groups...)
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
