package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ransoni/AoC/helpers"
)

func main() {
	input := helpers.ReadFileAsArray("data.txt")

	nums_raw := strings.Split(input[0], ",")
	numbers := make([]int, len(nums_raw))

	for i, n := range nums_raw {
		numbers[i], _ = strconv.Atoi(n)
	}

	boards_raw := input[2:]

	var boards [][5][5]int

	for r := 0; r < len(boards_raw); r += 6 {
		var board [5][5]int
		// fmt.Printf("\n\n")
		for rr := 0; rr < 5; rr++ {
			fmt.Sscanf(boards_raw[r+rr], "%d %d %d %d %d", &board[rr][0], &board[rr][1], &board[rr][2], &board[rr][3], &board[rr][4])
			// fmt.Printf("%v\n", board[rr])
		}
		boards = append(boards, board)
	}

	fmt.Printf("Num of boards: %v\n", len(boards))
	fmt.Printf("Board #1:\n%v", boards[0])

	var hits [100][5][5]bool
	var bingo bool = false
	var winningBoard, line, num, sum int
	var lastWinningBoard, lastNum int
	var lastBoard int = len(boards)

	for _, n := range numbers {
		for i, board := range boards {
			for r := 0; r < 5; r++ {
				// var lineSum int
				for c := 0; c < 5; c++ {
					if board[r][c] == n {
						hits[i][r][c] = true
					}
					// lineSum = lineSum + board[r][c]
				}
				if hits[i][r][0] == true && hits[i][r][1] == true && hits[i][r][2] == true && hits[i][r][3] == true && hits[i][r][4] == true {
					// if bingo {
					fmt.Println("BINGO!")
					bingo = true
					winningBoard = i
					line = r
					num = n
					lastBoard = lastBoard - 1
					// sum = lineSum
					break
					// } else {
					// 	// fmt.Println("BINGO!")
					// 	// bingo = true
					// 	lastWinningBoard = i
					// 	// lastLine = r
					// 	lastNum = n
					// }
				}
			}
			if bingo {
				break
			}
		}
		if bingo {
			break
		}
	}

	for r := 0; r < 5; r++ {
		fmt.Printf("%v", boards[winningBoard][r])
		fmt.Printf(" \t[")
		for c := 0; c < 5; c++ {
			if hits[winningBoard][r][c] == false {
				sum = sum + boards[winningBoard][r][c]
				fmt.Printf("o ")
			} else {
				fmt.Printf("X ")
			}
		}
		fmt.Printf("]\n")
	}

	fmt.Printf("Winning line is: %v\n", boards[winningBoard][line])
	fmt.Printf("Board: %d, Line: %d, Sum: %d\n", winningBoard, line, sum)
	total := sum * num
	fmt.Printf("ANSWER: %d\n", total)

	fmt.Printf("--- PART 2 ---\n")
	printBoard(boards[lastWinningBoard], hits[lastWinningBoard], lastNum)
}

func printBoard(board [5][5]int, hits [5][5]bool, lastNum int) {
	var sum int = 0
	for r := 0; r < 5; r++ {
		fmt.Printf("%v", board[r])
		fmt.Printf(" \t[")
		for c := 0; c < 5; c++ {
			if hits[r][c] == false {
				sum = sum + board[r][c]
				fmt.Printf("o ")
			} else {
				fmt.Printf("X ")
			}
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("TOTAL: %d\n", sum*lastNum)
}
