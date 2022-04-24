package main

import (
  "fmt"
  "sort"
)

// Logic:
// x+y+z = 0 => y+z = -x
// Find y and z that adds up to -x
// Then [-(-x), y, z] will result in 0

var triplets [][]int

func findTripletsZeroSum(nums []int) {
  for i := 0; i < len(nums); i++ {
    if i > 0 && nums[i] == nums[i-1] {
      continue
    }
    result := findPair(nums, -nums[i], i+1)
    if len(result) > 1 {
      triplets = append(triplets, result)
    }
  }
}

func findPair(nums []int, target,left int) []int {
  right := len(nums)-1
  var res []int

  for left < right {
    if nums[left] + nums[right] == target {
      res = append(res, -target)
      res = append(res, nums[left])
      res = append(res, nums[right])
      return res
    } else if nums[left] + nums[right] < target {
      left ++
    } else {
      right--
    }
  }
  return make([]int, 0)
}

func main() {
  nums := []int {-3, 0, 1, 2, -1, 1, -2}
  fmt.Println("Input: ", nums)
  sort.Ints(nums)
  findTripletsZeroSum(nums)
  fmt.Println("Triplet sum to 0  = ", triplets)
}
