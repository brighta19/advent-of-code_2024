package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("usage: " + args[0] + " <input file>")
	}

	data, err := os.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	var input string = string(data)
	lines := strings.Split(input, "\n")


	var left_list []int
	var right_list []int

	for _, line := range lines {
		nums := strings.Split(line, "   ")

		left_num, _ := strconv.Atoi(nums[0])
		right_num, _ := strconv.Atoi(nums[1])

		left_list = append(left_list, left_num)
		right_list = append(right_list, right_num)
	}

	slices.Sort(left_list)
	slices.Sort(right_list)

	total := 0

	for i, left_num := range left_list {
		right_num := right_list[i]

		difference := abs(right_num - left_num)
		total += difference
	}

	fmt.Println(total)
}
