package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type declFormGroup struct {
	Forms []string
}

// count of questions that are answered yes
func (d *declFormGroup) YesCount() int {
	count := make(map[byte]bool)
	for _, form := range d.Forms {
		for i := 0; i < len(form); i++ {
			count[form[i]] = true
		}
	}

	return len(count)
}

// count of questions that are answered yes by everyone in the group
func (d *declFormGroup) YesCountEveryone() int {
	count := make(map[byte]int)
	for _, form := range d.Forms {
		for i := 0; i < len(form); i++ {
			count[form[i]]++
		}
	}

	cnt := 0
	for _, v := range count {
		if v == len(d.Forms) {
			cnt++
		}
	}

	return cnt
}

func readInput(fn string) []declFormGroup {
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

	var groups []declFormGroup
	for reader.Scan() {
		line := reader.Text()
		groups = append(groups, declFormGroup{Forms: strings.Split(line, "\n")})
	}
	return groups
}

func solve(groups []declFormGroup) int {
	count := 0
	for _, group := range groups {
		// count += group.YesCount()
		count += group.YesCountEveryone()
	}
	return count
}

func main() {
	groups := readInput("day6.in")
	answer := solve(groups)
	fmt.Println(answer)
}
