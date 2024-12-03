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
	filename := "day1/1.in"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}
	input := string(data)

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		ns := strings.Fields(line)
		n1, _ := strconv.Atoi(ns[0])
		n2, _ := strconv.Atoi(ns[1])
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	if len(list1) != len(list2) {
		log.Fatalf("lists are not of the same length")
	}

	sum1 := 0
	for i, n1 := range list1 {
		n2 := list2[i]
		sum1 += int(math.Abs(float64(n1 - n2)))
	}
	println("Part 1:", sum1)

	m := make(map[int]int)
	for _, n := range list2 {
		if _, ok := m[n]; !ok {
			m[n] = 0
		}
		m[n] += 1
	}

	sum2 := 0
	for _, n1 := range list1 {
		if _, ok := m[n1]; ok {
			sum2 += n1 * m[n1]
		}
	}
	println("Part 2:", sum2)
}
