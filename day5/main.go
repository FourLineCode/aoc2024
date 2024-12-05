package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filename := "day5/5.in"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}
	input := string(data)

	slices_s := strings.Split(input, "\n\n")
	rules_s, updates_s := strings.Split(slices_s[0], "\n"), strings.Split(slices_s[1], "\n")

	rules, updates := make([][2]int, 0), make([][]int, 0)
	for _, rule := range rules_s {
		r := strings.Split(rule, "|")
		n1, _ := strconv.Atoi(r[0])
		n2, _ := strconv.Atoi(r[1])
		rules = append(rules, [2]int{n1, n2})
	}
	for _, update := range updates_s {
		u := strings.Split(update, ",")
		update := make([]int, 0)
		for _, s := range u {
			n, _ := strconv.Atoi(s)
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	sum1, sum2 := 0, 0
	for _, update := range updates {
		ok, flag := true, true
		corrected := make([]int, len(update))
		for flag {
			flag = false
			for _, rule := range rules {
				i1, i2 := slices.Index(update, rule[0]), slices.Index(update, rule[1])
				if i1 != -1 && i2 != -1 && i1 > i2 {
					ok, flag = false, true
					update := append(update[:i1], update[i1+1:]...)
					rest := slices.Clone(update[i2:])
					update = append(append(update[:i2], rule[0]), rest...)
					corrected = update
				}
			}
		}
		if ok {
			sum1 += update[len(update)/2]
		} else {
			sum2 += corrected[len(corrected)/2]
		}
	}
	println("Part 1:", sum1)
	println("Part 2:", sum2)
}
