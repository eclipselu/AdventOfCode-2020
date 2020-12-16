package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	Opcode  string
	Operand int
}

func (ins *instruction) Flip() instruction {
	if ins.Opcode == "nop" {
		return instruction{"jmp", ins.Operand}
	} else if ins.Opcode == "jmp" {
		return instruction{"nop", ins.Operand}
	}

	return *ins
}

func solve1(instructions []instruction) (int, bool) {
	visited := make(map[int]bool)
	pc := 0
	acc := 0
	terminates := true

	for pc < len(instructions) {
		if visited[pc] {
			terminates = false
			break
		}

		instr := instructions[pc]
		visited[pc] = true

		switch instr.Opcode {
		case "nop":
			pc++
		case "jmp":
			pc += instr.Operand
		case "acc":
			pc++
			acc += instr.Operand
		default:
			panic(fmt.Errorf("invalid instruction: %v", instr))
		}
	}

	return acc, terminates
}

func solve2(instructions []instruction) int {
	for i := 0; i < len(instructions); i++ {
		if instructions[i].Opcode == "acc" {
			continue
		}

		instructions[i] = instructions[i].Flip()
		if val, terminates := solve1(instructions); terminates {
			instructions[i] = instructions[i].Flip()
			return val
		}
		instructions[i] = instructions[i].Flip()
	}

	panic("cannot find the sole corrupted instruction")
}

func readInput(fn string) []instruction {
	file, err := os.Open(fn)
	if err != nil {
		panic("no such file")
	}
	defer file.Close()

	var instructions []instruction
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		fields := strings.Fields(line)
		opcode := fields[0]
		operand, _ := strconv.Atoi(fields[1])

		instructions = append(instructions, instruction{Opcode: opcode, Operand: operand})
	}
	return instructions
}

func main() {
	instructions := readInput("day8.in")
	// answer, terminates := solve1(instructions)
	answer := solve2(instructions)
	fmt.Println(answer)
}
