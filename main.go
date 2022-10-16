package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return
the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or
vertically. You may assume all four edges of the grid are all surrounded by water.
*/

/*
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
*/

/*
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
*/

/*
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

type Pair struct {
	x int
	y int
}

var logger *zap.Logger

func (p *Pair) String() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func isValidPair(pair Pair, width int, height int) bool {
	return pair.x >= 0 && pair.x < width && pair.y >= 0 && pair.y < height
}

func getAdjacentCells(cell Pair, width int, height int) []Pair {
	left := Pair{x: cell.x - 1, y: cell.y}
	right := Pair{x: cell.x + 1, y: cell.y}
	up := Pair{x: cell.x, y: cell.y - 1}
	down := Pair{x: cell.x, y: cell.y + 1}
	validPairs := make([]Pair, 0, 4)
	if isValidPair(left, width, height) {
		validPairs = append(validPairs, left)
	}
	if isValidPair(right, width, height) {
		validPairs = append(validPairs, right)
	}
	if isValidPair(up, width, height) {
		validPairs = append(validPairs, up)
	}
	if isValidPair(down, width, height) {
		validPairs = append(validPairs, down)
	}

	return validPairs
}

func visit(grid [][]byte, pair Pair, visited map[string]struct{}) {
	logger.Sugar().Debugf("Visiting x: %v, y: %v", pair.x, pair.y)
	current := grid[pair.y][pair.x]
	if current == 1 {
		// This is an island
		// Lets visit all parts of the island using DFS
		logger.Debug("Land")
		visited[pair.String()] = struct{}{}
		adjacentCells := getAdjacentCells(pair, len(grid[0]), len(grid))
		logger.Sugar().Debugf("Adjacent cells: %v\n", adjacentCells)
		for _, p := range adjacentCells {
			logger.Sugar().Debugf("Visited: %v", visited)
			if _, found := visited[p.String()]; !found {
				visit(grid, p, visited)
			}
		}
	} else {
		logger.Debug("Water")
	}
}

func numIslands(grid [][]byte) int {
	visited := map[string]struct{}{}
	countOfIslands := 0
	for row := range grid {
		for col := range grid[row] {
			pair := Pair{x: col, y: row}
			if grid[row][col] == 0 {
				continue
			}
			if _, found := visited[pair.String()]; !found {
				countOfIslands++
				visit(grid, pair, visited)
			}
		}
	}

	return countOfIslands
}

func main() {
	debugEnv := os.Getenv("DEBUG")
	if debugEnv == "1" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	defer zap.RedirectStdLog(logger)

	var grid [][]byte
	grid = [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	fmt.Println(numIslands(grid)) // Should be one

	grid = [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(numIslands(grid)) // Should be three
}
