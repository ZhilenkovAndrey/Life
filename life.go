package main

import (
	"fmt"
	"math/rand"
)

type Universe [][]bool

const (
	width  = 80
	height = 15
)

func NewUniverse() Universe {
	matrix := make(Universe, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
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
			fmt.Println(b)
			if b == 1 && col != 0 {
				matrix[i][j] = true
				col--
			}
		}
	}
	return matrix
}

func main() {
	NewUniverse().Seed().Show()
}
