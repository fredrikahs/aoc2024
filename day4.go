package main

func day4(lines []string) (int, int, error) {
	part1 := day4Part1(lines)
	part2 := day4Part2(lines)

	return part1, part2, nil
}

func day4Part1(lines []string) int {
	count := 0

	for y, line := range lines {
		for x, ch := range line {
			if ch != 'X' {
				continue
			}
			if x+3 < len(line) && line[x+1] == 'M' && line[x+2] == 'A' && line[x+3] == 'S' {
				count++
			}
			if x-3 >= 0 && line[x-1] == 'M' && line[x-2] == 'A' && line[x-3] == 'S' {
				count++
			}
			if y+3 < len(lines) && lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
				count++
			}
			if y-3 >= 0 && lines[y-1][x] == 'M' && lines[y-2][x] == 'A' && lines[y-3][x] == 'S' {
				count++
			}
			if y-3 >= 0 && x-3 >= 0 && lines[y-1][x-1] == 'M' && lines[y-2][x-2] == 'A' && lines[y-3][x-3] == 'S' {
				count++
			}
			if y-3 >= 0 && x+3 < len(line) && lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S' {
				count++
			}
			if y+3 < len(lines) && x-3 >= 0 && lines[y+1][x-1] == 'M' && lines[y+2][x-2] == 'A' && lines[y+3][x-3] == 'S' {
				count++
			}
			if y+3 < len(lines) && x+3 < len(line) && lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S' {
				count++
			}
		}
	}

	return count
}

func day4Part2(lines []string) int {
	count := 0

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			if lines[y][x] != 'A' {
				continue
			}

			nw := lines[y-1][x-1]
			ne := lines[y-1][x+1]
			sw := lines[y+1][x-1]
			se := lines[y+1][x+1]

			diag1 := (nw == 'M' && se == 'S') || (nw == 'S' && se == 'M')
			diag2 := (ne == 'M' && sw == 'S') || (ne == 'S' && sw == 'M')
			if diag1 && diag2 {
				count++
			}
		}
	}

	return count
}
