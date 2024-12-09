package main

// https://adventofcode.com/2024/day/8

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

func isInMap(width, height int, a pos) bool {
	return a.x >= 0 && a.x < width && a.y >= 0 && a.y < height
}

func getAntinodes1(antenna1, antenna2 pos) []pos {
	diff := pos{
		antenna1.x - antenna2.x,
		antenna1.y - antenna2.y,
	}

	var an1 pos
	an1.x = antenna1.x + diff.x
	an1.y = antenna1.y + diff.y

	var an2 pos
	an2.x = antenna2.x - diff.x
	an2.y = antenna2.y - diff.y

	return []pos{an1, an2}
}

func getAntinodes2(antenna1, antenna2 pos, width int) []pos {
	diff := pos{
		antenna1.x - antenna2.x,
		antenna1.y - antenna2.y,
	}

	start := antenna1

	var ans []pos
	var temp pos

	temp = start
	for temp.x >= 0 && temp.x < width {
		t := temp
		ans = append(ans, t)
		temp.x += diff.x
		temp.y += diff.y
	}

	temp = pos{start.x - diff.x, start.y - diff.y}
	for temp.x >= 0 && temp.x < width {
		t := temp
		ans = append(ans, t)
		temp.x -= diff.x
		temp.y -= diff.y
	}

	return ans
}

func countAntinodes1(frequencies map[string][]pos, width, height int) int {
	antinode_map := make([][]bool, height)
	for y := range antinode_map {
		antinode_map[y] = make([]bool, width)
	}

	var count int
	for _, antennas := range frequencies {
		n := len(antennas)
		for i := 0; i < n-1; i++ {
			a1 := antennas[i]
			for j := i+1; j < n; j++ {
				a2 := antennas[j]

				ans := getAntinodes1(a1, a2)
				for _, an := range ans {
					if isInMap(width, height, an) && !antinode_map[an.y][an.x]{
						antinode_map[an.y][an.x] = true
						count++
					}
				}
			}
		}
	}

	return count
}

func countAntinodes2(frequencies map[string][]pos, width, height int) int {
	antinode_map := make([][]bool, height)
	for y := range antinode_map {
		antinode_map[y] = make([]bool, width)
	}

	var count int
	for _, antennas := range frequencies {
		n := len(antennas)
		for i := 0; i < n-1; i++ {
			a1 := antennas[i]
			for j := i+1; j < n; j++ {
				a2 := antennas[j]

				ans := getAntinodes2(a1, a2, width)
				for _, an := range ans {
					if isInMap(width, height, an) && !antinode_map[an.y][an.x]{
						antinode_map[an.y][an.x] = true
						count++
					}
				}
			}
		}
	}

	return count
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	height := len(lines)
	width := len(lines[0])
	frequencies := make(map[string][]pos)

	for y, line := range lines {
		for x, c := range line {
			if char := string(c); char != "." {
				frequencies[char] = append(frequencies[char], pos{x, y})
			}
		}
	}

	count := countAntinodes1(frequencies, width, height)
	return count
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	height := len(lines)
	width := len(lines[0])
	frequencies := make(map[string][]pos)

	for y, line := range lines {
		for x, c := range line {
			if char := string(c); char != "." {
				frequencies[char] = append(frequencies[char], pos{x, y})
			}
		}
	}

	count := countAntinodes2(frequencies, width, height)
	return count
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
