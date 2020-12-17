package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func readInput(fn string) [][]byte {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()

	var board [][]byte
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		board = append(board, []byte(reader.Text()))
	}

	return board
}

func adjacentCount1(row, col int, board [][]byte) int {
	width, height := len(board[0]), len(board)
	cnt := 0
	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]
		if r < 0 || r >= height || c < 0 || c >= width {
			continue
		}
		if board[r][c] == '#' {
			cnt++
		}
	}
	return cnt
}

func adjacentCount2(row, col int, board [][]byte) int {
	width, height := len(board[0]), len(board)
	cnt := 0
	for _, dir := range directions {
		for r, c := row+dir[0], col+dir[1]; r >= 0 && r < height && c >= 0 && c < width; r, c = r+dir[0], c+dir[1] {
			if board[r][c] == '#' {
				cnt++
				break
			} else if board[r][c] == 'L' {
				break
			}
		}
	}
	return cnt
}

func iterate(board [][]byte, thresh int, adjacentCountFunc func(row, col int, bd [][]byte) int) ([][]byte, bool) {
	old := clone(board)
	width, height := len(board[0]), len(board)

	changed := false

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if old[r][c] == '.' {
				continue
			}
			cnt := adjacentCountFunc(r, c, old)
			if old[r][c] == 'L' && cnt == 0 {
				board[r][c] = '#'
				changed = true
			}
			if old[r][c] == '#' && cnt >= thresh {
				board[r][c] = 'L'
				changed = true
			}
		}
	}
	return board, changed
}

func print(board [][]byte) {
	for _, line := range board {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func clone(board [][]byte) [][]byte {
	newBoard := make([][]byte, len(board))
	for i := 0; i < len(board); i++ {
		newBoard[i] = make([]byte, len(board[i]))
		copy(newBoard[i], board[i])
	}

	return newBoard
}

func occupiedCount(board [][]byte) int {
	cnt := 0
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == '#' {
				cnt++
			}
		}
	}

	return cnt
}

func solve1(board [][]byte) int {
	var bd [][]byte
	var changed bool

	for {
		bd, changed = iterate(board, 4, adjacentCount1)
		if !changed {
			break
		}
	}

	return occupiedCount(bd)
}

func solve2(board [][]byte) int {
	var bd [][]byte
	var changed bool

	for {
		bd, changed = iterate(board, 5, adjacentCount2)
		if !changed {
			break
		}
	}

	return occupiedCount(bd)
}

func main() {
	board := readInput("day11.in")
	answer := solve2(board)
	fmt.Println(answer)
}
