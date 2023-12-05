package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeRemap struct {
	src uint
	dst uint
	cnt uint
}

func (rr RangeRemap) includes(num uint) (bool, uint) {
	if rr.src <= num && num < rr.src+rr.cnt {
		return true, rr.dst + (num - rr.src)
	}

	return false, 0
}

const TOTAL_MAPS = 7

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)
	s.Scan()
	seeds := []uint{}
	for _, seedid := range strings.Split(s.Text()[7:], " ") {
		seedid, _ := strconv.Atoi(seedid)
		seeds = append(seeds, uint(seedid))
	}
	s.Scan() // skip white line

	moreMaps := [TOTAL_MAPS][]RangeRemap{}
	invMoreMaps := [TOTAL_MAPS][]RangeRemap{}

	s.Scan() // skip header
	for mapIdx := 0; s.Scan(); {
		line := s.Text()
		if line == "" {
			s.Scan() // skip next header
			mapIdx += 1
			continue
		}

		lsplit := strings.Split(line, " ")
		dst, _ := strconv.Atoi(lsplit[0])
		src, _ := strconv.Atoi(lsplit[1])
		cnt, _ := strconv.Atoi(lsplit[2])

		moreMaps[mapIdx] = append(moreMaps[mapIdx], RangeRemap{uint(src), uint(dst), uint(cnt)})
		invMoreMaps[mapIdx] = append(invMoreMaps[mapIdx], RangeRemap{dst: uint(src), src: uint(dst), cnt: uint(cnt)})
	}

	var minLoc uint = 1e9 // don't do this in real code
	for _, seedid := range seeds {
		k := seedid

		for _, remaps := range moreMaps {
			for _, rr := range remaps {
				if ok, mapK := rr.includes(k); ok {
					k = mapK
					break
				}
			}
		}

		if k < minLoc {
			minLoc = k
		}
	}

	return fmt.Sprintf("%d", minLoc)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)
	s.Scan()
	seeds := []uint{}
	for _, seedid := range strings.Split(s.Text()[7:], " ") {
		seedid, _ := strconv.Atoi(seedid)
		seeds = append(seeds, uint(seedid))
	}
	s.Scan() // skip white line

	moreMaps := [TOTAL_MAPS][]RangeRemap{}
	invMoreMaps := [TOTAL_MAPS][]RangeRemap{}

	s.Scan() // skip header
	for mapIdx := 0; s.Scan(); {
		line := s.Text()
		if line == "" {
			s.Scan() // skip next header
			mapIdx += 1
			continue
		}

		lsplit := strings.Split(line, " ")
		dst, _ := strconv.Atoi(lsplit[0])
		src, _ := strconv.Atoi(lsplit[1])
		cnt, _ := strconv.Atoi(lsplit[2])

		moreMaps[mapIdx] = append(moreMaps[mapIdx], RangeRemap{uint(src), uint(dst), uint(cnt)})
		invMoreMaps[mapIdx] = append(invMoreMaps[mapIdx], RangeRemap{dst: uint(src), src: uint(dst), cnt: uint(cnt)})
	}

	var minLoc uint = 1e9 // don't do this in real code
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println("Seed group:", i>>1)
		for seedid := seeds[i]; seedid < seeds[i]+seeds[i+1]; seedid += 1 {
			k := seedid

			for _, remaps := range moreMaps {
				for _, rr := range remaps {
					if ok, mapK := rr.includes(k); ok {
						k = mapK
						break
					}
				}
			}

			if k < minLoc {
				minLoc = k
			}
		}
	}

	return fmt.Sprintf("%d", minLoc)
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
