package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const preambleLen = 25

func isValid(pos int, nums []int) bool {
	if pos < preambleLen {
		return true
	}

	visited := make(map[int]bool)
	num := nums[pos]
	for i := pos - preambleLen; i < pos; i++ {
		remain := num - nums[i]
		if visited[remain] {
			return true
		}

		visited[nums[i]] = true
	}

	return false
}

func solve1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if !isValid(i, nums) {
			return nums[i]
		}
	}

	panic("cannot find such number")
}

func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return min, max
}

func solve2(nums []int) int {
	invalidNum := solve1(nums)
	preSum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 2; j <= len(nums); j++ {
			if preSum[j]-preSum[i] != invalidNum {
				continue
			}
			// found one
			// fmt.Println(nums[i:j])
			min, max := minMax(nums[i:j])
			return min + max
		}
	}

	panic("cannot find such continguous set")
}

func readInput(fn string) []int {
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

func main() {
	nums := readInput("day9.in")
	answer := solve2(nums)
	fmt.Println(answer)
}
