package day

import (
	"advent-of-code-2022/pkg"
)

func Day3_1() (res int) {
	inputList := pkg.InputList("input-3-1")
	for _, v := range inputList {
		half := len(v) / 2
		comp1 := v[0:half]
		comp2 := v[half:]

		// fmt.Printf("c: %v \n", v)
		// fmt.Printf("c1: %v \n", comp1)
		// fmt.Printf("c2: %v \n", comp2)

		val := 0
		for _, c1 := range comp1 {
			for _, c2 := range comp2 {
				if c1 == c2 {
					// fmt.Printf("%v: %v \n", v, string(c1))
					if c1 >= 97 && c1 <= 122 {
						val = int(c1) - 96
						// fmt.Printf("x: %v-> %v \n", c1, val)
					} else if c1 >= 65 && c1 <= 90 {
						val = int(c1) - 38
						// fmt.Printf("xx: %v-> %v \n", c1, val)
					}
				}
			}
		}
		// fmt.Printf("val: %v \n", val)
		res += val
	}
	return res
}
