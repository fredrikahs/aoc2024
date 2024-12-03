package main

import (
	"regexp"
	"strconv"
)

func day3(line string) (int, int, error) {

	part1 := day3Part1(line)
	part2 := day3Part2(line)

	return part1, part2, nil
}

func day3Part1(line string) int {
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	sum := 0
	matches := r.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		m1, _ := strconv.Atoi(match[1])
		m2, _ := strconv.Atoi(match[2])
		sum += m1 * m2
	}

	return sum
}

func day3Part2(line string) int {
	sum := 0
	toggleEnabledRegex := regexp.MustCompile("(don't\\(\\))|(do\\(\\))")
	enableMatches := toggleEnabledRegex.FindAllStringIndex(line, -1)
	enabledRanges := [][]int{{0}}
	isEnabled := true
	for _, m := range enableMatches {
		if isEnabled && m[1]-m[0] == len("don't()") {
			x := len(enabledRanges) - 1
			ar := enabledRanges[x]
			ar = append(ar, m[1])
			enabledRanges[x] = ar
			isEnabled = false
		}
		if !isEnabled && m[1]-m[0] == len("do()") {
			enabledRanges = append(enabledRanges, []int{m[1]})
			isEnabled = true
		}
	}

	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	matches := r.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {

		matchEnabled := false
		for _, enabledRange := range enabledRanges {
			if match[0] > enabledRange[0] && (len(enabledRange) == 1 || match[0] < enabledRange[1]) {
				matchEnabled = true
				break
			}
		}
		if !matchEnabled {
			continue
		}

		match1 := line[match[2]:match[3]]
		m1, _ := strconv.Atoi(match1)
		match2 := line[match[4]:match[5]]
		m2, _ := strconv.Atoi(match2)
		sum += m1 * m2
	}

	return sum
}
