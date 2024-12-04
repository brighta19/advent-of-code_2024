package main

// https://adventofcode.com/2024/day/4

import (
	"fmt"
	"os"
	"strings"
)

// Checks all eight directions for "XMAS"
func findXmas(lines []string, x, y, width, height int) (count int) {
	if x <= width - 4 {
		// left to right
		if string(lines[y][x:x+4]) == "XMAS" {
			count++
		}
		// bottom left to top right
		if y >= 3 {
			s := []byte{lines[y][x], lines[y-1][x+1], lines[y-2][x+2], lines[y-3][x+3]}
			if string(s) == "XMAS" {
				count++
			}
		}
		// top left to bottom right
		if y <= height - 4 {
			s := []byte{lines[y][x], lines[y+1][x+1], lines[y+2][x+2], lines[y+3][x+3]}
			if string(s) == "XMAS" {
				count++
			}
		}
	}
	if x >= 3 {
		// right to left
		if string(lines[y][x-3:x+1]) == "SAMX" {
			count++
		}
		// bottom right to top left
		if y >= 3 {
			s := []byte{lines[y][x], lines[y-1][x-1], lines[y-2][x-2], lines[y-3][x-3]}
			if string(s) == "XMAS" {
				count++
			}
		}
		// top right to bottom left
		if y <= height - 4 {
			s := []byte{lines[y][x], lines[y+1][x-1], lines[y+2][x-2], lines[y+3][x-3]}
			if string(s) == "XMAS" {
				count++
			}
		}
	}
	// top to bottom
	if y <= height - 4 {
		s := []byte{lines[y][x], lines[y+1][x], lines[y+2][x], lines[y+3][x]}
		if string(s) == "XMAS" {
			count++
		}
	}
	// bottom to top
	if y >= 3 {
		s := []byte{lines[y][x], lines[y-1][x], lines[y-2][x], lines[y-3][x]}
		if string(s) == "XMAS" {
			count++
		}
	}

	return count
}

// Checks both diagonal directions for X-shaped "MAS"
func findOtherXmas(lines []string, x, y, width, height int) (count int) {
	if x == 0 || x == width - 1 || y == 0 || y == height - 1 {
		return 0
	}

	// top left to bottom right
	s1 := []byte{lines[y-1][x-1], lines[y][x], lines[y+1][x+1]}
	// top right to bottom left
	s2 := []byte{lines[y-1][x+1], lines[y][x], lines[y+1][x-1]}

	if t1, t2 := string(s1), string(s2); (t1 == "MAS" || t1 == "SAM") && (t2 == "MAS" || t2 == "SAM") {
		count++
	}

	return count
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	height := len(lines)
	width := len(lines[0])
	xmasTotalCount := 0

	for y, line := range lines {
		for x := range len(line) {
			char := string(line[x])

			if char == "X" {
				count := findXmas(lines, x, y, width, height)
				xmasTotalCount += count
			}
		}
	}

	return xmasTotalCount
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	height := len(lines)
	width := len(lines[0])
	xmasTotalCount := 0

	for y, line := range lines {
		for x := range len(line) {
			char := string(line[x])

			if char == "A" {
				count := findOtherXmas(lines, x, y, width, height)
				xmasTotalCount += count
			}
		}
	}

	return xmasTotalCount
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
