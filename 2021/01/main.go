package main

import (
	"fmt"
	"strconv"

	"github.com/ransoni/AoC/helpers"
)

func main() {
	data := helpers.ReadFileAsArray("data.txt")

	increments(data)
}

func increments(data []string) {

	var inc int
	var dec int
	var prev int

	for i, k := range data {
		num, err := strconv.Atoi(k)
		if err != nil {
			fmt.Printf("Could not convert!")
		}

		if i == 0 {
			prev = num
			continue
		}

		if num > prev {
			inc++
		} else {
			dec++
		}

		prev = num
	}

	fmt.Printf("Increased: %v\n", inc)
	fmt.Printf("Decreased: %v\n", dec)
}
