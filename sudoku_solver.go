package main

import (
  "fmt"
)

const  MAXROW int = 9
const  MAXCOLUMN int = 9

func printGrid(grid [][]int) {
  for i:=0; i < MAXROW; i++ {
    for j := 0; j < MAXCOLUMN; j++ {
      fmt.Printf("%d ", grid[i][j])
    }
    fmt.Printf("\n")
  }
}

func sudokuSolver(row, column int, grid [][]int) bool {
  // Check if the value is already present.
  // If so, move to the next column and possibly next row
  if row == MAXROW - 1 && column ==  MAXCOLUMN {
    // We are already at the end of the sudoku board
    // So the sudoku board is solvable
    return true
  }

  // Check if we are beyond the end of the column but there are rows yet to
  // be solved
  if column == MAXCOLUMN {
      row ++
      column = 0
  }

  // Check if the value already exists
  // If so, move to next column recursively
  if grid[row][column] > 0 {
    return sudokuSolver(row, column+1, grid)
  }

  // If not already solved, walk through all possible values and back track
  for i := 1; i <= 9; i++ {
    // Check if the value is acceptable
    // If so, recurse for the next column
    if isValid(row, column, grid, i) == true {
      grid[row][column] = i

      if  sudokuSolver(row, column+1, grid) == true {
        return true
      }
    }
    // Not valid
    // Reset value to 0
    grid[row][column] = 0
  }

  return false
}

func isValid(row, column int, grid [][]int, num int) bool {
  // Check if the value is valid by checking the following:
  // 1> value at grid[row][column] should not appear in any other
  // column on that row
  // 2> value at grid[row][column] should not appear in any other
  // row on that column
  // 3> value at grid[row][column] should not appear in any other cell of
  // 3x3 grid to which the [row, column] belong
  // If all of the above are true, return true. Else false
  for i := 0; i < MAXCOLUMN; i++ {
    if grid[row][i] == num {
      return false
    }
  }

  // Check if the column hasn't seen this value in the past
  for i := 0; i < MAXROW; i++ {
    if grid[i][column] == num {
      return false
    }
  }

  // Check the values in the grid
  // Grid's starting row is row - row % 3
  // Grid's starting column is column - column % 3
  startRow := row - row%3
  startColumn := column - column%3
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      // Check if the value is same as the one seen on grid[row][column]
      if grid[i+startRow] [j+startColumn] == num {
        return false
      }
    }
  }
  // All conditions valid
  // Return true
  return true
}

func main() {
  grid := [][]int {{ 3, 0, 6, 5, 0, 8, 4, 0, 0 },
                       { 5, 2, 0, 0, 0, 0, 0, 0, 0 },
                       { 0, 8, 7, 0, 0, 0, 0, 3, 1 },
                       { 0, 0, 3, 0, 1, 0, 0, 8, 0 },
                       { 9, 0, 0, 8, 6, 3, 0, 0, 5 },
                       { 0, 5, 0, 0, 9, 0, 6, 0, 0 },
                       { 1, 3, 0, 0, 0, 0, 2, 5, 0 },
                       { 0, 0, 0, 0, 0, 0, 0, 7, 4 },
                       { 0, 0, 5, 2, 0, 6, 3, 0, 0 } };

  solvable := sudokuSolver(0,0,grid)
  fmt.Println("is solvable = ", solvable)
  printGrid(grid)
}

