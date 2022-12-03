package day

import (
	"advent-of-code-2022/pkg"
	"strconv"
)

func sum(inputs []string) (sums []int) {
	sum := 0
	sums = []int{}
	for _, v := range inputs {
		if v != "" {
			n, _ := strconv.Atoi(v)
			sum += n
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}

	return sums
}

func Day1_1() (res int) {
	res, _ = pkg.Max(sum(pkg.InputList("input-1")))
	return res
}

func Day1_2() (res int) {
	sums := sum(pkg.InputList("input-1-1"))

	top3 := make([]int, 3)
	for i := 0; i < 3; i++ {
		res, ix := pkg.Max(sums)
		top3[i] = res
		sums[ix] = 0
	}

	return pkg.Sum(top3)
}
