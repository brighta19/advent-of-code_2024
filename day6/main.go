package main

// https://adventofcode.com/2024/day/6

import (
	"fmt"
	"os"
	"strings"
)

type lab_info struct {
	width, height int
	lab_map []string
}

func (l lab_info) at(x, y int) string {
	return string(l.lab_map[y][x])
}

type guard struct {
	x, y int
	direction string
}

func (g *guard) move() {
	switch g.direction {
	case "up":
		g.y--;
	case "down":
		g.y++;
	case "left":
		g.x--;
	case "right":
		g.x++;
	}
}

func (g *guard) turnRight() {
	switch g.direction {
	case "up":
		g.direction = "right"
	case "down":
		g.direction = "left"
	case "left":
		g.direction = "up"
	case "right":
		g.direction = "down"
	}
}

func findGuard(lab lab_info) guard {
	for y, row := range lab.lab_map {
		for x := range row {
			if lab.at(x, y) == "^" {
				return guard{x, y, "up"}
			}
		}
	}

	panic("where is the guard!?")
}

func isObstacleAhead(lab lab_info, g guard) bool {
	switch g.direction {
	case "up":
		return (g.y > 0 && lab.at(g.x, g.y - 1) == "#")
	case "down":
		return (g.y < lab.height - 1 && lab.at(g.x, g.y + 1) == "#")
	case "left":
		return (g.x > 0 && lab.at(g.x - 1, g.y) == "#")
	case "right":
		return (g.x < lab.width - 1 && lab.at(g.x + 1, g.y) == "#")
	}

	panic("what direction is that!?")
}

func isGuardInLab(lab lab_info, g guard) bool {
	return g.x >= 0 && g.x < lab.width &&
		g.y >= 0 && g.y < lab.height
}

func traceGuardPath(lab lab_info, g guard) int {
	path := make([][]bool, lab.height)
	for y := range path {
		path[y] = make([]bool, lab.width)
	}

	var numOfDistinctPositions int
	for isGuardInLab(lab, g) {
		if !path[g.y][g.x] {
			numOfDistinctPositions++
			path[g.y][g.x] = true
		}

		if isObstacleAhead(lab, g) {
			g.turnRight()
		} else {
			g.move()
		}
	}

	return numOfDistinctPositions
}

func Part1(input string) int {
	lab_map := strings.Split(input, "\n")

	lab := lab_info{
		len(lab_map[0]),
		len(lab_map),
		lab_map,
	}

	g := findGuard(lab)
	num := traceGuardPath(lab, g)
	return num
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
