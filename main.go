// Conway's game of life.
// 2 dimensional grid of cells.
// Each cell has 8 adjacent cells in horizontal, vertical and diagonal directions.
// Numbers of generations and in each generation, a cell can live or die based on neighbour.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	umap := make(Universe, height)
	for i := range umap {
		umap[i] = make([]bool, width)
	}
	return umap
}

func (u Universe) Show() {
	for _, row := range u {
		for _, cell := range row {
			switch {
			case cell:
				fmt.Printf("*")
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (u Universe) Seed() {
	for _, row := range u {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (width + x) % width
	y = (height + y) % height
	return u[y][x]
}

func (u Universe) Neighbours(x, y int) int {
	var neighbours int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if j == x && i == y {
				continue
			}
			if u.Alive(j, i) == true {
				neighbours++
			}
		}
	}
	return neighbours
}
func (u Universe) Next(x, y int) bool {
	n := u.Neighbours(x, y)
	alive := u.Alive(x, y)
	if n > 1 && n < 4 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}
func Step(a, b Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(j, i)
		}
	}
}

func main() {
	fmt.Println("\033c\x0c")
	rand.Seed(time.Now().UTC().UnixNano())
	newUniverse := NewUniverse()
	nextUniverse := NewUniverse()
	newUniverse.Seed()
	for {
		newUniverse.Show()
		Step(newUniverse, nextUniverse)
		newUniverse, nextUniverse = nextUniverse, newUniverse
		time.Sleep(50 * time.Millisecond)
		fmt.Println("\033c\x0c")
	}
}
