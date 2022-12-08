package day

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type (
	// Directory struct {
	// 	Name      string
	// 	File      []File
	// 	Directory []Directory
	// }
	File struct {
		Name string
		Size int
	}

	// Terminal struct {
	// 	Directories []Directory
	// }
)

// Command
// $ cd /
// $ ls print out all of the files and directories
// dir [name]
// [number] [name].xxx

const MAX = 5

func Day7_2() int {
	dirPaths := []string{}
	root := make(map[string][]File)

	commands := pkg.InputList("input-7-2")
	for _, c := range commands {
		line := strings.Split(c, " ")

		if len(line) == 3 {
			if line[1] == "cd" && line[2] != ".." && line[2] != "/" {
				dirPaths = append(dirPaths, line[2])
			} else if line[1] == "cd" && line[2] == ".." {
				dirPaths = dirPaths[:len(dirPaths)-1]
			} else {
				dirPaths = append(dirPaths, "/")
				root[line[2]] = []File{}
			}

		} else if len(line) == 2 {
			if line[0] == "dir" {
				root[strings.Join(dirPaths, "")+line[1]] = []File{}
			} else if line[1] != "ls" {
				size, _ := strconv.Atoi(line[0])
				tmep := root[strings.Join(dirPaths, "")]
				root[strings.Join(dirPaths, "")] = append(tmep, File{Name: line[1], Size: size})
			}
		}

	}

	sum := make(map[string]int)
	for k, v := range root {
		// fmt.Printf("[key: %v] \n", k)
		// fmt.Printf("value: %+v] \n", v)

		// sum in current dir
		for _, f := range v {
			sum[k] += f.Size
		}

		// find indirectly
		for k2, _ := range root {
			if k != k2 && strings.Contains(k, k2) {
				for _, f := range root[k] {
					sum[k2] += f.Size
				}
				// fmt.Printf("k2: %v indirectly k: %v \n", k2, k)
			}
		}

	}

	const DiskSpace = 70000000
	const forUpdate = 30000000
	available := DiskSpace - sum["/"]
	want := sum["/"] - 30000000
	fmt.Printf("root: %v \n", sum["/"])
	fmt.Printf("available: %v \n", available)
	fmt.Printf("want: %v \n", want)

	removeList := []int{}

	for k, v := range sum {
		// fmt.Printf("key: %v, file size: %v \n", k, v)

		used := sum["/"] - v
		avb := DiskSpace - used

		// fmt.Printf("useadble: %v  \n", avb)
		if avb >= forUpdate {
			removeList = append(removeList, v)
			// res += v
			fmt.Printf("key: %v, file size: %v, used: %v, available: %v \n", k, v, used, avb)
		}

		// fmt.Printf("unused: %v \n", unused)

		// if k != "/" && v > want {
		// 	fmt.Println("============================================================================")
		// 	fmt.Printf("key: %v, file size: %v, available size: %v \n", k, v, avb)
		// 	fmt.Println("============================================================================")
		// 	res += v
		// }
	}

	// fmt.Println("====================================")
	// for k, v := range sum {
	// 	fmt.Printf("[key: %v] \n", k)
	// 	fmt.Printf("value: %+v] \n", v)

	// 	if v < 100000 {
	// 		res += v
	// 	}
	// }

	return lo.Min(removeList)
}

