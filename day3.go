package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const SIZE = 140

type Symbol byte
type Number int

func contains(s []*Number, e *Number) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getNeighNumbers(table [SIZE][SIZE]any, cx, cy int) (res []*Number) {
	for i := cx - 1; i <= cx+1; i += 1 {
		if i < 0 || i > SIZE-1 {
			continue
		}

		for j := cy - 1; j <= cy+1; j += 1 {
			if j < 0 || j > SIZE-1 {
				continue
			}

			if v, ok := table[j][i].(*Number); ok && !contains(res, v) {
				res = append(res, v)
			}
		}
	}

	return
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	var table [SIZE][SIZE]any // really either Symbol or *Number
	row := 0
	for s.Scan() {
		line := s.Bytes()

		for i := 0; i < len(line); i += 1 {
			if line[i] == '.' {
				table[row][i] = Symbol('.')
				continue
			}

			var num Number = -1
			di := i
			for ; di < SIZE && unicode.IsDigit(rune(line[di])); di += 1 {
				if num == -1 {
					num = 0
				}

				num = num*10 + Number(line[di]-'0')
			}

			if num != -1 {
				for ; i < di; i += 1 {
					table[row][i] = &num
				}
				i -= 1
				continue
			}

			table[row][i] = Symbol(line[i])
		}

		row += 1
	}

	var sum Number = 0
	for i := 0; i < SIZE; i += 1 {
		for j := 0; j < SIZE; j += 1 {
			if v, ok := table[j][i].(Symbol); ok && v != '.' {
				for _, n := range getNeighNumbers(table, i, j) {
					sum += *n
				}
			}
		}
	}

	return fmt.Sprintf("%v", sum)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	var table [SIZE][SIZE]any // really either Symbol or *Number
	row := 0
	for s.Scan() {
		line := s.Bytes()

		for i := 0; i < len(line); i += 1 {
			if line[i] == '.' {
				table[row][i] = Symbol('.')
				continue
			}

			var num Number = -1
			di := i
			for ; di < SIZE && unicode.IsDigit(rune(line[di])); di += 1 {
				if num == -1 {
					num = 0
				}

				num = num*10 + Number(line[di]-'0')
			}

			if num != -1 {
				for ; i < di; i += 1 {
					table[row][i] = &num
				}
				i -= 1
				continue
			}

			table[row][i] = Symbol(line[i])
		}

		row += 1
	}

	var sum Number = 0
	for i := 0; i < SIZE; i += 1 {
		for j := 0; j < SIZE; j += 1 {
			if v, ok := table[j][i].(Symbol); ok && v == '*' {
				nums := getNeighNumbers(table, i, j)
				if len(nums) != 2 {
					continue
				}
				sum += *nums[0] * *nums[1]
			}
		}
	}

	return fmt.Sprintf("%v", sum)
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
