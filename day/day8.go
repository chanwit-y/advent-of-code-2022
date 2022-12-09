package day

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"strconv"
)

func Day8_1() (res int) {

	inputList := pkg.InputList("input-8-1")
	grid := make([][]int, len(inputList))

	for i, v := range inputList {
		grid[i] = make([]int, len(v))

		for j, n := range v {
			num, _ := strconv.Atoi(string(n))
			grid[i][j] = num
		}
	}
	interior := []int{}
	visible := []int{}

	for i, v := range grid {
		for j, n := range v {
			if i == 0 || i == len(grid)-1 {
				interior = append(interior, n)
			} else if j == 0 || j == len(v)-1 {
				interior = append(interior, n)
			} else {

				canSeeTop := true
				for top := 0; top < i; top++ {
					if canSeeTop {
						canSeeTop = grid[top][j] < n
					}
				}

				canSeeLeft := true
				for left := 0; left < j; left++ {
					if canSeeLeft {
						canSeeLeft = grid[i][left] < n
					}
				}

				canSeeRight := true
				for right := j + 1; right < len(grid[i]); right++ {
					if canSeeRight {
						canSeeRight = grid[i][right] < n
					}
				}

				canSeeBottom := true
				for bottom := i + 1; bottom < len(grid); bottom++ {
					if canSeeBottom {
						canSeeBottom = grid[bottom][j] < n
					}
				}

				// if n == 1 {
				// 	fmt.Printf("canSeeTop: %v \n", canSeeTop)
				// 	fmt.Printf("canSeeLeft: %v \n", canSeeLeft)
				// 	fmt.Printf("canSeeRight: %v \n", canSeeRight)
				// 	fmt.Printf(" canSeeBottom: %v \n", canSeeBottom)
				// }

				if canSeeTop || canSeeLeft || canSeeRight || canSeeBottom {
					visible = append(visible, n)
				}

			}
		}
	}

	fmt.Printf("interior: %v \n", len(interior))
	fmt.Printf("visible: %v \n", len(visible))

	return len(interior) + len(visible)
}
