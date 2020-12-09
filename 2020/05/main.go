package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	// validator "gopkg.in/validator.v2"
)

const (
	fname string = "data.txt"
)

var (
	// count   int
	valid        int64
	unvalid      int64
	strictPolicy int64
)

type data struct {
	array []string
	str   string
}

type seat struct {
	raw       string
	rawBinary string
	row       int64
	seat      int64
	ID        int64
}

type allSeats struct {
	seat []seat
}

var (
	seats = new(allSeats)
)

func main() {
	d := new(data)
	d.str = readFileToString(fname)
	d.array = readFileAsArray(fname)

	d.str = strings.ReplaceAll(d.str, "F", "0")
	d.str = strings.ReplaceAll(d.str, "B", "1")
	d.str = strings.ReplaceAll(d.str, "L", "0")
	d.str = strings.ReplaceAll(d.str, "R", "1")

	scanner := bufio.NewScanner(strings.NewReader(d.str))
	for scanner.Scan() {
		d.array = append(d.array, scanner.Text())
		fmt.Println(scanner.Text())
	}

	// os.Exit(0)

	// Sum of the rows
	n := 128
	total := int64(((n + 1) * (n + 2) / 2) * 8)
	fmt.Printf("Total: %d\n", total)

	var idTotal int64
	for r := 0; r < 128; r++ {
		for s := 0; s < 8; s++ {
			idTotal = int64(idTotal + int64(r*8+s))
			fmt.Printf("ID: %d\n", idTotal)
		}
	}
	// os.Exit(0)

	maxID := int64(0)
	oneSeat := new(seat)
	for _, binarySeat := range d.array {
		oneSeat.rawBinary = binarySeat
		fmt.Printf("RAW Row: %s, Seat: %s\n", binarySeat[:7], binarySeat[7:])
		oneSeat.row, _ = strconv.ParseInt(oneSeat.rawBinary[:7], 2, 64)
		oneSeat.seat, _ = strconv.ParseInt(oneSeat.rawBinary[7:], 2, 64)
		oneSeat.ID = int64(oneSeat.row*8 + oneSeat.seat)

		if oneSeat.ID > maxID {
			maxID = oneSeat.ID
		}

		idTotal = idTotal - oneSeat.ID

		seats.seat = append(seats.seat, *oneSeat)
		fmt.Printf("Row: %d, Seat: %d\n", oneSeat.row, oneSeat.seat)

		total = total - (oneSeat.row + 1)
	}

	fmt.Printf("Max ID: %d\n", maxID)
	// fmt.Printf("Missing ID: %d\n", idTotal)

}

func (d *data) cleanData() {
	fmt.Println("Cleaning up...")

	// j := 0
	// for _, line := range d.text {

	// }
}

func reSplit(text string, pattern string) (out []string) {
	re := regexp.MustCompile(pattern)
	out = re.Split(text, -1)

	return
}

func readFileToString(fname string) string {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

func readFileAsArray(fname string) []string {
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
