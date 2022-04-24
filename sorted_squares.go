package main

import (
  "fmt"
)

func sortedSquares(nums []int) []int {
  sortedSlice := make([]int, len(nums))
  leftIdx, rightIdx := 0 , len(nums)-1
  lastIdx := len(nums)-1
  leftSquare, rightSquare := 0, 0

  for leftIdx < rightIdx {
    leftSquare = nums[leftIdx] * nums[leftIdx]
    rightSquare = nums[rightIdx] * nums[rightIdx]

    if rightSquare > leftSquare {
      sortedSlice[lastIdx] = rightSquare
      rightIdx --
    } else {
      sortedSlice[lastIdx] = leftSquare
      leftIdx++
    }
    lastIdx -= 1
  }

  return sortedSlice
}

func main() {
  nums := []int {-2, -1, 0, 2, 3}
  fmt.Println("Sorted slice of ", nums, " = ", sortedSquares(nums))
}


