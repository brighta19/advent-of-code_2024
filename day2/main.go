package main

// https://adventofcode.com/2024/day/2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(input string) int {
	reports := strings.Split(input, "\n")
	safe_reports := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")

		first_level, _ := strconv.Atoi(levels[0])
		second_level, _ := strconv.Atoi(levels[1])

		var is_increasing bool
		if second_level > first_level {
			is_increasing = true
		} else if second_level < first_level {
			is_increasing = false
		} else {
			continue // unsafe
		}

		prev_level := first_level
		safe := true
		for i := 1; i < len(levels); i++ {
			level, _ := strconv.Atoi(levels[i])

			ordering_changed := (is_increasing && level < prev_level) || (!is_increasing && level > prev_level)
			no_difference := level - prev_level == 0
			big_difference := (is_increasing && level - prev_level > 3) || (!is_increasing && prev_level - level > 3)

			if ordering_changed || no_difference || big_difference {
				safe = false
				break
			}

			prev_level = level
		}

		if safe {
			safe_reports++
		}
	}
	return safe_reports
}

func Part2(input string) int {
	reports := strings.Split(input, "\n")
	safe_reports := 0

	for _, report := range reports {
		levels_2 := strings.Split(report, " ")

		for i := -1; i < len(levels_2); i++ {
			var levels []string
			if i == -1 {
				levels = levels_2
			} else {
				left := levels_2[0:i]
				right := levels_2[i+1:]
				for _, v := range left {
					levels = append(levels, v)
				}
				for _, v := range right {
					levels = append(levels, v)
				}
			}

			first_level, _ := strconv.Atoi(levels[0])
			second_level, _ := strconv.Atoi(levels[1])

			var is_increasing bool
			if second_level > first_level {
				is_increasing = true
			} else if second_level < first_level {
				is_increasing = false
			} else {
				continue // unsafe
			}

			prev_level := first_level
			safe := true
			for i := 1; i < len(levels); i++ {
				level, _ := strconv.Atoi(levels[i])

				ordering_changed := (is_increasing && level < prev_level) || (!is_increasing && level > prev_level)
				no_difference := level - prev_level == 0
				big_difference := (is_increasing && level - prev_level > 3) || (!is_increasing && prev_level - level > 3)

				if ordering_changed || no_difference || big_difference {
					safe = false
					break
				}

				prev_level = level
			}

			if safe {
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
