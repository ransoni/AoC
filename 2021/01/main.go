package main

import (
	"fmt"

	"github.com/ransoni/AoC/helpers"
)

func main() {
	// data := helpers.ReadFileAsArray("data.txt")
	data := helpers.ReadFileAsIntArray("data.txt")

	fmt.Printf("\nOne by One Results:\n")
	increments(data)

	fmt.Printf("\nSliding Window Results:\n")
	slidingWindow(data)
}

func slidingWindow(data []int) {
	var inc int
	var dec int

	for i := range data {
		// fmt.Printf("Len of window: %d", len(data[i:]))
		if len(data[i:]) < 3 {
			// fmt.Printf("Not enough rows available!\n")
			continue
		}

		prevWindow := countWindow(data[i : i+3])
		nextWindow := countWindow(data[i+1 : i+4])
		if nextWindow > prevWindow {
			// fmt.Printf("Its BIGGER!\n")
			// fmt.Printf("Prev: %v || %v :Next\n", prevWindow, nextWindow)
			inc++
		} else {
			// fmt.Printf("Its smaller!\n")
			// fmt.Printf("Prev: %v || %v :Next\n", prevWindow, nextWindow)
			dec++
		}
	}

	fmt.Printf("Increased: %v\n", inc)
	fmt.Printf("Decreased: %v\n", dec)

	// return
}

func countWindow(window []int) int {
	var sum int

	for _, num := range window {
		sum = sum + num
	}

	return sum
}

func increments(data []int) {

	var inc int
	var dec int
	var prev int

	for i, num := range data {
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
