package main

import (
  "fmt"
  "math"
)


var min int
var temp int


func max(i,j int) int {
  if i > j {
    return i
  }
  return j
}

// There are f floors and e eggs
// Start dropping the egg at kth floor.
// If egg breaks, remaining eggs are e-1 and threshold floor should be
// between 1 to k-1.
// If egg does not break, then eggs remain same. Remaining floors to test
// is f-k as threshold floor should be above k
// Take the minimum of these two recursive calls
func minAttempts(f,e int) int {
  // Base condition
  if f == 0 || f == 1 {
    return f
  }
  // If there is only one egg, at max f attempts can be made
  // Will start at lowest floor and keep going up till threshold floor
  // is reached which has upper bound of f.
  if e==1 {
    return f
  }

  for k:=1; k<= f; k++ {
    // One attempt is already  done at each iteration
    // Get the maximum of both choices and update min
    temp = 1+ max(minAttempts(f-k, e), minAttempts(k-1, e-1))
    if min > temp {
      min = temp
    }
  }

  return min
}


func main() {
  min = math.MaxInt
  fmt.Println("Minimum attempts for 5 floors and 3 eggs is = ",
    minAttempts(5, 3))
}
