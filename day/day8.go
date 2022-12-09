package day

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"strconv"

	"github.com/samber/lo"
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

				if canSeeTop || canSeeLeft || canSeeRight || canSeeBottom {
					visible = append(visible, n)
				}
			}
		}
	}

	return len(interior) + len(visible)
}

func Day8_2() (res int) {

	inputList := pkg.InputList("input-8-1")
	grid := make([][]int, len(inputList))

	for i, v := range inputList {
		grid[i] = make([]int, len(v))

		for j, n := range v {
			num, _ := strconv.Atoi(string(n))
			grid[i][j] = num
		}
	}
	// interior := []int{}
	// visible := []int{}

	result := []int{}

	for i, v := range grid {
		for j, n := range v {
			if i == 0 || i == len(grid)-1 {
				// interior = append(interior, n)
			} else if j == 0 || j == len(v)-1 {
				// interior = append(interior, n)
			} else {

				upScenic := 0
				isUpBlocked := false
				for up := i - 1; up >= 0; up-- {
					if grid[up][j] < n && !isUpBlocked {
						upScenic++
					} else if !isUpBlocked {
						upScenic++
						isUpBlocked = true
					}
				}

				leftScenic := 0
				isLeftBlocked := false

				// if i == 1 && j == 2 {
				// 	fmt.Println(j - 1)
				// }
				for left := j - 1; left >= 0; left-- {

					if grid[i][left] < n && !isLeftBlocked {
						leftScenic++
					} else if !isLeftBlocked {

						// if i == 1 && j == 2 {
						// 	fmt.Println(grid[i][left])
						// }
						leftScenic++
						isLeftBlocked = true
					}
				}

				rightScenic := 0
				isRightBlocked := false
				for right := j + 1; right < len(grid[i]); right++ {
					if grid[i][right] < n && !isRightBlocked {
						rightScenic++
					} else if !isRightBlocked {
						rightScenic++
						isRightBlocked = true
					}
				}

				bottomScenic := 0
				isBottomBlocked := false
				for bottom := i + 1; bottom < len(grid); bottom++ {
					if grid[bottom][j] < n && !isBottomBlocked {
						bottomScenic++
					} else if !isBottomBlocked {
						bottomScenic++
						isBottomBlocked = true
					}
				}

				if i == 3 && j == 2 {
					fmt.Printf("%v*%v*%v*%v = %v \n", upScenic, leftScenic, rightScenic, bottomScenic, upScenic*leftScenic*rightScenic*bottomScenic)
				}

				result = append(result, upScenic*leftScenic*rightScenic*bottomScenic)
			}
		}
	}

	// fmt.Printf("%+v \n", result)

	return lo.Max(result)
}
