package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == ALIVE {
				fmt.Printf("\033[32m%c\033[0m", cell) // set green color for alive cells
			} else {
				fmt.Printf("%c", cell)
			}
		}
		fmt.Println()
	}
}

/* Setup grid with random alive cells (10% chance) */
func setupGrid(grid [][]rune) {
	for row := range grid {
		for col := range grid[row] {
			grid[row][col] = DEAD
			if rand.Intn(100) < 10 {
				grid[row][col] = ALIVE
			}
		}
	}
}

/* Modulo operation with wrap around */
func mod(a, b int) int {
	return (a%b + b) % b
}

/* Count neighbors of a cell in a grid with wrap around */
func countNeighbors(grid [][]rune, row int, col int) int {
	neighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			ni := mod(row+i, NUM_ROWS)
			nj := mod(col+j, NUM_COLS)
			if grid[ni][nj] == ALIVE {
				neighbors++
			}
		}
	}
	return neighbors
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func updateGrid(grid [][]rune) {
	clearScreen()
	// To avoid updating in-place and affecting neighbor counts, use a copy
	newGrid := make([][]rune, NUM_ROWS)
	for i := range newGrid {
		newGrid[i] = make([]rune, NUM_COLS)
		for j := range newGrid[i] {
			if i == 0 && j == 0 {
				continue
			}
			neighbors := countNeighbors(grid, i, j)
			newState := DEAD
			if grid[i][j] == ALIVE {
				if neighbors == 2 || neighbors == 3 {
					newState = ALIVE
				}
			} else {
				if neighbors == 3 {
					newState = ALIVE
				}
			}

			newGrid[i][j] = newState
		}
	}
	// Copy newGrid back to grid
	for i := range grid {
		copy(grid[i], newGrid[i])
	}
}

var (
	NUM_COLS = 50
	NUM_ROWS = 20
	ALIVE    = 'o'
	DEAD     = '-'
)

func main() {

	if len(os.Args) > 1 {
		if val, err := strconv.Atoi(os.Args[1]); err == nil {
			NUM_COLS = val
		}
		if len(os.Args) > 2 {
			if val, err := strconv.Atoi(os.Args[2]); err == nil {
				NUM_ROWS = val
			}
		}
	}

	grid := make([][]rune, NUM_ROWS)
	for i := range grid {
		grid[i] = make([]rune, NUM_COLS)
	}
	setupGrid(grid)
	printGrid(grid)
	for {
		updateGrid(grid)
		printGrid(grid)
		time.Sleep(200 * time.Millisecond)
	}
}
