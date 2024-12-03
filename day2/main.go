package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filename := "day2/2.in"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}
	input := string(data)

	levels := make([][]int, 0)
	for i, line := range strings.Split(input, "\n") {
		levels = append(levels, make([]int, 0))
		for _, num := range strings.Split(strings.TrimSpace(line), " ") {
			n, _ := strconv.Atoi(num)
			levels[i] = append(levels[i], n)
		}
	}

	sum1 := 0
	for _, level := range levels {
		if IsLevelSafe(level) {
			sum1++
		}
	}
	println("Part 1:", sum1)

	sum2 := 0
	for _, level := range levels {
		if IsLevelSafe(level) {
			sum2++
		} else {
			for i := 0; i < len(level); i++ {
				c := make([]int, 0)
				for j, n := range level {
					if j != i {
						c = append(c, n)
					}
				}
				if IsLevelSafe(c) {
					sum2++
					break
				}
			}
		}
	}

	println("Part 2:", sum2)
}

func IsLevelSafe(level []int) bool {
	safe := true
	for i, num := range level {
		if i != 0 {
			diff := math.Abs(float64(num - level[i-1]))
			if diff < 1 || diff > 3 {
				safe = false
			}
		}
	}
	c := make([]int, len(level))
	copy(c, level)
	slices.Sort(c)
	inc := IntArrayEquals(c, level)
	slices.Reverse(c)
	dec := IntArrayEquals(c, level)

	if !inc && !dec {
		safe = false
	}

	return safe
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
