package main

import (
	"fmt"
	"math/rand"
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
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

func (matrix Universe) Seed() Universe {
	col := width * height / 4

	for i := range matrix {
		for j := range matrix[i] {
			b := rand.Intn(4)
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

func (matrix Universe) SelsNeighbors(x, y int) int {
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

func main() {
	m := NewUniverse()
	m.Seed().Show()
	fmt.Println(m.SelsNeighbors(1, 1))
	fmt.Println(m.IsCellAlive(1, 1))
}
