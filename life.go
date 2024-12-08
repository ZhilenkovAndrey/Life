package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Universe [][]bool

const (
	width  = 80
	height = 20
)

func NewUniverse() Universe {
	matrix := make(Universe, width)
	for i := range matrix {
		matrix[i] = make([]bool, height)
	}
	return matrix
}

func (matrix Universe) ToString() string {
	b := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if matrix[x][y] {
				b = append(b, '*')
			} else {
				b = append(b, ' ')
			}
		}
		b = append(b, '\n')
	}
	return string(b)
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (matrix Universe) Show() {
	ClearScreen()
	fmt.Println()
	fmt.Print(matrix.ToString())
}

func (matrix Universe) Seed() Universe {
	col := (width + 1) * height / 10

	for i := range matrix {
		for j := range matrix[i] {
			b := rand.Intn(10)
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
	return matrix.CellsNeighbors(x, y) == 3 || (matrix.IsCellAlive(x, y) &&
		matrix.CellsNeighbors(x, y) == 2)
}

func NextStep(a, b Universe) Universe {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[j][i] = a.NextGeneration(j, i)
		}
	}
	return b
}

func main() {
	universeFirst := NewUniverse()
	universeNext := NewUniverse()

	universeFirst.Seed()

	for i := 0; i < 150; i++ {
		universeFirst.Show()
		universeFirst = NextStep(universeFirst, universeNext)
		time.Sleep(time.Second / 20)
	}
}
