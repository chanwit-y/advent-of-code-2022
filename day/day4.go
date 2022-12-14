package day

import (
	"advent-of-code-2022/pkg"
	"strconv"
	"strings"
)

func getNumber(v string) (a1, a2, b1, b2 int) {
	section := strings.Split(v, ",")
	s1 := strings.Split(section[0], "-")
	s2 := strings.Split(section[1], "-")

	a1, _ = strconv.Atoi(s1[0])
	a2, _ = strconv.Atoi(s1[1])

	b1, _ = strconv.Atoi(s2[0])
	b2, _ = strconv.Atoi(s2[1])

	return a1, a2, b1, b2
}

// a in b
//  b1   a1   b2   a2
// (4 <= 6 && 6 >= 6) = T
// ||
// b in a
//  a1   b1   a2   b2
// (2 <= 3 && 8 >= 7)

// 2 <= 6 && 4 >= 8 = F
// ||
// 6 <= 2 && 8 >= 4 = F
func Day4_1() (res int) {
	inputList := pkg.InputList("input-4-1")
	for _, v := range inputList {
		a1, a2, b1, b2 := getNumber(v)
		if (b1 <= a1 && b2 >= a2) || (a1 <= b1 && a2 >= b2) {
			res++
		}
	}

	return res
}

func Day4_2() (res int) {
	inputList := pkg.InputList("input-4-2")
	for _, v := range inputList {
		a1, a2, b1, b2 := getNumber(v)
		if ((a1 >= b1 && a1 <= b2) || (a2 >= b1 && a2 <= b2)) ||
			((b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2)) {
			res++
		}
	}

	return res
}
