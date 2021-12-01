package main

import (
	"fmt"

	// validator "gopkg.in/validator.v2"
	// helper "github.com/ransoni/aoc/2020/helpers"

	"aoc2020.com/helpers"
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
	input := helpers.ReadFileAsArray(fname)

	fmt.Println(input)

}

// func (d *data) cleanData() {
// 	fmt.Println("Cleaning up...")

// 	// j := 0
// 	// for _, line := range d.text {

// 	// }
// }

// func reSplit(text string, pattern string) (out []string) {
// 	re := regexp.MustCompile(pattern)
// 	out = re.Split(text, -1)

// 	return
// }

// func readFileToString(fname string) string {
// 	content, err := ioutil.ReadFile(fname)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Convert []byte to string and print to screen
// 	text := string(content)
// 	return text
// }

// func readFileAsArray(fname string) []string {
// 	file, err := os.Open(fname)

// 	if err != nil {
// 		log.Fatalf("failed to open")

// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	scanner.Split(bufio.ScanLines)
// 	var text []string

// 	for scanner.Scan() {
// 		text = append(text, scanner.Text())
// 	}

// 	file.Close()

// 	return text
// }
