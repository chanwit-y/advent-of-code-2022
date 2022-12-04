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

func indicated(y int) int {
	switch y {
	case 1:
		return 0
	case 2:
		return 3
	case 3:
		return 6
	default:
		return 0
	}
}

func needToEnd(o, i int) (res int) {
	switch i {
	case 1:
		return toLose(o)
	case 2:
		return o
	case 3:
		return toWin(o)
	default:
		return res
	}
}

func toLose(o int) (res int) {
	switch o {
	case 1:
		return 3
	case 2:
		return 1
	case 3:
		return 2
	default:
		return res
	}

}

func toWin(o int) (res int) {
	switch o {
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 1
	default:
		return res
	}

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

func Day2_2() (res int) {
	inputList := pkg.InputList("input-2-2")
	for _, v := range inputList {
		inputs := strings.Split(v, " ")
		opponent := convStrToEnumVal(inputs[0])
		ind := convStrToEnumVal(inputs[1])

		// fmt.Printf("%v \n", needToEnd(int(opponent), int(ind)))
		// fmt.Printf("%v \n", indicated(int(ind)))
		res += needToEnd(int(opponent), int(ind)) + indicated(int(ind))
	}

	return res
}
