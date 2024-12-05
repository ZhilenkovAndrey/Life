package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Universe [][]bool

const (
	width  = 20
	height = 15
)

func NewUniverse() Universe {
	matrix := make(Universe, width)
	for i := range matrix {
		matrix[i] = make([]bool, height)
	}
	return matrix
}

func (matrix Universe) Show() {
	fmt.Println()
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

func (matrix Universe) Seed() Universe {
	col := width * height / 2

	for i := range matrix {
		for j := range matrix[i] {
			b := rand.Intn(2)
			if b == 1 && col != 0 {
				matrix[i][j] = true
				col--
			}
		}
	}
	return matrix
}

func (matrix Universe) IsCellAlive(x, y int) bool {
	if x < 0 {
		x = (x % width) + width
	}

	if y < 0 {
		y = (y % height) + height
	}

	if x >= width {
		x = x % width
	}

	if y >= height {
		y = y % height
	}

	return matrix[x][y]
}

func (matrix Universe) CellsNeighbors(x, y int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) && matrix.IsCellAlive(x+j, y+i) {
				sum++
			}
		}
	}
	return sum
}

func (matrix Universe) NextGeneration(x, y int) bool {
	if matrix.IsCellAlive(x, y) == true &&
		(matrix.CellsNeighbors(x, y) <= 2 || matrix.CellsNeighbors(x, y) >= 3) {
		matrix[x][y] = false
	} else {
		matrix[x][y] = true
	}

	if matrix.IsCellAlive(x, y) == false && matrix.CellsNeighbors(x, y) == 3 {
		matrix[x][y] = true
	} else {
		matrix[x][y] = false
	}

	return matrix[x][y]
}

func NextStep(a, b Universe) Universe {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			b[i][j] = a.NextGeneration(i, j)
		}
	}
	return b
}

func main() {
	universeFirst := NewUniverse()
	universeNext := NewUniverse()

	universeFirst.Seed().Show()
	time.Sleep(3 * time.Second)

	for i := 0; i < 10; i++ {
		universeFirst = NextStep(universeFirst, universeNext)
		universeFirst.Seed().Show()
		time.Sleep(3 * time.Second)
	}
}
