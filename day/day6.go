package day

import (
	"advent-of-code-2022/pkg"
	"fmt"
)

// const maxChar = 14

func Day6_1(maxChar int) (res int) {
	input := pkg.ReadFileToStr("input-6-1")
	i := 0
	for i < len(input) {
		isDup := false
		if i+maxChar < len(input) {
			text := input[i : i+maxChar]
			fmt.Printf("%v \n", text)
			j := 0
			for j < len(text) {
				k := j + 1
				for k < len(text) {
					if text[j] == text[k] {
						isDup = true
					}
					k++
				}
				j++
			}
		}
		if !isDup {

			if i+maxChar < len(input) {
				return len(input[:i+maxChar])
			}
			// fmt.Printf("none dup: %v \n", input[:i+4])
		}
		i++
	}
	return res
}

func _Day6_1() (res int) {
	// var result []string
	resStr := ""
	text := ""
	input := pkg.ReadFileToStr("input-6-1")
	i := 0
	foundIndex := 0
	for i < len(input) {
		// fmt.Printf("foundIndex: %v \n", foundIndex)
		if foundIndex != 0 {
			i = foundIndex
			foundIndex = 0
		} else if i != 0 {
			// fmt.Printf("res str: %v, text: %v \n", resStr, text[1:])
			return len(resStr + text[1:])
		}
		end4 := i + 4
		// fmt.Printf("i: %v \n", i)
		if end4 <= len(input) {
			fmt.Printf("sub input 4: %v, i: %v \n", input[i:end4], i)
			text = input[i:end4]

			j := 0
			// fmt.Printf("%v \n", text)
			for j < len(text) {
				k := j + 1
				// fmt.Printf("%v \n", len(text)-j)
				for k < len(text) && k < 4 {
					// fmt.Printf("text[j:%v]: %v, text[k:%v]: %v \n", j, string(text[j]), k, string(text[k]))
					if text[j] == text[k] {
						if j == 0 {
							foundIndex = i + 1
						} else {
							foundIndex = i + j
						}
						if i > 0 {
							resStr += input[i+1 : foundIndex+1]
						} else {
							resStr += input[i : foundIndex+1]
						}
						// i = j
						// fmt.Printf("found! i=%v resStr: %v, now text: %v \n", i, resStr, text)
						// 	// foundIndex = k
						// 	// i += k
						// 	// fmt.Printf("found index: %v, k:%v, i:%v  \n", foundIndex, k, i)
					}
					// if i != 0 &&  {
					// 	resStr += text
					// }
					k++
				}
				j++
			}
			// fmt.Println("-----------------------")

		}

		// fmt.Printf("xxx: %v \n", foundIndex)
		// if foundIndex == 0 {
		// 	i++
		// }
		i++
	}

	return res
}
