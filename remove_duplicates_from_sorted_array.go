package main

import (
  "fmt"
)

func removeDuplicates(nums []int) int {
  i := 0

  for j:= 1; j < len(nums); j++ {
    // Check if nums[i] and nums[j] are same.
    // If so, increment j
    if nums[i] == nums[j] {
      continue
    }
    // If they are not same, increment i and assign
    // nums[i] = nums[j]
    if nums[i] != nums[j] {
      i++
      nums[i] = nums[j]
    }
  }

  return i+1
}

func main() {
  nums := []int{1,1,2,2,2,3,3}
  fmt.Println("Length of unique elements = ", removeDuplicates(nums))
}