func Day7_1() (res int) {
	dirPaths := []string{}
	root := make(map[string][]File)

	commands := pkg.InputList("input-7-1")
	for _, c := range commands {
		line := strings.Split(c, " ")

		if len(line) == 3 {
			if line[1] == "cd" && line[2] != ".." && line[2] != "/" {
				dirPaths = append(dirPaths, line[2])
			} else if line[1] == "cd" && line[2] == ".." {
				dirPaths = dirPaths[:len(dirPaths)-1]
			} else {
				dirPaths = append(dirPaths, "/")
				root[line[2]] = []File{}
			}

		} else if len(line) == 2 {
			if line[0] == "dir" {
				root[strings.Join(dirPaths, "")+line[1]] = []File{}
			} else if line[1] != "ls" {
				size, _ := strconv.Atoi(line[0])
				tmep := root[strings.Join(dirPaths, "")]
				root[strings.Join(dirPaths, "")] = append(tmep, File{Name: line[1], Size: size})
			}
		}

	}

	sum := make(map[string]int)
	for k, v := range root {
		// fmt.Printf("[key: %v] \n", k)
		// fmt.Printf("value: %+v] \n", v)

		// sum in current dir
		for _, f := range v {
			sum[k] += f.Size
		}

		// find indirectly
		for k2, _ := range root {
			if k != k2 && strings.Contains(k, k2) {
				for _, f := range root[k] {
					sum[k2] += f.Size
				}
				// fmt.Printf("k2: %v indirectly k: %v \n", k2, k)
			}
		}

	}
	fmt.Println("====================================")
	for k, v := range sum {
		fmt.Printf("[key: %v] \n", k)
		fmt.Printf("value: %+v] \n", v)

		if v < 100000 {
			res += v
		}
	}

	return res
}

// func deepCreateDir(ter Directory, curPath, dirName string) Directory {
// 	if ter.Name == curPath {
// 		fmt.Printf("1 %v \n", ter.Name)
// 		fmt.Printf("ter: %v \n", ter)
// 		fmt.Printf("ter.Directory: %v \n", ter.Directory)
// 		ter.Directory = append(ter.Directory, Directory{Name: dirName})
// 		return ter
// 	}
// 	for _, t := range ter.Directory {
// 		if t.Name == curPath {
// 			fmt.Printf("found! \n")
// 			fmt.Printf("before add t: %v \n", t)
// 			t.Directory = append(t.Directory, Directory{Name: dirName})
// 			fmt.Printf("after add t: %v \n", t)

// 			// for i := 0; i < len(ter.Directory); i++ {
// 			// 	if ter.Directory[i].Name == curPath {

// 			// 	}
// 			// }
// 			// ter.Directory = append(ter.Directory, t.Directory...)
// 			return ter
// 		}
// 		fmt.Printf("xxxxxxxxxxxxxxxxxxxxx %+v \n", ter)
// 		ter = deepCreateDir(t, curPath, dirName)
// 		// fmt.Printf("xxxxxxxxxxxxxxxxxxxxx %+v \n", ter)
// 	}
// 	// for i := 0; i < len(ter.Directory); i++ {
// 	// 	if ter.Name == curPath {
// 	// 		fmt.Printf("found! \n")
// 	// 		fmt.Printf("found!: %v \n", ter.Directory[i].Directory)
// 	// 		// ter.Directory[i] = append(ter.Directory[i], Directory{Name: dirName})
// 	// 		ter.Directory[i].Directory = append(ter.Directory[i].Directory, Directory{Name: dirName})
// 	// 	}
// 	// 	deepCreateDir(&ter.Directory[i], curPath, dirName)
// 	// 	fmt.Printf("===========> %+v \n", ter)
// 	// }

// 	return ter
// }

// func Day7_1() (res int) {
// 	dirPaths := []string{}
// 	// curX := 0
// 	// root := make([]Directory, MAX)
// 	root := Directory{}
// 	// for i := range terminals {
// 	// 	terminals[i] = make([]Directory, MAX)
// 	// }

// 	commands := pkg.InputList("input-7-1")
// 	for _, c := range commands {
// 		line := strings.Split(c, " ")

// 		if len(line) == 3 {
// 			if line[1] == "cd" && line[2] != ".." && line[2] != "/" {
// 				dirPaths = append(dirPaths, line[2])
// 			} else if line[1] == "cd" && line[2] == ".." {
// 				dirPaths = dirPaths[:len(dirPaths)-1]
// 			} else {
// 				dirPaths = append(dirPaths, "/")
// 				root.Name = line[2]
// 			}

