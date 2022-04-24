package main

import (
  "fmt"
  "math"
  "sort"
)

func abs(i int) int {
  if i < 0 {
    return -i
  }
  return i
}


func findTripletWithSmallestDifference(nums []int, target int) int {
  smallestDiffSoFar := math.MaxInt
  sort.Ints(nums)
  var diff int
  for i := 0; i < len(nums)-1; i++ {
    left := i+1
    right := len(nums)-1
    for left < right {
      curSum := nums[i] + nums[left] + nums[right]
      diff = target - curSum
      if diff == 0 {
        return nums[i] + nums[left] + nums[right]
      } else if diff < 0 {
        right --
      } else {
        left++
      }
      if abs(diff) < abs(smallestDiffSoFar) ||
          abs(diff) == abs(smallestDiffSoFar) && diff > smallestDiffSoFar {
        smallestDiffSoFar = diff
      }
    }
  }

  return target - smallestDiffSoFar
}

func main() {
  nums := []int {-3, -1, 1, 2}
  target := 1
  fmt.Println("Minimum diff close to ", target, " in ", nums, "= ",
    findTripletWithSmallestDifference(nums, target))

  nums = []int {-2, 0, 1, 2}
  target = 2
  fmt.Println("Minimum diff close to ", target, " in ", nums, "= ",
    findTripletWithSmallestDifference(nums, target))

  nums = []int{1, 0, 1, 1}
  target = 100
  fmt.Println("Minimum diff close to ", target, " in ", nums, "= ",
    findTripletWithSmallestDifference(nums, target))
}





