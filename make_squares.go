package main

import (
  "fmt"
)

func makeSquares(nums []int) []int {

  // Logic: 2 pointer method
  // Initialize left to 0 and right to len-1
  // Calculate square of num @left and num @right
  // Move pointers depending on which value is bigger
  left := 0
  right := len(nums) -1
  lastIndex := len(nums) -1
  output := make([]int, len(nums))
  leftVal, rightVal := 0,0

  for left <= right {
    leftVal = nums[left] * nums[left]
    rightVal = nums[right] * nums[right]
    if leftVal > rightVal {
      output[lastIndex] = leftVal
      lastIndex--
      left ++
    } else {
      output[lastIndex] = rightVal
      lastIndex--
      right --
    }
  }

  return output
}

func main() {
  nums := []int{-2, -1, 0, 2, 3}

  fmt.Println("Squared array in sorted order = ", makeSquares(nums))
}


