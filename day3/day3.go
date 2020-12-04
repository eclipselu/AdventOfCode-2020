package main

import (
	"bufio"
	"fmt"
	"os"
)

const tree = '#'

type board struct {
	Map    []string
	Width  int
	Height int
}

func (b *board) Get(x, y int) (byte, error) {
	if x < 0 || x >= b.Height || y < 0 {
		return byte(' '), fmt.Errorf("out of bound: x=%v, y=%v", x, y)
	}

	return b.Map[x][y%b.Width], nil
}

func (b *board) Solve(dx, dy int) int {
	answer := 0
	for x, y := 0, 0; x < b.Height; x, y = x+dx, y+dy {
		ch, _ := b.Get(x, y)
		if ch == tree {
			answer++
		}
	}
	return answer
}

func readInput(fn string) board {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var bd []string
	for reader.Scan() {
		bd = append(bd, reader.Text())
	}

	if len(bd) == 0 {
		panic("empty map")
	}

	return board{
		Map:    bd,
		Width:  len(bd[0]),
		Height: len(bd),
	}
}

func main() {
	b := readInput("day3.in")
	// part 1
	// slopes := [][]int{{1, 3}}
	// part 2
	slopes := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	answer := 1
	for _, slope := range slopes {
		dx, dy := slope[0], slope[1]
		answer = answer * b.Solve(dx, dy)
	}
	fmt.Println(answer)
}