// 			fmt.Printf("Path: %v \n", dirPaths)
// 		} else if len(line) == 2 {
// 			if line[0] == "dir" {
// 				currentPath := dirPaths[len(dirPaths)-1]
// 				var setDir func(dir []Directory) []Directory
// 				setDir = func(dir []Directory) []Directory {
// 					for _, d := range dir {
// 						if d.Name == currentPath {
// 							d.Directory = append(d.Directory, Directory{Name: line[1]})
// 							dir = lo.Map(dir, func(item Directory, index int) Directory {
// 								if item.Name == currentPath {
// 									return d
// 								}
// 								return item
// 							})
// 							return dir
// 						}
// 						root.Directory = setDir(d.Directory)
// 					}
// 					return dir
// 				}
// 				if root.Name == currentPath {
// 					root.Directory = append(root.Directory, Directory{Name: line[1]})
// 				} else {
// 					root.Directory = setDir(root.Directory)
// 				}
// 				// fmt.Printf("root: %+v \n", root)
// 			} else if line[1] != "ls" {
// 				currentPath := dirPaths[len(dirPaths)-1]
// 				var setFile func(dir []Directory) []Directory
// 				setFile = func(dir []Directory) []Directory {
// 					for _, d := range dir {
// 						fmt.Printf("current in file %v \n", currentPath)
// 						if d.Name == "e" {
// 							fmt.Printf("data E %+v \n", d)
// 						}
// 						if d.Name == currentPath {
// 							size, _ := strconv.Atoi(line[0])
// 							d.File = append(d.File, File{Name: line[1], Size: size})
// 							dir = lo.Map(dir, func(item Directory, index int) Directory {
// 								if item.Name == currentPath {
// 									return d
// 								}
// 								return item
// 							})
// 							return dir
// 						}
// 						root.Directory = setFile(d.Directory)
// 					}
// 					return dir
// 				}
// 				if root.Name == currentPath {
// 					size, _ := strconv.Atoi(line[0])
// 					root.File = append(root.File, File{Name: line[1], Size: size})
// 				} else {
// 					root.Directory = setFile(root.Directory)
// 				}
// 				// fmt.Printf("root: %+v \n", root)
// 			}
// 		}

// 	}

// 	// for _, t := range terminals {
// 	// 	// for _, d := range t {
// 	// 	fmt.Printf("dir '%v': %+v \n", t.Name, t)
// 	// 	// }
// 	// 	fmt.Printf("------------------------------------- \n")
// 	// }

// 	fmt.Printf("\n root: %+v \n", root)

// 	return res
// }

// func Day7_1() (res int) {
// 	var currenIndex int
// 	commands := pkg.InputList("input-7-1")
// 	terminal := Terminal{}

// 	for _, c := range commands {
// 		fmt.Printf("%v \n", c)
// 		line := strings.Split(c, " ")
// 		//$ cd++
// 		if len(line) == 3 {
// 			if subStr[1] == "cd" && subStr[2] != ".." && subStr[2] != "/" {
// 				// into dir
// 				currenIndex++
// 				fmt.Printf("index: %v \n", currenIndex)
// 			} else if subStr[1] == "cd" && subStr[2] == ".." {
// 				// back dir
// 				currenIndex--
// 				fmt.Printf("index: %v \n", currenIndex)
// 			} else {
// 				currenIndex = 0
// 			}

// 		} else if len(subStr) == 2 {

// 			if subStr[0] == "dir" {
// 				// create dir
// 				if currenIndex == 0 {
// 					terminal.Directories = append(terminal.Directories, Directory{Name: subStr[1]})
// 				} else {
// 					// tempTerminal := Directory{}
// 					for i := 0; i <= currenIndex; i++ {
// 						tempTerminal := &terminal.Directories[i]
// 						if i == currenIndex {
// 							tempTerminal.Directory = append(tempTerminal.Directory, Directory{Name: subStr[1]})
// 						}
// 					}
// 					// fmt.Printf("goto current index: %v, dir: %+v \n", currenIndex, tempTerminal)
// 				}
// 			} else if subStr[1] != "ls" {
// 				//create file

// 			}

// 		}

// 	}

// 	fmt.Printf("terminal: %+v \n", terminal)

// 	return res
// }
