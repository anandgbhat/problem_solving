package main

import (
  "fmt"
)

func binarySearch(nums []int, left int, right int, element int) int {
  if left > right {
    return -1
  }

  mid := (left + right) / 2
  if element == nums[mid] {
    return mid
  }

  if element < nums[mid] {
    // recurse left
    return binarySearch(nums, left, mid-1, element)
  }

  return binarySearch(nums, mid +1, right, element)
}

func main() {
  nums := []int {-1, 0, 1, 2, 3, 4, 7, 9, 10, 20}
  element := 10
  fmt.Printf("Position of %d in %v is %d\n", element, nums, binarySearch(nums, 0, len(nums)-1, element))
}
