package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	filename := "day4/4.in"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}
	input := string(data)

	g := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		g = append(g, strings.Split(line, ""))
	}

	rc, cc := len(g), len(g[0])
	inBounds := func(coords [][]int) bool {
		for _, coord := range coords {
			if coord[0] < 0 || coord[0] >= rc || coord[1] < 0 || coord[1] >= cc {
				return false
			}
		}
		return true
	}

	count1 := 0
	for r := 0; r < rc; r++ {
		for c := 0; c < cc; c++ {
			if g[r][c] != "X" {
				continue
			}

			coords := [][][]int{
				{{r, c + 1}, {r, c + 2}, {r, c + 3}},
				{{r, c - 1}, {r, c - 2}, {r, c - 3}},
				{{r + 1, c}, {r + 2, c}, {r + 3, c}},
				{{r - 1, c}, {r - 2, c}, {r - 3, c}},
				{{r + 1, c + 1}, {r + 2, c + 2}, {r + 3, c + 3}},
				{{r - 1, c - 1}, {r - 2, c - 2}, {r - 3, c - 3}},
				{{r + 1, c - 1}, {r + 2, c - 2}, {r + 3, c - 3}},
				{{r - 1, c + 1}, {r - 2, c + 2}, {r - 3, c + 3}},
			}
			for _, coord := range coords {
				if inBounds(coord) {
					c1 := g[coord[0][0]][coord[0][1]]
					c2 := g[coord[1][0]][coord[1][1]]
					c3 := g[coord[2][0]][coord[2][1]]
					if c1 == "M" && c2 == "A" && c3 == "S" {
						count1++
					}
				}
			}
		}
	}
	println("Part 1:", count1)

	count2 := 0
	for r := 0; r < rc; r++ {
		for c := 0; c < cc; c++ {
			if g[r][c] != "A" {
				continue
			}

			d1 := [][]int{{r + 1, c + 1}, {r - 1, c - 1}}
			d2 := [][]int{{r + 1, c - 1}, {r - 1, c + 1}}
			if inBounds(d1) && inBounds(d2) {
				c11 := g[d1[0][0]][d1[0][1]]
				c12 := g[d1[1][0]][d1[1][1]]
				c21 := g[d2[0][0]][d2[0][1]]
				c22 := g[d2[1][0]][d2[1][1]]

				ok1 := (c11 == "M" && c12 == "S") || (c11 == "S" && c12 == "M")
				ok2 := (c21 == "M" && c22 == "S") || (c21 == "S" && c22 == "M")
				if ok1 && ok2 {
					count2++
				}
			}
		}
	}
	println("Part 2:", count2)
}
