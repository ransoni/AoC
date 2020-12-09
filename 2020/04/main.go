package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/go-playground/validator.v9"
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
	text []string
	str  string
}

// Passport is...
type Passport struct {
	//(Birth Year)
	Byr string `validate:"required"`
	// (Issue Year)
	Iyr string `validate:"required"`
	// (Expiration Year)
	Eyr string `validate:"required"`
	// (Height)
	Hgt string `validate:"required"`
	// (Hair Color)
	Hcl string `validate:"required"`
	// (Eye Color)
	Ecl string `validate:"required"`
	// (Passport ID)
	Pid string `validate:"required"`
	// (Country ID)
	Cid string //`validate:"required"`
}

type passports struct {
	passport Passport
	valid    bool
}

// AllPassports is ...
type AllPassports struct {
	Passport []passports
}

var (
	allPassports = new(AllPassports)
	// validate     *validator.Validate
)

const (
	debug = false
)

var validate *validator.Validate

func main() {
	d := new(data)
	d.str = readFileToString(fname)
	validate = validator.New()

	// fmt.Printf("Lines: %d\n", len(reSplit(d.str, "^\\s*$")))
	fmt.Printf("Lines: %d\n", len(reSplit(d.str, "\n\n")))

	// Split by empty lines
	for _, passportsRaw := range reSplit(d.str, "\n\n") {
		// Now we have raw passport data in "array"
		// fmt.Println(passportRaw)

		// onePassport := new(Passport)
		onePassport := new(passports)
		for _, passportRaw := range reSplit(passportsRaw, "(\\s|\n)") {
			// allPassports.passport = append(allPassports.passport, allPassports.clasifyData(passportRaw))

			// onePassport.clasifyData(passportRaw)
			onePassport.passport.clasifyData(passportRaw)
			// onePassport += passportRaw
		}

		// fmt.Printf("Validate passport: %+v\n", onePassport)

		onePassport.valid = onePassport.passport.validatePassport()

		allPassports.Passport = append(allPassports.Passport, *onePassport)
		// d.text = append(d.text, onePassport)

		// if i == 20 {
		// 	if debug {
		// 		fmt.Printf("Passport: %+v: %v\n", allPassports)
		// 	}
		// 	fmt.Printf("Num of passports: %d\n", len(allPassports.Passport))

		// 	fmt.Printf("Valid passports: %d\n", allPassports.countValid())
		// 	// testValidate()
		// 	os.Exit(0)
		// }
	}

	fmt.Printf("Num of passports: %d\n", len(allPassports.Passport))

	fmt.Printf("Valid passports: %d\n", allPassports.countValid())

	// d.text = text

	// d.cleanData()

}

func (p *AllPassports) countValid() int {
	count := 0
	for _, passport := range p.Passport {
		if passport.valid {
			count++
		}
	}
	return count
}

func (passport *Passport) validatePassport() bool {
	if debug {
		fmt.Printf("Validate passport: %+v\n", passport)
	}

	err := validate.Struct(passport)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return false
		}

		for _, err := range err.(validator.ValidationErrors) {
			if debug {
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			}
		}

		// from here you can create your own error messages in whatever language you wish
		return false
	}

	if debug {
		fmt.Printf("Passport VALID!\n")
	}
	return true
}

// func (p *passport) clasifyData(pass string) (passport passport) {
func (passport *Passport) clasifyData(pass string) {
	if debug {
		fmt.Printf("clasify, pass: %s\n", pass)
	}
	dataField := strings.Split(pass, ":")

	switch dataField[0] {
	case "byr":
		if debug {
			fmt.Printf("%s: %s\n", dataField[0], dataField[1])
		}
		passport.Byr = dataField[1]
	case "iyr":
		if debug {
			fmt.Printf("%s: %s\n", dataField[0], dataField[1])
		}
		passport.Iyr = dataField[1]
	case "eyr":
		if debug {
			fmt.Printf("%s: %s\n", dataField[0], dataField[1])
		}
		passport.Eyr = dataField[1]
	case "hgt":
		if debug {
			fmt.Printf("%s: %s\n", dataField[0], dataField[1])
		}
		passport.Hgt = dataField[1]
	case "hcl":
		if debug {
			fmt.Printf("%s: %s\n", dataField[0], dataField[1])
		}
		passport.Hcl = dataField[1]
	case "ecl":
		if debug {
			fmt.Printf("%s: %s\n", dataField[0], dataField[1])
		}
		passport.Ecl = dataField[1]
	case "pid":
		if debug {
			fmt.Printf("%s: %s\n", dataField[0], dataField[1])
		}
		passport.Pid = dataField[1]
	case "cid":
		if debug {
			fmt.Printf("%s: %s", dataField[0], dataField[1])
		}
		passport.Cid = dataField[1]
	}
	return
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

func testValidate() { // TEST VALIDATION
	testPassport := &Passport{
		Byr: "1945",
		Iyr: "2018",
		Eyr: "2020",
		Hgt: "165cm",
		// Hcl: "#18171d",
		Ecl: "blu",
		// Pid: "322103252",
		Cid: "240",
	}

	fmt.Printf("%+v", testPassport)

	err := validate.Struct(testPassport)
	if err != nil {
		fmt.Printf("Passport NOT VALID!\n")
		fmt.Println(err)
	}
}
