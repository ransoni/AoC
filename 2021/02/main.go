package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ransoni/AoC/helpers"
)

func main() {
	input := helpers.ReadFileAsArray("data.txt")

	fmt.Printf("PART1\n")
	part1(input)

	fmt.Printf("\n\nPART2\n")
	part2(input)
}

func part2(input []string) {
	var horizontal int
	var vertical int
	var aim int = 0

	for _, steer := range input {
		command := strings.Split(steer, " ")
		direction := command[0]
		amount, err := strconv.Atoi(command[1])
		if err != nil {
			fmt.Printf("Could not convert!")
		}

		// Get the direction and amount
		switch direction {
		case "forward":
			horizontal = horizontal + amount
			vertical = vertical + (aim * amount)
		case "up":
			aim = aim - amount
		case "down":
			aim = aim + amount
		}
	}

	fmt.Printf("POSITIONS:\n")
	fmt.Printf("Horizontal: %v\n", horizontal)
	fmt.Printf("Vertical: %v\n", vertical)

	fmt.Printf("Multiply of navigation: %d\n", horizontal*vertical)
}

func part1(input []string) {
	var horizontal int
	var vertical int

	for _, steer := range input {
		command := strings.Split(steer, " ")
		direction := command[0]
		amount, err := strconv.Atoi(command[1])
		if err != nil {
			fmt.Printf("Could not convert!")
		}

		// Get the direction and amount
		switch direction {
		case "forward":
			horizontal = horizontal + amount
		case "up":
			vertical = vertical - amount
		case "down":
			vertical = vertical + amount
		}
	}

	fmt.Printf("POSITIONS:\n")
	fmt.Printf("Horizontal: %v\n", horizontal)
	fmt.Printf("Vertical: %v\n", vertical)

	fmt.Printf("Multiply of navigation: %d\n", horizontal*vertical)

}
