package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type V struct {
	x, y int
}
type Graph map[V][]V
type Grid [][]byte

// p1 top pipe, p2 bottom pipe, p1 p2 adjacent cells
func pipesConnectVert(p1, p2 byte) bool {
	if p1 == '.' || p2 == '.' {
		return false
	}

	if p1 == '-' || p1 == 'L' || p1 == 'J' {
		return false
	}

	if p2 == '-' || p2 == '7' || p2 == 'F' {
		return false
	}

	return true
}

// p1 left pipe, p2 right pipe, p1 p2 adjacent cells
func pipesConnectHoriz(p1, p2 byte) bool {
	if p1 == '.' || p2 == '.' {
		return false
	}

	if p1 == '|' || p1 == 'J' || p1 == '7' {
		return false
	}

	if p2 == '|' || p2 == 'L' || p2 == 'F' {
		return false
	}

	return true
}

// adds edges in gr to neighbouring cells to v in g if
// the pipes have connecting ends
func addConnection(v V, g Grid, gr Graph) {
	other := v
	if other.x = v.x - 1; other.x >= 0 && pipesConnectHoriz(g[other.y][other.x], g[v.y][v.x]) {
		gr[v] = append(gr[v], other)
	}

	if other.x = v.x + 1; other.x <= len(g[v.y])-1 && pipesConnectHoriz(g[v.y][v.x], g[other.y][other.x]) {
		gr[v] = append(gr[v], other)
	}

	other.x = v.x
	if other.y = v.y - 1; other.y >= 0 && pipesConnectVert(g[other.y][other.x], g[v.y][v.x]) {
		gr[v] = append(gr[v], other)
	}

	if other.y = v.y + 1; other.y <= len(g)-1 && pipesConnectVert(g[v.y][v.x], g[other.y][other.x]) {
		gr[v] = append(gr[v], other)
	}
}

type From struct {
	p    V
	dist int
}

type LoopNode struct {
	v      V
	pa, pb V
	dist   int
}

// finds a loop in gr which contains start and
// returns paths explored and the node where the loop was detected
func findLoop(start V, gr Graph) (map[V]From, LoopNode) {
	dists := map[V]From{start: From{start, 0}}
	next := []V{start}
	for len(next) > 0 {
		n := next[0]
		next = next[1:]

		for _, neigh := range gr[n] {
			f, ok := dists[neigh]
			if ok && f.dist < dists[n].dist {
				continue
			}
			if ok /* f.dist >= dists[n].dist */ {
				return dists, LoopNode{neigh, f.p, n, f.dist}
			}
			next = append(next, neigh)
			dists[neigh] = From{n, dists[n].dist + 1}
		}
	}
	return nil, LoopNode{dist: -1}
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	// parse whole grid
	g := Grid{}
	for y := 0; s.Scan(); y++ {
		line := s.Bytes()
		gl := []byte{}
		for _, b := range line {
			gl = append(gl, b)
		}
		g = append(g, gl)
	}

	// create graph from the grid pipe connections
	gr := Graph{}
	var start V
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			addConnection(V{j, i}, g, gr)
			if g[i][j] == 'S' {
				start = V{j, i}
			}
		}
	}

	_, ln := findLoop(start, gr)

	return fmt.Sprintf("%d", ln.dist)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	// parse the whole grid
	g := Grid{}
	for y := 0; s.Scan(); y++ {
		line := s.Bytes()
		gl := []byte{}
		for _, b := range line {
			gl = append(gl, b)
		}
		g = append(g, gl)
	}

	// create a graph from pipe connections in the grid
	gr := Graph{}
	var start V
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			addConnection(V{j, i}, g, gr)
			if g[i][j] == 'S' {
				start = V{j, i}
			}
		}
	}

	// find the loop and the paths leading to it
	paths, ln := findLoop(start, gr)

	// create a boundary map of the grid from the paths of the loop
	boundMap := make([][]bool, len(g))
	for i, l := range g {
		boundMap[i] = make([]bool, len(l))
	}

	g[start.y][start.x] = mark(start, paths, ln, boundMap)

	// scan boundary map and count the cells that fall inside the loop
	total := 0
	for i, l := range boundMap {
		var expectedJoint byte
		inside := false
		for j, b := range l {
			if !b {
				if inside {
					total += 1
					g[i][j] = 'X' // for visualisation
				} else {
					g[i][j] = ' '
				}
				continue
			}

			// when we cross a | pipe, we for sure either
			// go inside-out, or outside-in. But if we encounter
			// a L or F pipe, we must wait to see if we encounter
			// a 7 or J pipe (which must come after L or F to form
			// a proper loop).
			// Following cases mean we change sides:
			// 1)      |                  2)             |
			//  Side A L---7 Side B          Side A F----J Side B
			//             |                        |
			// But for following we stay on the same side.
			// 1)      |   |              2)
			//  Side A L---J Side A         Side A  F----7 Side A
			//                                      |    |

			pipe := g[i][j]
			if pipe == '-' {
				continue
			} else if pipe == '|' {
				inside = !inside
			} else if pipe == 'L' {
				expectedJoint = '7'
			} else if pipe == 'F' {
				expectedJoint = 'J'
			} else {
				if pipe == expectedJoint {
					inside = !inside
				}
				expectedJoint = 0
			}
		}
	}

	// Uncomment for visualisation
	// for _, l := range g {
	// 	fmt.Println(string(l))
	// }

	return fmt.Sprintf("%d", total)
}

// traverses the paths from loop node all the way to the start,
// marking the boundary cells along the way. Also computes the start
// pipe that completes the loop in g and returns the pipe as a result
func mark(start V, paths map[V]From, ln LoopNode, boundMap [][]bool) byte {
	boundMap[start.y][start.x] = true
	boundMap[ln.v.y][ln.v.x] = true

	var startNeighA, startNeighB V
	for next := ln.pa; next != start; next = paths[next].p {
		boundMap[next.y][next.x] = true
		startNeighA = next
	}

	for next := ln.pb; next != start; next = paths[next].p {
		boundMap[next.y][next.x] = true
		startNeighB = next
	}

	if startNeighA.y == startNeighB.y {
		return '-'
	}

	if startNeighA.x == startNeighB.x {
		return '|'
	}

	if startNeighA.x > start.x || startNeighB.x > start.x {
		if startNeighB.y < start.y || startNeighA.y < start.y {
			return 'L'
		} else {
			return 'F'
		}
	} else {
		if startNeighB.y < start.y || startNeighA.y < start.y {
			return 'J'
		} else {
			return '7'
		}
	}
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
