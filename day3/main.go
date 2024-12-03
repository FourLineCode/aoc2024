package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filename := "day3/3.in"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}
	input := string(data)

	sum1 := Compute(input)
	println("Part 1:", sum1)

	for strings.Contains(input, "don't()") {
		dont_i := strings.Index(input, "don't()")
		do_i := len(input)
		res := strings.Index(input[dont_i:], "do()")
		if res != -1 {
			do_i = len(input[:dont_i]) + res
		}
		input = input[:dont_i] + input[do_i:]
	}

	sum2 := Compute(input)
	println("Part 2:", sum2)
}

func Compute(program string) int {
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	sum := 0
	for _, match := range r.FindAllString(program, -1) {
		s := strings.TrimRight(strings.TrimLeft(match, "mul("), ")")
		nums := strings.Split(s, ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		sum += num1 * num2
	}
	return sum
}
