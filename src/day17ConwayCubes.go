package src

import (
	"fmt"

	"nikpappas.com/adventofcode2021/files"
	"nikpappas.com/adventofcode2021/maps"
)

type ConwayCube struct {
	board map[int](map[int](map[int](string)))
	size  int
}
type ConwayHyperCube struct {
	board map[int](map[int](map[int](map[int](string))))
	size  int
}

func Day17() {
	lines := files.ReadLines("inputs/day17.txt")
	boardSlice := maps.MapLinesToBoard(lines)
	cube := ConwayCube{nil, 0}

	for i, line := range boardSlice {
		for j, val := range line {
			cube = putInCube(cube, i-len(boardSlice)/2, j-len(line)/2, 0, val)
		}
	}

	live := day17sol1(cube, 6)
	fmt.Println(live)
	hyper := ConwayHyperCube{nil, 0}

	for i, line := range boardSlice {
		for j, val := range line {
			hyper = putInHyperCube(hyper, i-len(boardSlice)/2, j-len(line)/2, 0, 0, val)
		}

	}
	live = day17sol2(hyper, 6)
	fmt.Println(live)
}

func day17sol1(cube ConwayCube, cycles int) int {
	for cycle := 0; cycle < cycles; cycle++ {
		buffer := cloneCube(cube)
		printCube(cube)
		fmt.Println(countAlive(cube), "===============================")
		size := buffer.size + 1
		fmt.Println("'size'", size)
		for i := -size; i <= size; i++ {
			for j := -size; j <= size; j++ {
				for k := -size; k <= size; k++ {
					aliveNeighbours := countAliveNeighbours(buffer, i, j, k)
					isCurAlive := isAlive(buffer, i, j, k)
					if isCurAlive && !(aliveNeighbours == 2 || aliveNeighbours == 3) {
						cube = putInCube(cube, i, j, k, ".")
					}
					if !isCurAlive && aliveNeighbours == 3 {
						cube = putInCube(cube, i, j, k, "#")
					}
				}
			}
		}

	}

	return countAlive(cube)
}
func day17sol2(cube ConwayHyperCube, cycles int) int {
	for cycle := 0; cycle < cycles; cycle++ {
		buffer := cloneHyperCube(cube)
		fmt.Println(countAliveHyper(cube), "===============================")
		size := buffer.size + 1
		fmt.Println("'size'", size)
		for i := -size; i <= size; i++ {
			for j := -size; j <= size; j++ {
				for k := -size; k <= size; k++ {
					for l := -size; l <= size; l++ {
						aliveNeighbours := countAliveNeighboursHyper(buffer, i, j, k, l)
						isCurAlive := isAliveHyper(buffer, i, j, k, l)
						if isCurAlive && !(aliveNeighbours == 2 || aliveNeighbours == 3) {
							cube = putInHyperCube(cube, i, j, k, l, ".")
						}
						if !isCurAlive && aliveNeighbours == 3 {
							cube = putInHyperCube(cube, i, j, k, l, "#")
						}
					}
				}
			}
		}

	}

	return countAliveHyper(cube)
}

func countAliveNeighbours(cube ConwayCube, i, j, k int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				if isAlive(cube, i+x, j+y, k+z) {
					count++
				}
			}
		}

	}
	return count

}
func countAliveNeighboursHyper(cube ConwayHyperCube, i, j, k, l int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					if isAliveHyper(cube, i+x, j+y, k+z, l+w) {
						count++
					}
				}
			}
		}

	}
	return count

}

func putInCube(cube ConwayCube, i, j, k int, val string) ConwayCube {

	if abs(i) > cube.size {
		fmt.Println("size i->", abs(i))
		cube.size = abs(i)
	}
	if abs(j) > cube.size {
		fmt.Println("size j->", j)
		cube.size = abs(j)
	}
	if abs(k) > cube.size {
		fmt.Println("size k->", k)
		cube.size = abs(k)
	}
	if cube.board == nil {
		cube.board = make(map[int](map[int](map[int](string))))
	}
	_, ok := cube.board[i]
	if !ok {
		cube.board[i] = make(map[int](map[int](string)))
	}
	_, ok = cube.board[i][j]
	if !ok {
		cube.board[i][j] = make(map[int](string))
	}
	cube.board[i][j][k] = val
	return cube
}
func putInHyperCube(cube ConwayHyperCube, i, j, k, l int, val string) ConwayHyperCube {

	if abs(i) > cube.size {
		cube.size = abs(i)
	}
	if abs(j) > cube.size {
		cube.size = abs(j)
	}
	if abs(k) > cube.size {
		cube.size = abs(k)
	}
	if abs(l) > cube.size {
		cube.size = abs(l)
	}
	if cube.board == nil {
		cube.board = make(map[int](map[int](map[int](map[int](string)))))
	}
	_, ok := cube.board[i]
	if !ok {
		cube.board[i] = make(map[int](map[int](map[int](string))))
	}
	_, ok = cube.board[i][j]
	if !ok {
		cube.board[i][j] = make(map[int](map[int](string)))
	}
	_, ok = cube.board[i][j][k]
	if !ok {
		cube.board[i][j][k] = make(map[int](string))
	}
	cube.board[i][j][k][l] = val
	return cube
}

func isAlive(cube ConwayCube, i, j, k int) bool {
	_, ok := cube.board[i]
	if !ok {
		return false
	}
	_, ok = cube.board[i][j]
	if !ok {
		return false
	}
	v, ok := cube.board[i][j][k]
	if !ok {
		return false
	}
	if v != "#" {
		return false
	}
	return true

}
func isAliveHyper(cube ConwayHyperCube, i, j, k, l int) bool {
	_, ok := cube.board[i]
	if !ok {
		return false
	}
	_, ok = cube.board[i][j]
	if !ok {
		return false
	}
	_, ok = cube.board[i][j][k]
	if !ok {
		return false
	}
	v, ok := cube.board[i][j][k][l]
	if !ok {
		return false
	}
	if v != "#" {
		return false
	}
	return true

}

func cloneCube(cube ConwayCube) ConwayCube {
	fmt.Println("clonging ")
	var clone ConwayCube
	for i, plane := range cube.board {
		for j, line := range plane {
			for k, val := range line {
				clone = putInCube(clone, i, j, k, val)
			}
		}
	}
	return clone
}
func cloneHyperCube(hyper ConwayHyperCube) ConwayHyperCube {
	fmt.Println("clonging ")
	var clone ConwayHyperCube
	for i, cube := range hyper.board {
		for j, plane := range cube {
			for k, line := range plane {
				for l, val := range line {
					clone = putInHyperCube(clone, i, j, k, l, val)
				}
			}
		}
	}
	return clone
}
func printCube(cube ConwayCube) {
	size := cube.size

	for i := -size; i <= size; i++ {
		fmt.Println("Z", i, "+=======================")
		for j := -size; j <= size; j++ {
			for k := -size; k <= size; k++ {
				if isAlive(cube, i, j, k) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println("+=======================")
	}
}

func countAlive(cube ConwayCube) int {
	count := 0
	for _, plane := range cube.board {
		for _, line := range plane {
			for _, val := range line {
				if val == "#" {
					count++
				}
			}
		}
	}
	return count
}
func countAliveHyper(hyper ConwayHyperCube) int {
	count := 0
	for _, cube := range hyper.board {
		for _, plane := range cube {
			for _, line := range plane {
				for _, val := range line {
					if val == "#" {
						count++
					}
				}
			}
		}
	}
	return count
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
