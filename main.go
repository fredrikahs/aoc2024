package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input/day1.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	part1, part2, err := day1(lines)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
