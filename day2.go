package main

import (
	"fmt"
	"slices"
	"strings"
)

func day2(lines []string) (int, int, error) {
	var reports [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}
		splits := strings.Split(line, " ")
		var levels []int
		for _, split := range splits {
			var x int
			_, err := fmt.Sscanf(split, "%d", &x)
			if err != nil {
				return -1, -1, err
			}

			levels = append(levels, x)
		}

		reports = append(reports, levels)
	}

	part1 := day2Part1(reports)
	part2 := day2Part2(reports)

	return part1, part2, nil
}

func day2Part1(reports [][]int) int {
	count := 0
	for _, levels := range reports {
		if isLevelsSafe(levels) {
			count++
		}
	}

	return count
}

func day2Part2(reports [][]int) int {
	count := 0
	for _, levels := range reports {
		for i := 0; i < len(levels); i++ {
			dampened := slices.Concat(levels[:i], levels[i+1:len(levels)])
			if isLevelsSafe(dampened) {
				count++
				break
			}
		}

	}

	return count
}

func isLevelsSafe(levels []int) bool {
	var isIncreasing bool
	for i, curr := range levels {
		if i == 0 {
			isIncreasing = curr < levels[i+1]
			continue
		}

		prev := levels[i-1]
		diff := curr - prev

		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if diff < 0 == isIncreasing {
			return false
		}
	}

	return true
}
