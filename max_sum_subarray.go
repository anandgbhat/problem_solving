package main

import (
  "fmt"
)

func maxSumSubarray(nums []int, numElements int)  int {
  // Logic : have 2 pointers- left and right.
  // right = left +numElements
  // Calculate the sum and slide the window by 1 till we hit the end
  left := 0
  right := left + numElements - 1 // Inclusive
  maxSum := 0

  for right < len(nums) {
    sum := 0
    // Calculate the sum
    for i:= left; i <= right; i++ {
      sum += nums[i]
    }
    if maxSum < sum {
      maxSum = sum
    }
    left ++
    right ++
  }

  return maxSum
}

func main() {
  nums := []int {2, 1, 5, 1, 3, 2}
  fmt.Printf("Max sum in a sliding window of 3 in %v = %d\n",
    nums, maxSumSubarray(nums, 3))

  nums = []int {2, 3, 4, 1, 5}
  fmt.Printf("Max sum in a sliding window of 3 in %v = %d\n",
    nums, maxSumSubarray(nums, 3))
}
