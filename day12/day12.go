package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type moveAction struct {
	Action byte
	Value  int
}

func (a moveAction) String() string {
	return fmt.Sprintf("{Action: %c, Value: %d}", a.Action, a.Value)
}

type position struct {
	X   int
	Y   int
	Deg int // degree between direction and east direction
}

func (p position) Next(a moveAction) position {
	switch a.Action {
	case 'N':
		p.Y += a.Value
	case 'S':
		p.Y -= a.Value
	case 'E':
		p.X += a.Value
	case 'W':
		p.X -= a.Value
	case 'L':
		p.Deg = (p.Deg + a.Value + 360) % 360
	case 'R':
		p.Deg = (p.Deg - a.Value + 360) % 360
	case 'F':
		dx := a.Value * int(math.Round(math.Cos(float64(p.Deg)/180.0*math.Pi)))
		dy := a.Value * int(math.Round(math.Sin(float64(p.Deg)/180.0*math.Pi)))
		p.X += dx
		p.Y += dy
	}

	return p
}

type position2 struct {
	X         int
	Y         int
	WaypointX int
	WaypointY int
}

func (p position2) Len() float64 {
	return math.Sqrt(float64(p.WaypointX*p.WaypointX + p.WaypointY*p.WaypointY))
}

func (p position2) Rotate(degree int) position2 {
	angle := math.Atan2(float64(p.WaypointY), float64(p.WaypointX))
	angle += float64(degree) / 180.0 * math.Pi
	length := p.Len()

	// don't forget to round
	p.WaypointX = int(math.Round(length * math.Cos(angle)))
	p.WaypointY = int(math.Round(length * math.Sin(angle)))
	return p
}

func (p position2) Next(a moveAction) position2 {
	switch a.Action {
	case 'N':
		p.WaypointY += a.Value
	case 'S':
		p.WaypointY -= a.Value
	case 'E':
		p.WaypointX += a.Value
	case 'W':
		p.WaypointX -= a.Value
	case 'L':
		p = p.Rotate(a.Value)
	case 'R':
		p = p.Rotate(-a.Value)
	case 'F':
		p.X += p.WaypointX * a.Value
		p.Y += p.WaypointY * a.Value
	}

	return p
}

func readInput(fn string) []moveAction {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()

	var actions []moveAction
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		value, _ := strconv.Atoi(line[1:])
		actions = append(actions, moveAction{Action: line[0], Value: value})
	}

	return actions
}

func solve1(actions []moveAction) int {
	pos := position{X: 0, Y: 0, Deg: 0}
	for _, action := range actions {
		pos = pos.Next(action)
		// fmt.Println(pos)
	}

	return int(math.Abs(float64(pos.X)) + math.Abs(float64(pos.Y)))
}

func solve2(actions []moveAction) int {
	pos := position2{X: 0, Y: 0, WaypointX: 10, WaypointY: 1}
	for _, action := range actions {
		pos = pos.Next(action)
		// fmt.Printf("%v %+v\n", action, pos)
	}

	return int(math.Abs(float64(pos.X)) + math.Abs(float64(pos.Y)))
}

func main() {
	actions := readInput("day12.in")
	answer := solve2(actions)
	fmt.Println(answer)
}
