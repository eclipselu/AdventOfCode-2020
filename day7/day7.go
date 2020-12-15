package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagRule struct {
	Color string
	Bags  map[string]int // color name -> count
}

type bagCount struct {
	Color string
	Count int
}

func readInput(fn string) []bagRule {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()

	var rules []bagRule
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		rules = append(rules, parseRule(line))
	}
	return rules
}

func parseRule(ruleText string) bagRule {
	ruleText = strings.TrimSuffix(ruleText, ".")
	splits := strings.Split(ruleText, "contain")
	if len(splits) != 2 {
		panic(fmt.Errorf("invalid input: %v", ruleText))
	}

	color, bgsInfo := strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])
	color = strings.TrimSuffix(color, " bags")

	var bags map[string]int
	if bgsInfo != "no other bags" {
		bgs := strings.Split(bgsInfo, ", ")
		bags = make(map[string]int)
		for _, bg := range bgs {
			fields := strings.Fields(bg)
			count, _ := strconv.Atoi(fields[0])
			name := strings.Join(fields[1:len(fields)-1], " ")
			bags[name] = count
		}
	}

	return bagRule{
		Color: color,
		Bags:  bags,
	}
}

func solve1(rules []bagRule) int {
	graph := make(map[string][]string) // if A contains B, then B -> A
	for _, rule := range rules {
		for bagColor := range rule.Bags {
			graph[bagColor] = append(graph[bagColor], rule.Color)
		}
	}

	visited := make(map[string]bool)
	start := "shiny gold"

	queue := list.New()
	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		cur := queue.Front()
		for _, next := range graph[cur.Value.(string)] {
			if !visited[next] {
				visited[next] = true
				queue.PushBack(next)
			}
		}
		queue.Remove(cur)
	}

	return len(visited) - 1
}

func solve2(rules []bagRule) int {
	graph := make(map[string][]bagCount) // if A contains B, then B -> A
	for _, rule := range rules {
		for bagColor, count := range rule.Bags {
			graph[rule.Color] = append(graph[rule.Color], bagCount{Color: bagColor, Count: count})
		}
	}

	return dfs("shiny gold", graph) - 1
}

func dfs(start string, graph map[string][]bagCount) int {
	if graph[start] == nil {
		return 1
	}

	count := 1
	for _, bc := range graph[start] {
		cnt := dfs(bc.Color, graph)
		count += cnt * bc.Count
	}

	// fmt.Printf("color: %s, count: %d\n", start, count)
	return count
}

func main() {
	rules := readInput("day7.in")
	fmt.Println(solve2(rules))
}
