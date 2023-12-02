package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sample struct {
	red   int
	green int
	blue  int
}

func ParseSample(in string) Sample {
	res := Sample{}
	var err error
	for _, color := range strings.Split(in, ", ") {
		if bcountS, ok := strings.CutSuffix(color, " red"); ok {
			res.red, err = strconv.Atoi(bcountS)
		}
		if bcountS, ok := strings.CutSuffix(color, " green"); ok {
			res.green, err = strconv.Atoi(bcountS)
		}
		if bcountS, ok := strings.CutSuffix(color, " blue"); ok {
			res.blue, err = strconv.Atoi(bcountS)
		}
	}

	if err != nil {
		panic("something happened parsing Sample")
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	posGameIds := []int{}
	for s.Scan() {
		line := s.Text()

		splits := strings.Split(line, ": ")
		gameId, err := strconv.Atoi(strings.TrimPrefix(splits[0], "Game "))
		if err != nil {
			panic("crashed parsing game id")
		}

		gamePossible := true
		splits = strings.Split(splits[1], "; ")
		for _, s := range splits {
			sample := ParseSample(s)

			if sample.red > 12 || sample.green > 13 || sample.blue > 14 {
				gamePossible = false
				break
			}
		}

		if gamePossible {
			posGameIds = append(posGameIds, gameId)
		}

	}

	var sum int
	for _, gameId := range posGameIds {
		sum += gameId
	}

	return fmt.Sprintf("%d", sum)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	powers := []int{}
	for s.Scan() {
		line := s.Text()

		splits := strings.Split(line, ": ")
		splits = strings.Split(splits[1], "; ")

		var minRed, minGreen, minBlue int
		for _, s := range splits {
			sample := ParseSample(s)

			if sample.red != 0 {
				minRed = max(minRed, sample.red)
			}
			if sample.green != 0 {
				minGreen = max(minGreen, sample.green)
			}
			if sample.blue != 0 {
				minBlue = max(minBlue, sample.blue)
			}
		}

		powers = append(powers, minRed*minGreen*minBlue)
	}

	var sum int
	for _, pow := range powers {
		sum += pow
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
