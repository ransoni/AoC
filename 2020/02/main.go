package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	fname = "data.txt"
	// count   int
	valid        int64
	unvalid      int64
	strictPolicy int64
)

func main() {

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

	for _, line := range text {
		splitLine := strings.Split(line, " ")
		loStr := strings.Split(splitLine[0], "-")[0]
		lo, _ := strconv.ParseInt(loStr, 10, 64)
		hiStr := strings.Split(splitLine[0], "-")[1]
		hi, _ := strconv.ParseInt(hiStr, 10, 64)
		ltr := splitLine[1][:1]
		passwd := splitLine[2]

		fmt.Printf("LO: %v, HI: %v, LTR: %s, PASSWD: %s\n", lo, hi, ltr, passwd)

		count := strings.Count(passwd, ltr)
		fmt.Printf("COUNT: %d (%v)\n", count, reflect.TypeOf(count))
		cnt := int64(count)
		if cnt >= lo && cnt <= hi {
			valid++
			fmt.Printf("VALID! (%d)\n", valid)
		}
		if cnt < lo || cnt > hi {
			fmt.Println("UNVALID!")
			unvalid++
		}
		lo = lo - 1
		hi = hi - 1
		if !(lo-1 > int64(len(passwd)) == true) || !(hi > int64(len(passwd)) == true) {
			if !(string(passwd[lo]) == ltr && string(passwd[hi]) == ltr) && (string(passwd[lo]) == ltr || string(passwd[hi]) == ltr) {
				fmt.Println("Strict compliant!")
				strictPolicy++
			}
		} else {
			fmt.Println("Not strict compliant!")
		}

	}

	fmt.Println()
	fmt.Printf("VALID: %d\nUNVALID: %d\n", valid, unvalid)
	fmt.Printf("STRICT: %d\n", strictPolicy)

}
