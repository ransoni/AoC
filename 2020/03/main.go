package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	fname = "data.txt"
	// count   int
	valid        int64
	unvalid      int64
	strictPolicy int64
)

type data struct {
	text []string
}

type pointer struct {
	col, line int
}

type counter struct {
	open, tree int
}

var (
	d = new(data)
	s = new(data)
	p = pointer{0, 0}
)

func main() {
	// d := new(data)
	// s := new(data)
	// p := pointer{0, 0}

	runSlope(1, 2)

}

func runSlope(col int, line int) {
	slope := pointer{col: col, line: line}
	count := counter{open: 0, tree: 0}

	text := readFile(fname)
	s.text = text
	d.text = text

	for _, line := range d.text {
		for i, chr := range line {
			if i == p.col {
				if chr == '.' {
					// fmt.Println("Open slot!")
					fmt.Printf("O")
					count.open++
				} else {
					// fmt.Printl("YOU HIT TREE!")
					fmt.Printf("X")
					count.tree++
				}
			} else {
				fmt.Printf("%c", chr)
			}
			if i+1 == len(line) {
				fmt.Printf(" (%c)", chr)
			}
		}

		// if i > 20 {
		// 	break
		// }

		p.col = p.col + slope.col
		p.line = p.line + slope.line

		if ((p.col + slope.col) >= len(line)) && len(d.text) > p.line {
			fmt.Printf(" Running out of lines!")
			d.addColumn()
		}

		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")

	fmt.Printf("OPEN: %d, TREE: %d\n", count.open, count.tree)
	fmt.Println("END")
}

func (d *data) addColumn() {
	for i, line := range d.text {
		// fmt.Println(line)
		d.text[i] = line + s.text[i]
		// fmt.Println(d.text[i])
	}
	// return out
}

func readFile(fname string) []string {
	file, err := os.Open(fname)

	if err != nil {
		log.Fatalf("failed to open")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	return text
}
