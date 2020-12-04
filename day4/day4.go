package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passport struct {
	BirthYear  string
	IssueYear  string
	ExpireYear string
	Height     string
	HairColor  string
	EyeColor   string
	PID        string
	CID        string
}

func (p *passport) Valid() bool {
	if p.BirthYear != "" && p.IssueYear != "" && p.ExpireYear != "" && p.Height != "" &&
		p.HairColor != "" && p.EyeColor != "" && p.PID != "" {
		return true
	}
	return false
}

func (p *passport) Valid2() bool {
	return validYear(p.BirthYear, 1920, 2002) && validYear(p.IssueYear, 2010, 2020) &&
		validYear(p.ExpireYear, 2020, 2030) && validHeight(p.Height) &&
		validHairColor(p.HairColor) && validEyeColor(p.EyeColor) && validPID(p.PID)
}

func validYear(s string, min int, max int) bool {
	if len(s) != 4 {
		return false
	}

	year, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return year >= min && year <= max
}

func validHeight(s string) bool {
	isCM := strings.HasSuffix(s, "cm")
	isIN := strings.HasSuffix(s, "in")
	if !isCM && !isIN {
		return false
	}

	number, err := strconv.Atoi(s[:len(s)-2])
	if err != nil {
		return false
	}

	if isCM {
		return number >= 150 && number <= 193
	} else {
		return number >= 59 && number <= 76
	}
}

func validHairColor(s string) bool {
	if len(s) != 7 || s[0] != '#' {
		return false
	}

	for i := 1; i < 7; i++ {
		if !(s[i] >= '0' && s[i] <= '9') && !(s[i] >= 'a' && s[i] <= 'f') {
			return false
		}
	}

	return true
}

func validEyeColor(s string) bool {
	colors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	return colors[s]
}

func validPID(s string) bool {
	if len(s) != 9 {
		return false
	}

	_, err := strconv.Atoi(s)
	return err == nil
}

func readInput(fn string) []passport {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data)-1; i++ {
			if data[i] == '\n' && data[i+1] == '\n' {
				return i + 2, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	reader.Split(split)

	var passports []passport
	for reader.Scan() {
		line := reader.Text()
		elem := parseEntry(line)
		// fmt.Printf("%+v\n", elem)
		passports = append(passports, elem)
	}
	return passports
}

func parseEntry(s string) passport {
	fields := strings.Fields(s)
	m := make(map[string]string)

	for _, field := range fields {
		kvpair := strings.Split(field, ":")
		m[kvpair[0]] = kvpair[1]
	}
	return passport{
		BirthYear:  m["byr"],
		IssueYear:  m["iyr"],
		ExpireYear: m["eyr"],
		Height:     m["hgt"],
		HairColor:  m["hcl"],
		EyeColor:   m["ecl"],
		PID:        m["pid"],
		CID:        m["cid"],
	}
}

func main() {
	passports := readInput("day4.in")
	answer := 0
	for _, passport := range passports {
		if passport.Valid2() {
			answer++
		}
	}
	fmt.Println(answer)
}
