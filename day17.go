package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type V17K struct {
	x, y  int
	dir   byte // 0. right, 1. up, 2. left, 4. down
	sleft int
}

func dV17K(x, y int) V17K {
	return V17K{x, y, 255, 1000}
}

func buildGraph(grid [][]byte) map[V17K][]V17K {
	w, h := len(grid[0]), len(grid)

	g := map[V17K][]V17K{}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if i != 0 {
				g[dV17K(i, j)] = append(g[dV17K(i, j)], dV17K(i-1, j))
			}
			if j != 0 {
				g[dV17K(i, j)] = append(g[dV17K(i, j)], dV17K(i, j-1))
			}
			if i != w-1 {
				g[dV17K(i, j)] = append(g[dV17K(i, j)], dV17K(i+1, j))
			}
			if j != h-1 {
				g[dV17K(i, j)] = append(g[dV17K(i, j)], dV17K(i, j+1))
			}
		}
	}

	return g
}

func getDir(a, b V17K) byte {
	if a.y == b.y {
		if a.x < b.x {
			return 0
		} else {
			return 2
		}
	}
	if a.x == b.x {
		if a.y > b.y {
			return 1
		} else {
			return 3
		}
	}

	panic("should not happen")
}

func shortPaths(grid [][]byte, g map[V17K][]V17K, visited map[V17K]bool, from V17K, cache map[V17K]int, paths map[V17K]V17K) int {
	w, h := len(grid[0]), len(grid)
	if from.x == w-1 && from.y == h-1 {
		return int(grid[from.y][from.x])
	}

	if v, ok := cache[from]; ok {
		return v
	}

	visited[from] = true
	minHeatLoss := int(1e6)
	for _, nv := range g[dV17K(from.x, from.y)] {
		n := nv
		n.dir = getDir(from, n)
		if (n.dir-from.dir+4)%4 == 2 {
			continue
		}
		if n.dir == from.dir {
			if from.sleft == 0 {
				continue
			} else {
				n.sleft = from.sleft - 1
			}
		} else {
			n.sleft = 2
		}

		if k, ok := visited[n]; ok && k {
			continue
		}

		if v := shortPaths(grid, g, visited, n, cache, paths); v < minHeatLoss {
			minHeatLoss = v
			paths[from] = n
		}
	}
	visited[from] = false

	if minHeatLoss < 1e6 {
		cache[from] = minHeatLoss + int(grid[from.y][from.x])
		return cache[from]
	} else {
		return minHeatLoss
	}
}

func shortPathsPart2(grid [][]byte, g map[V17K][]V17K, visited map[V17K]bool, from V17K, cache map[V17K]int, paths map[V17K]V17K) (int, bool) {
	w, h := len(grid[0]), len(grid)
	if from.x == w-1 && from.y == h-1 {
		return int(grid[from.y][from.x]), true
	}

	if v, ok := cache[from]; ok {
		return v, true
	}

	visited[from] = true
	minHeatLoss := int(1e6)
	found := false
	for _, nv := range g[dV17K(from.x, from.y)] {
		n := nv
		n.dir = getDir(from, n)
		if (n.dir+4-from.dir)%4 == 2 {
			continue
		}
		if n.dir == from.dir && from.sleft == 10 {
			continue
		}
		if from.sleft < 4 && n.dir != from.dir {
			continue
		}

		if n.dir == from.dir {
			n.sleft = from.sleft + 1
		} else {
			n.sleft = 1
		}

		if k, ok := visited[n]; ok && k {
			continue
		}

		if v, ok := shortPathsPart2(grid, g, visited, n, cache, paths); ok && v < minHeatLoss {
			minHeatLoss = v
			paths[from] = n
			found = true
		}
	}
	visited[from] = false

	if found {
		cache[from] = minHeatLoss + int(grid[from.y][from.x])
		return cache[from], true
	} else {
		return minHeatLoss, false
	}
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	grid := [][]byte{}
	for i := 0; s.Scan(); i++ {
		grid = append(grid, []byte{})
		for _, b := range s.Bytes() {
			grid[i] = append(grid[i], b-'0')
		}
	}

	g := buildGraph(grid)
	cache := map[V17K]int{}
	visited := map[V17K]bool{}
	paths := map[V17K]V17K{}
	minLoss := shortPaths(grid, g, visited, V17K{0, 0, 0, 3}, cache, paths) - int(grid[0][0])

	// stepMap := [][]byte{}
	// for i := range grid {
	// 	stepMap = append(stepMap, make([]byte, len(grid[i])))
	// 	for j := range grid[i] {
	// 		stepMap[i][j] = '.'
	// 	}
	// }

	// d2c := map[byte]byte{
	// 	0: '>',
	// 	1: '^',
	// 	2: '<',
	// 	3: 'v',
	// }
	// fmt.Println(cache[V17K{5, 0, 1, 2}])
	// fmt.Println(cache[V17K{5, 0, 3, 2}])
	// for p, ok := paths[V17K{0, 0, 0, 3}]; ok; p, ok = paths[p] {
	// 	stepMap[p.y][p.x] = d2c[p.dir]
	// }

	// for _, l := range stepMap {
	// 	fmt.Println(string(l))
	// }

	return fmt.Sprintf("%d", minLoss)
}

func part2(in string) string {
	// 	rdr := strings.NewReader(in)
	// 	s := bufio.NewScanner(rdr)

	// 	grid := [][]byte{}
	// 	for i := 0; s.Scan(); i++ {
	// 		grid = append(grid, []byte{})
	// 		for _, b := range s.Bytes() {
	// 			grid[i] = append(grid[i], b-'0')
	// 		}
	// 	}

	// 	g := buildGraph(grid)
	// 	cache := map[V17K]int{}
	// 	visited := map[V17K]bool{}
	// 	paths := map[V17K]V17K{}
	// 	minLoss, _ := shortPathsPart2(grid, g, visited, V17K{0, 0, 3, 0}, cache, paths)
	// 	minLoss -= int(grid[0][0])

	// 	stepMap := [][]byte{}
	// 	for i := range grid {
	// 		stepMap = append(stepMap, make([]byte, len(grid[i])))
	// 		for j := range grid[i] {
	// 			stepMap[i][j] = '.'
	// 		}
	// 	}

	// 	d2c := map[byte]byte{
	// 		0: '>',
	// 		1: '^',
	// 		2: '<',
	// 		3: 'v',
	// 	}
	// 	fmt.Println(cache[V17K{5, 0, 1, 2}])
	// 	fmt.Println(cache[V17K{5, 0, 3, 2}])
	// 	for p, ok := paths[V17K{0, 0, 3, 0}]; ok; p, ok = paths[p] {
	// 		stepMap[p.y][p.x] = d2c[p.dir]
	// 	}

	// 	for _, l := range stepMap {
	// 		fmt.Println(string(l))
	// 	}

	// 	return fmt.Sprintf("%d", minLoss)

	return "it really was djikstra after all"
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
