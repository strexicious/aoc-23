package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// source for GCD and LCM: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type Node struct {
	left, right string
}

func (n Node) move(dir byte) string {
	if dir == 'L' {
		return n.left
	}

	if dir == 'R' {
		return n.right
	}

	panic("invalid move instruction")
}

func part1(in string) string {
	splits := strings.Split(in, "\n\n")
	moveInsts := splits[0]

	rdr := strings.NewReader(splits[1])
	s := bufio.NewScanner(rdr)

	network := map[string]Node{}
	for s.Scan() {
		// example: line = "NFK = (LMH, RSS)"
		line := s.Text()

		currNodeId := line[:3]
		leftNodeId := line[7:10]
		rightNodeId := line[12:15]

		network[currNodeId] = Node{leftNodeId, rightNodeId}
	}

	steps := 1
	currNode := network["AAA"].move(moveInsts[0])
	for currNode != "ZZZ" {
		currNode = network[currNode].move(moveInsts[steps%len(moveInsts)])
		steps += 1
	}
	return fmt.Sprintf("%d", steps)
}

type CycleDesc struct {
	tillZs []int

	// index of tillZs to which loop back to after
	// last element of tillZs
	cycleIdx int
}

// This function computes n. of steps needed starting from `start` to reach
// a node ending with 'Z'. After that, it recursively computes n. of steps
// needed to reach next node ending with 'Z'. It does so until it reaches
// a node ending with 'Z' at which it's been before, and with the same
// next instruction to follow.
func computeZSteps(start string, network map[string]Node, insts string) CycleDesc {
	type VisitInfo struct {
		nodeId  string
		instIdx int
	}

	desc := CycleDesc{}
	visited := map[VisitInfo]int{}

	steps := 0
	instIdx := 0
	currNode := start
	for {
		if currNode[2] == 'Z' {
			desc.tillZs = append(desc.tillZs, steps)
			steps = 0

			if idx, ok := visited[VisitInfo{currNode, instIdx}]; ok {
				desc.cycleIdx = idx
				break
			}
			visited[VisitInfo{currNode, instIdx}] = len(desc.tillZs) - 1
		}

		currNode = network[currNode].move(insts[instIdx])
		instIdx = (instIdx + 1) % len(insts)
		steps += 1
	}

	return desc
}

func part2(in string) string {
	splits := strings.Split(in, "\n\n")
	moveInsts := splits[0]

	rdr := strings.NewReader(splits[1])
	s := bufio.NewScanner(rdr)

	network := map[string]Node{}
	for s.Scan() {
		// example: line = "NFK = (LMH, RSS)"
		line := s.Text()

		currNodeId := line[:3]
		leftNodeId := line[7:10]
		rightNodeId := line[12:15]

		network[currNodeId] = Node{leftNodeId, rightNodeId}
	}

	uniqueNSteps := []int{}
	for k, _ := range network {
		if k[2] == 'A' {
			cycleDesc := computeZSteps(k, network, moveInsts)
			// WARNING: Now, at least with the test input, and my input,
			// `cycleDesc.tillZs` elements were all same. So we could
			// just say that starting from `k`, we need to take this unique
			// number of steps to reach a Z-node, and same n. of steps
			// next time as well to reach a Z-node again. So to solve the
			// minimum number of steps needed for all A-nodes to finish
			// at Z-nodes, we need to find the LCM(n. of steps for all A-nodes)
			uniqueNSteps = append(uniqueNSteps, cycleDesc.tillZs[0])
		}
	}

	return fmt.Sprintf("%d", LCM(uniqueNSteps[0], uniqueNSteps[1], uniqueNSteps[2:]...))
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
