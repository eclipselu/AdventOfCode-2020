package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

func findPairProduct(nums []int, target int, start int, end int) (int, error) {
	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			return nums[start] * nums[end], nil
		} else if sum < target {
			start++
		} else {
			end--
		}
	}

	return -1, fmt.Errorf("no such pair found")
}

func findTripletProduct(nums []int, target int) (int, error) {
	n := len(nums)
	for i, num := range nums {
		prod, err := findPairProduct(nums, target - num, i + 1, n - 1)
		if err == nil {
			return prod * num, nil
		}
	}

	return -1, fmt.Errorf("no such triplet")
}

func main() {
	nums := readInput("day1.in")
	sort.Ints(nums)
	answer1, err := findPairProduct(nums, 2020, 0, len(nums) - 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1: ", answer1)

	answer2, err := findTripletProduct(nums, 2020)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2: ", answer2)
}
