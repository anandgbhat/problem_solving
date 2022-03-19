package main

import (
  "fmt"
)

func firstNegativeNumberK(nums []int, k int) []int {
  var result []int
  var negatives []int
  start := 0
  end := 0

  // Find the sliding window first
  for idx := start; idx < k; idx++ {
    if nums[idx] < 0 {
      negatives = append(negatives, nums[start])
      end++
    }
  }

  for end < len(nums) {
    // If some negative numbers found, check if the first value
    // in the negatives slices matches starting of the sliding
    // window. If so, add it to result and move the element
    // out of negatives slice.
    if len(negatives) > 0 && nums[start] == negatives[0] {
      result = append(result, negatives[0])
      negatives = negatives[1:]
    }
    start++
    end++
    // Check if nums[end] is negative as it is a new element
    // in the sliding window
    if end < len(nums) && nums[end] < 0 {
      negatives = append(negatives, nums[end])
    }
  }

  return result
}

func main() {
  nums := []int {-1, 2, 3, -5, -6, -8, 9, 10, 23, -2, -3, 4, -5}
  fmt.Println("input : ", nums)
  fmt.Println(firstNegativeNumberK(nums, 3))
}
