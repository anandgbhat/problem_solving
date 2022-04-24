package main

import (
  "fmt"
)

func removeDuplicates(nums []int, element int) []int {
  // Two pointer approach
  // One pointer traverses the array and next pointer keeps track of last
  // valid element
  lastValid := 0

  for i:=0; i < len(nums); i++ {
    // Copy the first element that's != element to nums[lastValid]
    // lastValid tracks the last element in the final array.
    if nums[i] != element {
      nums[lastValid] = nums[i]
      lastValid++
    }
  }

  return nums[0:lastValid]
}

func main() {
  nums := []int {3, 2, 3, 6, 3, 10, 9, 3}
  fmt.Println(nums)
  fmt.Println("Array after removing 3 is ",
   removeDuplicates(nums, 3))
}
