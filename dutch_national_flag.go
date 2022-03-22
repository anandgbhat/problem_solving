package main

import (
  "fmt"
)

func swap(i,j *int) {
  tmp := *i
  *i = *j
  *j = tmp
}

// 0 - RED
// 1 - WHITE
// 2 - BLUE

func sortColors(nums []int) []int {
  // Three pointers
  // start is the rightmost boundary of 0s
  // end is leftmost boundary of 2s
  // Any index below start is 0
  // Any index after end is 2
  // Swap value of nums[curr] with nums[start] if it is 0 and increment both
  // Swap value of nums[curr] with nums[end] if it is 2 and decrement end
  start := 0
  end := len(nums) -1
  curr := 0

  for curr <= end {
    if nums[curr] == 0 {
      swap(&nums[curr], &nums[start])
      curr ++
      start++
    } else if nums[curr] == 2 {
      swap(&nums[curr], &nums[end])
      end--
    } else {
      curr++
    }
  }
  return nums
}

func main() {
  nums := []int {2,0,2,1,1,0}
  nums = sortColors(nums)
  fmt.Println("Final array = ", nums)
}
