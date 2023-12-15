package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Lens struct {
	label       string
	focalLength byte
}

type HashMap struct {
	boxes [256][]Lens
}

func hash(input string) byte {
	var res byte
	for _, b := range []byte(input) {
		res += b
		res *= 17
	}
	return res
}

func part1(in string) string {
	total := 0
	for _, ss := range strings.Split(in, ",") {
		total += int(hash(ss))
	}

	return fmt.Sprintf("%d", total)
}

func part2(in string) string {
	hmap := HashMap{}
	stepRE := regexp.MustCompile("([a-z]+)(=[0-9]|-)")
	for _, ss := range strings.Split(in, ",") {
		splits := stepRE.FindStringSubmatch(ss)[1:]
		boxId := hash(splits[0])
		didFind := false
		for i, l := range hmap.boxes[boxId] {
			if splits[1][0] == '-' && l.label == splits[0] {
				hmap.boxes[boxId][i].label = ""
				didFind = true // this one may not be needed
				break
			}

			if splits[1][0] == '=' && l.label == splits[0] {
				hmap.boxes[boxId][i].focalLength = splits[1][1] - '0'
				didFind = true
				break
			}
		}
		if splits[1][0] == '=' && !didFind {
			hmap.boxes[boxId] = append(hmap.boxes[boxId], Lens{splits[0], splits[1][1] - '0'})
		}
	}

	total := 0
	for i, b := range hmap.boxes {
		slotNum := 0
		for _, l := range b {
			if l.label == "" {
				continue
			}
			slotNum++
			total += (i + 1) * slotNum * int(l.focalLength)
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
