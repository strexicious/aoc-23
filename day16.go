package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type RayCast struct {
	x, y int
	dir  byte // 0. right, 1. up, 2. left, 3. down
}

func newCoords(x, y, dir int) (int, int) {
	if dir == 0 {
		return x + 1, y
	}
	if dir == 1 {
		return x, y - 1
	}
	if dir == 2 {
		return x - 1, y
	}
	if dir == 3 {
		return x, y + 1
	}
	panic("dir was not 0, 1, 2 or 3!")
}

func shootRay(grid [][]byte, ray RayCast, energyMap [][]byte) (newRays []RayCast) {
	w, h := len(grid[0]), len(grid)
	for {
		ray.x, ray.y = newCoords(ray.x, ray.y, int(ray.dir))
		if ray.x < 0 || ray.x >= w || ray.y < 0 || ray.y >= h {
			return
		}
		if energyMap[ray.y][ray.x]&(1<<ray.dir) != 0 {
			return
		} else {
			energyMap[ray.y][ray.x] |= 1 << ray.dir
		}
		// split if necessary and return
		if grid[ray.y][ray.x] == '|' && (ray.dir == 0 || ray.dir == 2) {
			newRays = append(newRays, RayCast{ray.x, ray.y, 1})
			newRays = append(newRays, RayCast{ray.x, ray.y, 3})
			return
		}
		if grid[ray.y][ray.x] == '-' && (ray.dir == 1 || ray.dir == 3) {
			newRays = append(newRays, RayCast{ray.x, ray.y, 0})
			newRays = append(newRays, RayCast{ray.x, ray.y, 2})
			return
		}
		if grid[ray.y][ray.x] == '/' {
			if ray.dir == 0 {
				ray.dir = 1
			} else if ray.dir == 1 {
				ray.dir = 0
			} else if ray.dir == 3 {
				ray.dir = 2
			} else if ray.dir == 2 {
				ray.dir = 3
			}
		}
		if grid[ray.y][ray.x] == '\\' {
			if ray.dir == 0 {
				ray.dir = 3
			} else if ray.dir == 3 {
				ray.dir = 0
			} else if ray.dir == 1 {
				ray.dir = 2
			} else if ray.dir == 2 {
				ray.dir = 1
			}
		}
	}
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	grid := [][]byte{}
	energyMap := [][]byte{}
	for s.Scan() {
		grid = append(grid, bytes.Clone(s.Bytes()))
		energyMap = append(energyMap, make([]byte, len(s.Bytes())))
	}

	rays := []RayCast{RayCast{-1, 0, 0}}
	for len(rays) > 0 {
		ray := rays[0]
		rays = rays[1:]
		for _, nr := range shootRay(grid, ray, energyMap) {
			rays = append(rays, nr)
		}
	}

	total := 0
	for _, l := range energyMap {
		for _, b := range l {
			if b != 0 {
				total++
			}
		}
	}

	return fmt.Sprintf("%d", total)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	grid := [][]byte{}
	for s.Scan() {
		grid = append(grid, bytes.Clone(s.Bytes()))
	}

	maxTotal := 0
	for i := 0; i < len(grid); i++ {
		for d, j := range []int{0, len(grid) - 1, len(grid[i]) - 1, 0} {
			energyMap := [][]byte{}
			for k := 0; k < len(grid); k++ {
				energyMap = append(energyMap, make([]byte, len(grid[i])))
			}

			var rays []RayCast
			if d == 0 || d == 2 {
				rays = append(rays, RayCast{j, i, byte(d)})
			} else {
				rays = append(rays, RayCast{i, j, byte(d)})
			}

			for len(rays) > 0 {
				ray := rays[0]
				rays = rays[1:]
				for _, nr := range shootRay(grid, ray, energyMap) {
					rays = append(rays, nr)
				}
			}

			total := 0
			for _, l := range energyMap {
				for _, b := range l {
					if b != 0 {
						total++
					}
				}
			}
			if total > maxTotal {
				maxTotal = total
			}
		}
	}

	return fmt.Sprintf("%d", maxTotal)
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
