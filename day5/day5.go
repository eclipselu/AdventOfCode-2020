package main

import (
	"bufio"
	"fmt"
	"os"
)

type boardingPass struct {
	Row int
	Col int
}

func fromPassStr(passStr string) boardingPass {
	return boardingPass{
		Row: getRow(passStr),
		Col: getCol(passStr),
	}
}

func (s *boardingPass) ID() int {
	return s.Row*8 + s.Col
}

func getRow(pass string) int {
	row := 0
	for i := 0; i < 7; i++ {
		row = row * 2
		if pass[i] == 'B' {
			row++
		}
	}
	return row
}

func getCol(pass string) int {
	col := 0
	for i := 7; i < 10; i++ {
		col = col * 2
		if pass[i] == 'R' {
			col++
		}
	}
	return col
}

func readInput(fn string) []boardingPass {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var passes []boardingPass
	for reader.Scan() {
		passes = append(passes, fromPassStr(reader.Text()))
	}

	return passes
}

func maxID(passes []boardingPass) int {
	answer := 0
	for _, pass := range passes {
		id := pass.ID()
		if id > answer {
			answer = id
		}
	}
	return answer
}

func missingSeat(passes []boardingPass) boardingPass {
	countsByRow := make(map[int]int)
	for _, pass := range passes {
		countsByRow[pass.Row]++
	}

	row := -1
	for i := 1; i < 127; i++ {
		if countsByRow[i] == 7 {
			row = i
			break
		}
	}

	if row == -1 {
		panic("cannot find missing row")
	}

	// sum of 0..7
	col := (0 + 7) * 8 / 2
	for _, pass := range passes {
		if pass.Row == row {
			col = col - pass.Col
		}
	}

	return boardingPass{
		Row: row,
		Col: col,
	}
}

func main() {
	passes := readInput("day5.in")
	answer := missingSeat(passes)
	fmt.Println(answer, answer.ID())
}
