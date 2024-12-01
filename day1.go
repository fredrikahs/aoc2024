package main

import (
	"fmt"
	"slices"
)

func day1(lines []string) (int, int, error) {
	var leftNumbers, rightNumbers []int

	for _, line := range lines {
		var left, right int
		_, err := fmt.Sscanf(line, "%d   %d", &left, &right)
		if err != nil {
			return -1, -1, err
		}

		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	part1, err := day1Part1(leftNumbers, rightNumbers)
	if err != nil {
		return -1, -1, err
	}

	part2, err := day1Part2(leftNumbers, rightNumbers)
	if err != nil {
		return part1, -1, err
	}

	return part1, part2, nil
}

func day1Part1(leftNumbers, rightNumbers []int) (int, error) {
	distance := 0
	for i, left := range leftNumbers {
		distance += abs(left - rightNumbers[i])
	}

	return distance, nil
}

func day1Part2(leftNumbers, rightNumbers []int) (int, error) {
	similarityScore := 0
	rightStart := 0
	for _, left := range leftNumbers {
		rightCount := 0
		for i := rightStart; i < len(rightNumbers); i++ {
			right := rightNumbers[i]
			if right > left {
				break
			}
			if left < right {
				rightStart = i
			}
			if left == right {
				rightCount++
			}
		}

		similarityScore += left * rightCount
	}

	return similarityScore, nil
}

func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}
