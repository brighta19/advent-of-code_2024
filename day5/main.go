package main

// https://adventofcode.com/2024/day/5

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

func swapStrings(a, x, y string) string {
	a = strings.Replace(a, x, "temp", -1)
	a = strings.Replace(a, y, x, -1)
	a = strings.Replace(a, "temp", y, -1)
	return a
}

func getRules(section string) [][]string {
	var rulesStr = strings.Split(section, "\n")

	rules := make([][]string, len(rulesStr))

	for i, line := range rulesStr {
		nums := strings.Split(line, "|")
		rules[i] = append([]string{}, nums[0], nums[1])
	}

	return rules
}

// Returns false if an update breaks a rule
func isUpdateCorrect(update string, rules [][]string) (correct bool, brokenRules [][]string) {
	for _, rule := range rules {
		index0 := strings.Index(update, rule[0])
		index1 := strings.Index(update, rule[1])

		if index0 < 0 || index1 < 0 {
			continue
		}

		if index0 > index1 {
			brokenRules = append(brokenRules, rule)
		}
	}

	return len(brokenRules) == 0, brokenRules
}

func fixUpdate(update string, brokenRule []string) string {
	update = swapStrings(update, brokenRule[0], brokenRule[1])
	return update
}

func Part1(input string) int {
	sections := strings.Split(input, "\n\n")
	rules := getRules(sections[0])
	updates := strings.Split(sections[1], "\n")

	var total int
	for _, update := range updates {
		if correct, _ := isUpdateCorrect(update, rules); correct {
			nums := strings.Split(update, ",")
			middleIndex := len(nums) / 2
			middleNum := toInt(nums[middleIndex])
			total += middleNum
		}
	}

	return total
}

func Part2(input string) int {
	sections := strings.Split(input, "\n\n")
	rules := getRules(sections[0])
	updates := strings.Split(sections[1], "\n")

	var total int
	for _, update := range updates {
		if correct, brokenRules := isUpdateCorrect(update, rules); !correct {
			update = fixUpdate(update, brokenRules[0]);
			correct, brokenRules = isUpdateCorrect(update, rules)

			for !correct {
				update = fixUpdate(update, brokenRules[0]);
				correct, brokenRules = isUpdateCorrect(update, rules)
			}

			nums := strings.Split(update, ",")
			middleIndex := len(nums) / 2
			middleNum := toInt(nums[middleIndex])
			total += middleNum
		}
	}

	return total
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
