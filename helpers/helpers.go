package helpers

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// ReSplit helps to split string with Regex
// Inputs:
// text:string String to split
// pattern:string Regex pattern to use for splitting
// Return:
// out:array Array of substrings
func ReSplit(text string, pattern string) (out []string) {
	re := regexp.MustCompile(pattern)
	out = re.Split(text, -1)

	return
}

// ReadFileToString reads a file from path and returns string with files content
// input:string File name to read
// Return: string
func ReadFileToString(fname string) string {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

// ReadFileAsArray reads a file from given path, and returns a string array by splitting each line
// fname:string Name of the file
// Return: []string
func ReadFileAsArray(fname string) []string {
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
