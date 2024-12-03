package main

// https://adventofcode.com/2024/day/3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func toInt(x string) int {
	y, _ := strconv.Atoi(x)
	return y
}

func Part1(input string) int {
	r, _ := regexp.Compile("mul\\((\\d{1,3},\\d{1,3})\\)") // ex. mul(123, 45)
	matches := r.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		numsStr := match[1] // ex: "45,5"
		nums := strings.Split(numsStr, ",")

		sum += toInt(nums[0]) * toInt(nums[1])
	}

	return sum
}

func Part2(input string) int {
	r, _ := regexp.Compile("(mul\\((\\d{1,3},\\d{1,3})\\)|do\\(\\)|don't\\(\\))") // ex. mul(4, 34) OR do() OR don't()
	matches := r.FindAllStringSubmatch(input, -1)

	sum := 0
	mul_activated := true
	for _, match := range matches {
		instruction := match[1]

		switch instruction {
		case "do()":
			mul_activated = true
		case "don't()":
			mul_activated = false
		default: // mul
			if mul_activated {
				numsStr := match[2] // ex: "45,5"
				nums := strings.Split(numsStr, ",")

				sum += toInt(nums[0]) * toInt(nums[1])
			}
		}
	}

	return sum
}

func main() {
	args := os.Args
	if len(args) != 3 {
		panic("usage: " + args[0] + " <input file> <part1|part2>")
	}

	input_filename, part := args[1], args[2]

	data, err := os.ReadFile(input_filename)
	if err != nil {
		panic(err)
	}

	var input string = string(data)

	if part == "part1" {
		solution := Part1(input)
		fmt.Println(solution)
	} else if part == "part2" {
		solution := Part2(input)
		fmt.Println(solution)
	} else {
		panic("unknown part. usage: " + args[0] + " <input file> <part1|part2>")
	}
}
