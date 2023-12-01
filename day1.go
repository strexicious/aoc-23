package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func reverseString(in string) string {
	res := make([]rune, len([]rune(in)))

	for i, r := range in {
		res[len(res)-i-1] = r
	}

	return string(res)
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	vals := make([]int, 0, 1000)
	for s.Scan() {
		line := s.Text()

		var a int
		for _, c := range line {
			if unicode.IsNumber(c) {
				a = int(c - '0')
				break
			}
		}

		var b int
		for _, c := range line {
			if unicode.IsNumber(c) {
				b = int(c - '0')
			}
		}

		vals = append(vals, a*10+b)
		fmt.Println(a, b)
	}

	var total int
	for _, x := range vals {
		total += x
	}
	return fmt.Sprintf("%d", total)
}

func part2(in string) string {
	w2d := func(in string) int {
		switch in {
		case "one":
			return 1
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		default:
			if len([]rune(in)) != 1 || !unicode.IsNumber([]rune(in)[0]) {
				panic("wooo no mapping: " + in)
			}
			return int([]rune(in)[0] - '0')
		}
	}

	rf := regexp.MustCompile("0|1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine")
	rb := regexp.MustCompile(reverseString("0|1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine"))

	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	vals := make([]int, 0, 1000)
	for s.Scan() {
		line := s.Text()

		linef := rf.FindString(line)
		lineb := reverseString(rb.FindString(reverseString(line)))

		a := w2d(linef)
		b := w2d(lineb)

		vals = append(vals, a*10+b)
	}

	var total int
	for _, x := range vals {
		total += x
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
