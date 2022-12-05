package day

import (
	"advent-of-code-2022/pkg"
	"fmt"
)

func Day3_1() (res int) {
	inputList := pkg.InputList("input-3-1")
	for _, v := range inputList {
		half := len(v) / 2
		comp1 := v[0:half]
		comp2 := v[half:]

		val := 0
		for _, c1 := range comp1 {
			for _, c2 := range comp2 {
				if c1 == c2 {
					if c1 >= 97 && c1 <= 122 {
						val = int(c1) - 96
					} else if c1 >= 65 && c1 <= 90 {
						val = int(c1) - 38
					}
				}
			}
		}
		res += val
	}
	return res
}

func Day3_2() (res int) {
	inputList := pkg.InputList("input-3-2")
	for i := 0; i < len(inputList); {
		b1s := inputList[i]
		b2s := inputList[i+1]
		b3s := inputList[i+2]

		val := 0
		for _, b1 := range b1s {
			valB1and2 := []rune{}
			for _, b2 := range b2s {
				if b1 == b2 {
					valB1and2 = append(valB1and2, b2)
				}
			}

			for _, b3 := range b3s {
				for _, vb1b2 := range valB1and2 {
					if b3 == vb1b2 {
						if b3 >= 97 && b3 <= 122 {
							val = int(b3) - 96
							fmt.Printf("%v:%v \n", string(b3), val)
						} else if b3 >= 65 && b3 <= 90 {
							val = int(b3) - 38
							fmt.Printf("%v:%v \n", string(b3), val)
						}
					}
				}
			}
		}
		res += val

		i += 3
	}

	return res
}
