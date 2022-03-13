package main

import (
  "fmt"
  "math"
)


var min int

func solve(nums []int, i,j int) int {
  // Base condition
  if i>= j {
    return 0
  }

  // Loop through all possible split values and calculate the
  // minimum cost
  for k:=i; k<=j-1; k++ {
    temp := solve(nums, i, k) + solve(nums, k+1, j) +
      nums[i-1] * nums[k] * nums[j]

    if temp < min {
      min = temp
    }
  }

  return min
}

func main() {
  nums := []int {40, 20, 30, 10, 30}
  min = math.MaxInt

  fmt.Println("Minimum cost of multiplication = ", solve(nums, 1, len(nums)-1))
}
