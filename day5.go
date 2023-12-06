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

type Range struct {
	// both inclusive
	start, end uint
}

// Splits range `r` such that points overlapping with
// `or` are extracted as a new range and remaining non-contiguous
// pieces of `r` are also returned as new ranges along with the
// extracted range. If no overlap, empty slice is returned.
func (r Range) split(or Range) []Range {
	if or.end < r.start || r.end < or.start {
		return []Range{r}
	}

	if r.start < or.start && or.end < r.end {
		return []Range{
			Range{r.start, or.start - 1},
			or, Range{or.end + 1, r.end},
		}
	}

	if or.end < r.end {
		return []Range{
			Range{or.end + 1, r.end},
			Range{r.start, or.end},
		}
	}

	if or.start > r.start {
		return []Range{
			Range{r.start, or.start - 1},
			Range{or.start, r.end},
		}
	}

	if or.start <= r.start && r.end <= or.end {
		return []Range{r}
	}

	panic("should not reach?")
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
	seedRanges := []Range{}
	for seedSplit, i := strings.Split(s.Text()[7:], " "), 0; i < len(seedSplit); i += 2 {
		start, _ := strconv.Atoi(seedSplit[i])
		count, _ := strconv.Atoi(seedSplit[i+1])
		seedRanges = append(seedRanges, Range{uint(start), uint(start) + uint(count)})
	}
	s.Scan() // skip white line

	moreMaps := [TOTAL_MAPS][]RangeRemap{}

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
	}

	for level := 0; level < TOTAL_MAPS; level += 1 {
		for _, rr := range moreMaps[level] {
			for i := 0; i < len(seedRanges); i += 1 {
				cr := seedRanges[i]
				newRanges := cr.split(Range{start: rr.src, end: rr.src + rr.cnt - 1})
				if len(newRanges) == 1 {
					continue
				}

				// we will mark cr as invalid by setting
				// end = 0 and start = 1, so that later
				// these can be filtered out
				seedRanges[i] = Range{1, 0}

				for _, nr := range newRanges {
					seedRanges = append(seedRanges, nr)
				}
			}
		}

		for i, r := range seedRanges {
			if r.end == 0 && r.start == 1 {
				continue
			}

			for _, rr := range moreMaps[level] {
				if ok, dst := rr.includes(r.start); ok {
					seedRanges[i] = Range{dst, dst + (r.end - r.start)}
					break
				}
			}
		}
	}

	var minLoc uint = 1e9 // don't do this in real code
	for _, r := range seedRanges {
		if r.end == 0 && r.start == 1 {
			continue
		}

		if r.start < minLoc {
			minLoc = r.start
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
