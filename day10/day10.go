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

	sort.Ints(nums)
	nums = append([]int{0}, nums...)
	nums = append(nums, nums[len(nums)-1]+3)
	return nums
}

func solve1(nums []int) int {
	prev := 0
	count1, count3 := 0, 0

	for _, num := range nums {
		diff := num - prev
		prev = num
		switch diff {
		case 1:
			count1++
		case 3:
			count3++
		}
	}

	fmt.Println(count1, count3)
	return count1 * count3
}

func solve2(nums []int) uint64 {
	memo := make(map[int]uint64)
	return dfs(0, nums, memo)
}

func dfs(start int, nums []int, memo map[int]uint64) uint64 {
	if val, ok := memo[start]; ok {
		return val
	}

	if start >= len(nums)-1 {
		memo[start] = uint64(1)
		return uint64(1)
	}

	count := uint64(0)
	for i := start + 1; i < len(nums) && nums[start]+3 >= nums[i]; i++ {
		count += dfs(i, nums, memo)
	}

	memo[start] = count
	return count
}

func main() {
	nums := readInput("day10.in")
	// nums := readInput("day10.example")
	answer := solve2(nums)
	fmt.Println(answer)
}
