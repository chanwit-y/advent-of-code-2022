package day

import (
	"advent-of-code-2022/pkg"
	"strconv"
)

func Day1_1() (res int) {
	inputs := pkg.InputList("input-1")

	sum := 0
	sums := []int{}
	for _, v := range inputs {
		if v != "" {
			n, _ := strconv.Atoi(v)
			sum += n
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}

	return pkg.Max(sums)
}
