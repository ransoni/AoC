package main

import (
	"fmt"
	"strconv"

	"github.com/ransoni/AoC/helpers"
)

func main() {
	input := helpers.ReadFileAsArray("data.txt")

	var gamma string
	var epsilon string

	for i := range input[0] {
		var sum int = 0
		for _, c := range input {
			bitNum, _ := strconv.Atoi(string(c[i]))
			sum = sum + bitNum
		}
		if sum > len(input)/2 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		fmt.Printf("Converting bit to Int FAILED!\n")
	}
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		fmt.Printf("Converting bit to Int FAILED!\n")
	}
	fmt.Printf("GAMMA: %s (%d)\n", gamma, gammaInt)
	fmt.Printf("EPSILON: %s (%d)\n\n", epsilon, epsilonInt)
	fmt.Printf("ANSWER: %d\n", gammaInt*epsilonInt)
}
