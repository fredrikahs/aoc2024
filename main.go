package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("input/day3.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	// lines := strings.Split(string(file), "\n")

	part1, part2, err := day3(string(file))

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
