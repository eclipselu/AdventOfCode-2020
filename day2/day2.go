package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type password struct {
	Min    int
	Max    int
	Letter byte
	Phrase string
}

func (p *password) IsValid() bool {
	cnt := 0
	for i := 0; i < len(p.Phrase); i++ {
		ch := p.Phrase[i]
		if ch == p.Letter {
			cnt++
		}
	}
	return cnt >= p.Min && cnt <= p.Max
}

func (p *password) IsValid2() bool {
	cnt := 0
	if p.Min >= 1 && p.Min <= len(p.Phrase) && p.Phrase[p.Min-1] == p.Letter {
		cnt++
	}
	if p.Max >= 1 && p.Max <= len(p.Phrase) && p.Phrase[p.Max-1] == p.Letter {
		cnt++
	}
	return cnt == 1
}

func readInput(fn string) []password {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()

	var pwlist []password
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		pwd, err := parseLine(reader.Text())
		if err != nil {
			panic(err)
		}
		pwlist = append(pwlist, pwd)
	}

	return pwlist
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func parseLine(line string) (password, error) {
	splits := strings.FieldsFunc(line, func(r rune) bool {
		return r == '-' || r == ':' || unicode.IsSpace(r)
	})
	if len(splits) != 4 {
		return password{}, fmt.Errorf("splits should be 4, invalid line: %v", line)
	}

	if !isInt(splits[0]) || !isInt(splits[1]) || len(splits[2]) != 1 {
		return password{}, fmt.Errorf("invalid format: %v", splits)
	}

	min, _ := strconv.Atoi(splits[0])
	max, _ := strconv.Atoi(splits[1])
	letter, phrase := splits[2][0], splits[3]

	return password{
		Min:    min,
		Max:    max,
		Letter: letter,
		Phrase: phrase,
	}, nil
}

func main() {
	pwlist := readInput("day2.in")
	count := 0
	for _, pwd := range pwlist {
		// if pwd.IsValid() {
		if pwd.IsValid2() {
			count++
		}
	}
	fmt.Println(count)
}
