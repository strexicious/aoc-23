package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type V11 struct {
	x, y int
	pos  int
}

func part1(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	totalNodes := 0
	verts := []V11{}
	var nCols, nRows int
	for ; s.Scan(); nRows++ {
		nCols = len(s.Bytes())
		for x, b := range s.Bytes() {
			if b == '#' {
				verts = append(verts, V11{x, nRows, totalNodes})
				totalNodes++
			}
		}
	}

	colGal := make([]bool, nCols)
	rowGal := make([]bool, nRows)
	for _, v := range verts {
		colGal[v.x] = true
		rowGal[v.y] = true
	}

	adjMat := make([][]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		adjMat[i] = make([]int, totalNodes)
	}

	for i := 0; i < totalNodes; i++ {
		for j := i + 1; j < totalNodes; j++ {
			adjMat[i][j] = int(math.Abs(float64(verts[i].x-verts[j].x)) + math.Abs(float64(verts[i].y-verts[j].y)))
			adjMat[j][i] = adjMat[i][j]
		}
	}

	totalDists := 0
	// take into account expanded columns
	for i := 0; i < totalNodes; i++ {
		for j := i + 1; j < totalNodes; j++ {
			minX, maxX := verts[i].x, verts[j].x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			for _, e := range colGal[minX:maxX] {
				if !e {
					adjMat[i][j]++
					adjMat[j][i]++
				}
			}

			minY, maxY := verts[i].y, verts[j].y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			for _, e := range rowGal[minY:maxY] {
				if !e {
					adjMat[i][j]++
					adjMat[j][i]++
				}
			}
			totalDists += adjMat[i][j]
		}
	}

	return fmt.Sprintf("%d", totalDists)
}

func part2(in string) string {
	rdr := strings.NewReader(in)
	s := bufio.NewScanner(rdr)

	totalNodes := 0
	verts := []V11{}
	var nCols, nRows int
	for ; s.Scan(); nRows++ {
		nCols = len(s.Bytes())
		for x, b := range s.Bytes() {
			if b == '#' {
				verts = append(verts, V11{x, nRows, totalNodes})
				totalNodes++
			}
		}
	}

	colGal := make([]bool, nCols)
	rowGal := make([]bool, nRows)
	for _, v := range verts {
		colGal[v.x] = true
		rowGal[v.y] = true
	}

	adjMat := make([][]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		adjMat[i] = make([]int, totalNodes)
	}

	for i := 0; i < totalNodes; i++ {
		for j := i + 1; j < totalNodes; j++ {
			adjMat[i][j] = int(math.Abs(float64(verts[i].x-verts[j].x)) + math.Abs(float64(verts[i].y-verts[j].y)))
			adjMat[j][i] = adjMat[i][j]
		}
	}

	totalDists := 0
	// take into account expanded columns
	for i := 0; i < totalNodes; i++ {
		for j := i + 1; j < totalNodes; j++ {
			minX, maxX := verts[i].x, verts[j].x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			for _, e := range colGal[minX:maxX] {
				if !e {
					adjMat[i][j] += 999_999
					adjMat[j][i] += 999_999
				}
			}

			minY, maxY := verts[i].y, verts[j].y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			for _, e := range rowGal[minY:maxY] {
				if !e {
					adjMat[i][j] += 999_999
					adjMat[j][i] += 999_999
				}
			}
			totalDists += adjMat[i][j]
		}
	}

	return fmt.Sprintf("%d", totalDists)
}

func main() {
	data, err := os.ReadFile(os.Stdin.Name())
	in := string(data)

	if err != nil {
		panic("yo it crashed!")
	}

	fmt.Println(part2(in))
}
