package main

// https://adventofcode.com/2024/day/7

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func toInt(x string) int {
	y, _ := strconv.Atoi(x)
	return y
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func isEquationTrue(testValue int, nums []string) bool{
	numOfOperations := len(nums) - 1
	numOfArrangements := pow(2, numOfOperations)
	for a := 0; a < numOfArrangements; a++ {
		total := toInt(nums[0])
		for i := 0; i < numOfOperations; i++ {
			num := toInt(nums[i+1])
			// uses binary to determine which operation to use
			if (a & pow(2, i)) == 0 {
				total += num
			} else {
				total *= num
			}
		}

		if total == testValue {
			return true
		}
	}

	return false
}

func Part1(input string) int {
	equations := strings.Split(input, "\n")

	var sum int
	for _, equation := range equations {
		equationSplit := strings.Split(equation, ":")
		testValue := toInt(equationSplit[0])
		nums := strings.Split(equationSplit[1][1:], " ")

		if isEquationTrue(testValue, nums) {
			sum += testValue
		}
	}

	return sum
}

func Part2(input string) int {
	panic("part 2 not done!")
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
