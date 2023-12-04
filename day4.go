package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	sum := 0
	for s.Scan() {
		line := s.Text() + " EOL"

		lrdr := strings.NewReader(line)
		sl := bufio.NewScanner(lrdr)
		sl.Split(bufio.ScanWords)

		sl.Scan()
		sl.Scan()

		winning := map[int]bool{}
		for sl.Scan() {
			if sl.Text() == "|" {
				break
			}

			digs := sl.Text()
			num, err := strconv.Atoi(digs)
			if err != nil {
				panic("it can't be!")
			}
			winning[num] = true
		}

		totalMatches := -1
		for sl.Scan() {
			if sl.Text() == "EOL" {
				break
			}

			digs := sl.Text()
			num, err := strconv.Atoi(digs)
			if err != nil {
				panic("it can't be again!")
			}

			if _, ok := winning[num]; ok {
				totalMatches += 1
			}
		}

		if totalMatches != -1 {
			sum += int(math.Pow(2, float64(totalMatches)))
		}
	}

	return fmt.Sprintf("%d", sum)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	multFactors := map[int]int{}
	for currCard := 1; s.Scan(); currCard += 1 {
		line := s.Text() + " EOL"

		if _, ok := multFactors[currCard]; !ok {
			multFactors[currCard] = 1
		}

		lrdr := strings.NewReader(line)
		sl := bufio.NewScanner(lrdr)
		sl.Split(bufio.ScanWords)

		sl.Scan()
		sl.Scan()

		winning := map[int]bool{}
		for sl.Scan() {
			if sl.Text() == "|" {
				break
			}

			digs := sl.Text()
			num, err := strconv.Atoi(digs)
			if err != nil {
				panic("it can't be!")
			}
			winning[num] = true
		}

		totalMatches := 0
		for sl.Scan() {
			if sl.Text() == "EOL" {
				break
			}

			digs := sl.Text()
			num, err := strconv.Atoi(digs)
			if err != nil {
				panic("it can't be again!")
			}

			if _, ok := winning[num]; ok {
				totalMatches += 1
			}
		}

		ccmf := multFactors[currCard]
		for i := currCard + 1; i < currCard+1+totalMatches; i += 1 {
			if _, ok := multFactors[i]; !ok {
				multFactors[i] = 1
			}

			multFactors[i] += ccmf
		}
	}

	sum := 0
	for _, v := range multFactors {
		sum += v
	}

	return fmt.Sprintf("%d", sum)
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
