package main

import (
	"fmt"
	"os"
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

	keys_used := map[int]bool{}
	left_frequencies := map[int]int{}
	right_frequencies := map[int]int{}

	// if val, ok := left_right_frequencies[4]; !ok {
	// 	fmt.Println(val, ok)
    // }


	for _, line := range lines {
		nums := strings.Split(line, "   ")

		left_num, _ := strconv.Atoi(nums[0])
		right_num, _ := strconv.Atoi(nums[1])

		left_frequencies[left_num]++
		right_frequencies[right_num]++

		keys_used[left_num] = true
		keys_used[right_num] = true
	}

	similarity_score := 0

	for key := range keys_used {
		left_frequency := left_frequencies[key]
		right_frequency := right_frequencies[key]

		similarity_score += key * left_frequency * right_frequency
	}

	fmt.Println(similarity_score)
}
