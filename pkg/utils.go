package pkg

import (
	"fmt"
	"os"
	"strings"
)

func ReadFileToStr(name string) string {
	dat, _ := os.ReadFile(fmt.Sprintf("data/%s.txt", name))
	return string(dat)
}

func InputList(name string) (inputs []string) {
	dat, _ := os.ReadFile(fmt.Sprintf("data/%s.txt", name))
	return strings.Split(string(dat), "\n")
}

func Max(v []int) (max, index int) {
	max = 0
	index = -1
	for i, e := range v {
		if i == 0 || e > max {
			max = e
			index = i
		}
	}

	return max, index
}

func Sum(vs []int) (sum int) {
	sum = 0
	for _, v := range vs {
		sum += v
	}

	return sum
}

// func MaxIndex(v []int) (index int) {
// 	m := 0
// 	for i, e := range v {
// 		if i == 0 || e > m {
// 			m = e
// 			index = i
// 		}
// 	}

// 	return index
// }
