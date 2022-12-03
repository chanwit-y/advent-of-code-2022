package pkg

import (
	"fmt"
	"os"
	"strings"
)

func InputList(name string) (inputs []string) {
	dat, _ := os.ReadFile(fmt.Sprintf("data/%s.txt", name))
	return strings.Split(string(dat), "\n")
}

func Max(v []int) (m int) {
	m = 0
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}

	return m
}

func MaxIndex(v []int) (index int) {
	m := 0
	for i, e := range v {
		if i == 0 || e > m {
			m = e
			index = i
		}
	}

	return index
}
