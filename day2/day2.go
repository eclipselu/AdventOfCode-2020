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
	Min int
	Max int
	Letter rune
	Phrase string
}

func (p *password) IsValid() bool {
	cnt := 0
	for _, ch := range p.Phrase {
		if ch == p.Letter {
			cnt++
		}
	}
	return cnt >= p.Min && cnt <= p.Max
}

func readInput(fn string) []password {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()

	var nums []int
	reader := bufio.NewScanner(file)
	for reader.Scan() {

		num, _ := strconv.Atoi(reader.Text())
		nums = append(nums, num)
	}

	return nums
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}

func parseLine(line string) (password, error) {
	splits := strings.FieldsFunc(line, func(r rune) bool {
		return r == '-' || r == ':' || unicode.IsSpace(r)
	})
	if len(splits) != 4 {
		return password{}, fmt.Errorf("invalid line: %v", line)
	}

	min, max, letter, phrase := splits[0], splits[1], splits[2], splits[3]

}

func main() {
	pwd := &password{
		Min: 1,
		Max: 2,
		Letter: 'c',
		Phrase: "abccccd",
	}
	fmt.Println(pwd, pwd.IsValid())
}
