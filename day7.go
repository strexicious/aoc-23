package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const VALID_CARDS = "AKQJT98765432"

type Hand [5]byte
type Bid struct {
	h Hand
	b int
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

func cardStregnth(c byte) int {
	return -strings.IndexByte(VALID_CARDS, c)
}

func cardStregnthPart2(c byte) int {
	if c == 'J' {
		return -20
	}
	return cardStregnth(c)
}

func newHand(cs []byte) Hand {
	if len(cs) != 5 {
		panic("a hand has 5 cards, this one does not??")
	}

	for _, c := range cs {
		if !strings.ContainsRune(VALID_CARDS, rune(c)) {
			panic("not valid card")
		}
	}

	return Hand(cs)
}

func countType(counts []int) int {
	if counts[0] == 5 {
		return FiveKind
	}

	if counts[0] == 1 && counts[1] == 4 {
		return FourKind
	}

	if counts[0] == 2 && counts[1] == 3 {
		return FullHouse
	}

	if counts[2] == 3 {
		return ThreeKind
	}

	if counts[2] == 2 {
		return TwoPair
	}

	if counts[3] == 2 {
		return OnePair
	}

	return HighCard
}

func (h Hand) handType() int {
	counter := map[byte]int{}
	for _, c := range h {
		if _, ok := counter[c]; !ok {
			counter[c] = 0
		}

		counter[c] += 1
	}

	counts := []int{}
	for _, c := range counter {
		counts = append(counts, c)
	}
	sort.IntSlice(counts).Sort()

	return countType(counts)
}

func (h Hand) handTypePart2() int {
	jokers := 0
	counter := map[byte]int{}
	for _, c := range h {
		if c == 'J' {
			jokers += 1
			continue
		}

		if _, ok := counter[c]; !ok {
			counter[c] = 0
		}
		counter[c] += 1
	}

	if jokers == 5 {
		return FiveKind
	}

	counts := []int{}
	for _, c := range counter {
		counts = append(counts, c)
	}
	sort.IntSlice(counts).Sort()
	counts[len(counts)-1] += jokers

	return countType(counts)
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	bids := []Bid{}
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		hand := newHand([]byte(line[0]))
		bid, _ := strconv.Atoi(line[1])
		bids = append(bids, Bid{hand, bid})
	}

	sort.Slice(bids, func(i, j int) bool {
		ha := bids[i].h.handType()
		hb := bids[j].h.handType()
		if ha != hb {
			return ha < hb
		}

		// TODO: hardcoded 5, assumed to be n. of cards in a Hand
		for idx := 0; idx < 5; idx += 1 {
			ca := cardStregnth(bids[i].h[idx])
			cb := cardStregnth(bids[j].h[idx])
			if ca != cb {
				return ca < cb
			}
		}

		return false
	})

	var total int
	for i, b := range bids {
		total += b.b * (i + 1)
	}
	return fmt.Sprintf("%d", total)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	bids := []Bid{}
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		hand := newHand([]byte(line[0]))
		bid, _ := strconv.Atoi(line[1])
		bids = append(bids, Bid{hand, bid})
	}

	sort.Slice(bids, func(i, j int) bool {
		ha := bids[i].h.handTypePart2()
		hb := bids[j].h.handTypePart2()
		if ha != hb {
			return ha < hb
		}

		// TODO: hardcoded 5, assumed to be n. of cards in a Hand
		for idx := 0; idx < 5; idx += 1 {
			ca := cardStregnthPart2(bids[i].h[idx])
			cb := cardStregnthPart2(bids[j].h[idx])
			if ca != cb {
				return ca < cb
			}
		}

		return false
	})

	var total int
	for i, b := range bids {
		total += b.b * (i + 1)
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
