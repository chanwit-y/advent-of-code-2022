package day

import (
	"advent-of-code-2022/pkg"
	"strings"
)

type RPS int

const (
	Rock     RPS = 1
	Paper        = 2
	Scissors     = 3
)

func convStrToEnumVal(v string) (res RPS) {
	switch {
	case v == "X" || v == "A":
		res = Rock
	case v == "Y" || v == "B":
		res = Paper
	case v == "Z" || v == "C":
		res = Scissors
	}
	return res
}

// 3 -> 2
// 2 -> 3
// 1 -> 3
func findScore(op, you int) (score int) {
	if op == you {
		return 3
	}
	if (you == 3 && op == 2) ||
		(you == 2 && op == 1) ||
		(you == 1 && op == 3) {
		return 6
	}
	return score
}

func Day2_1() (res int) {
	inputList := pkg.InputList("input-2-1")
	for _, v := range inputList {
		inputs := strings.Split(v, " ")
		opponent := convStrToEnumVal(inputs[0])
		you := convStrToEnumVal(inputs[1])

		res += int(you) + findScore(int(opponent), int(you))
	}

	return res
}
