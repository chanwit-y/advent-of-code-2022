package day

import (
	"advent-of-code-2022/pkg"
	"strconv"
	"strings"
)

// var stacks = [][]string{{"Z", "N"}, {"M", "C", "D"}, {"P"}}

var stacks1 = [][]string{
	{"P", "F", "M", "Q", "W", "G", "R", "T"},
	{"H", "F", "R"},
	{"P", "Z", "R", "V", "G", "H", "S", "D"},
	{"Q", "H", "P", "B", "F", "W", "G"},
	{"P", "S", "M", "J", "H"},
	{"M", "Z", "T", "H", "S", "R", "P", "L"},
	{"P", "T", "H", "N", "M", "L"},
	{"F", "D", "Q", "R"},
	{"D", "S", "C", "N", "L", "P", "H"},
}

func getInput(v string) (move, from, to int) {
	inputs := strings.Split(v, " ")
	move, _ = strconv.Atoi(inputs[1])
	from, _ = strconv.Atoi(inputs[3])
	to, _ = strconv.Atoi(inputs[5])

	from--
	to--

	return move, from, to
}

func Day5_1() (res string) {
	// fmt.Printf("%+v \n", stacks)
	inputList := pkg.InputList("input-5-1")
	for _, v := range inputList {
		move, from, to := getInput(v)
		// fmt.Printf("move: %v, from: %v, to: %v \n", move, from, to)

		for i := 0; i < move; i++ {
			lastIndex := len(stacks1[from]) - 1
			// fmt.Printf("lastIndex: %v \n", lastIndex)
			crate := stacks1[from][lastIndex]
			stacks1[from] = stacks1[from][:lastIndex]
			stacks1[to] = append(stacks1[to], crate)
			// fmt.Printf("crate: %v ,pop: %v \n", crate, stacks)
		}
	}
	for _, v := range stacks1 {
		res += v[len(v)-1]
	}
	return res
}

func Day5_2() (res string) {
	inputList := pkg.InputList("input-5-2")

	for _, v := range inputList {
		move, from, to := getInput(v)
		// fmt.Printf("input: %v \n", v)
		// for i := 0; i < move; i++ {
		// lastIndex := len(stacks2[from]) - 1
		startIndex := len(stacks1[from]) - move
		// fmt.Printf("stacks2[from]: %v, lastIndex: %v, move: %v, startIndex: %v, to: %v \n", stacks2[from], lastIndex, move, startIndex, to)
		crates := stacks1[from][startIndex:]
		// fmt.Printf("crates: %+v \n", crates)
		stacks1[from] = stacks1[from][:startIndex]
		stacks1[to] = append(stacks1[to], crates...)

		// fmt.Printf("stacks2: %v \n", stacks2)
		// }
	}

	for _, v := range stacks1 {
		res += v[len(v)-1]
	}
	return res
}
