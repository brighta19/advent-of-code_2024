package main

// https://adventofcode.com/2024/day/2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(x string) int {
	y, _ := strconv.Atoi(x)
	return y
}

func removeAtIndex(list []string, i int) []string {
	left := list[0:i]
	right := list[i+1:]

	var new_list []string
	for _, v := range left {
		new_list = append(new_list, v)
	}
	for _, v := range right {
		new_list = append(new_list, v)
	}

	return new_list
}

func isReportSafe(levels []string) bool {
	first_level := toInt(levels[0])
	second_level := toInt(levels[1])

	var is_increasing bool
	if second_level > first_level {
		is_increasing = true
	} else if second_level < first_level {
		is_increasing = false // is decreasing
	} else {
		return false
	}

	prev_level := first_level
	for i := 1; i < len(levels); i++ {
		level := toInt(levels[i])

		ordering_changed := (is_increasing && level < prev_level) || (!is_increasing && level > prev_level)
		no_difference := level-prev_level == 0
		big_difference := (is_increasing && level-prev_level > 3) || (!is_increasing && prev_level-level > 3)

		if ordering_changed || no_difference || big_difference {
			return false
		}

		prev_level = level
	}

	return true
}

func Part1(input string) int {
	reports := strings.Split(input, "\n")
	safe_reports := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")

		if isReportSafe(levels) {
			safe_reports++
		}
	}

	return safe_reports
}

func Part2(input string) int {
	reports := strings.Split(input, "\n")
	safe_reports := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")

		for i := -1; i < len(levels); i++ {
			var modified_levels []string
			if i == -1 {
				modified_levels = levels
			} else {
				modified_levels = removeAtIndex(levels, i)
			}

			if isReportSafe((modified_levels)) {
				safe_reports++
				break
			}
		}
	}

	return safe_reports
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
