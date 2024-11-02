package main

import (
	. "fmt"
	. "math/rand"
	_ "time"
)

func main() {
	// write your code here
	var size int
	Scan(&size)
	Evolve(size)
}

func CreateUniverse(size int) map[int]map[int]string {
	universe := make(map[int]map[int]string, size)

	// create the universe
	for i := 0; i < size; i++ {
		universe[i] = make(map[int]string, size)
		for j := 0; j < size; j++ {
			random := Intn(2)
			if random == 1 {
				universe[i][j] = "O"
			} else {
				universe[i][j] = " "
			}
		}
	}
	return universe
}

func getCurrentState(size int) [][]string {
	universe := CreateUniverse(size)
	currentStateArr := make([][]string, size)

	// current state
	for i := range currentStateArr {
		innerCurr := make([]string, size)
		for j := range innerCurr {
			innerCurr[j] = universe[i][j]
		}
		currentStateArr[i] = innerCurr
	}
	return currentStateArr
}

// uniformal counting of live neighbors per generation
func countLiveNeighbors(universe [][]string, x, y, size int) int {
	count := 0
	// these are rows
	for i := -1; i <= 1; i++ {
		// these are columns
		for j := -1; j <= 1; j++ {
			// check we are not including the cell itself whose neighbors we need to count for
			if !(i == 0 && j == 0) {
				ni := (x + i + size) % size
				nj := (y + j + size) % size
				if universe[ni][nj] == "O" {
					count++
				}
			}
		}
	}
	return count
}

func Evolve(size int) {
	current := getCurrentState(size)
	count := 1
	// next state 2D array is re-generated upon next generation
	for count < 20 {
		// dur := Duration(Intn(1000)) * Millisecond
		nextStateArr := make([][]string, size)

		Printf("Generation: #%d\n", count)
		var aliveCells int
		for i := range nextStateArr {
			// we generate a fresh inner slice upon the actual row of the next-gen slice we iterate over
			nextStateArr[i] = make([]string, size)
			for j := range nextStateArr[i] {
				// this line declares and re-declares at the same time upon gen iteration
				liveNeighbors := countLiveNeighbors(current, i, j, size)

				// we then apply the evolved cells into the fresh inner slice
				if current[i][j] == "O" {
					if liveNeighbors < 2 || liveNeighbors > 3 {
						nextStateArr[i][j] = " "
					} else {
						nextStateArr[i][j] = "O"
						aliveCells++
					}
				} else {
					if liveNeighbors == 3 {
						nextStateArr[i][j] = "O"
						aliveCells++
					} else {
						nextStateArr[i][j] = " "
					}
				}
			}
		}
		Printf("Alive: %d\n", aliveCells)
		// apply the new independent copies into the current state
		current = nextStateArr
		for i := range current {
			for j := range current[i] {
				Printf("%s", current[i][j])
			}
			Println()
		}
		// Sleep(dur)
		count++
	}
}
