package main

import (
  "fmt"
)


var result [][]int

func findSubArrays(nums []int, target int) [][]int {
  // basic checks
  if len(nums) == 0 || target < 0 {
    return [][]int(nil)
  }

  result = nil

  // Add single elements into subarray as long as they are less than target
  for _, v := range nums {
    if v < target {
      r := make([]int, 0)
      r = append(r, v)
      result = append(result, r)
    }
  }

  // Sliding window
  left := 0
  product := nums[left]
  for right:=left+1; right < len(nums) ; right ++ {
    product *= nums[right]
    if product >= target {
      product /= nums[left]
      left++
      if product < target {
        addToResult(nums, left, right)
      }
    } else {
      addToResult(nums, left, right)
    }
  }

  return result
}

func addToResult(nums []int, left, right int) {
  var r []int

  for j := left; j <= right; j++ {
    r = append(r, nums[j])
  }
  result = append(result, r)
}

func main() {
  nums := []int{2, 5, 3, 10}
  target := 30
  fmt.Println("Subarrays in ", nums, " whose product is less than ",
    target, " = ", findSubArrays(nums, target))
  nums = []int{8, 2, 6, 5}
  target = 50
  fmt.Println("Subarrays in ", nums, " whose product is less than ",
    target, " = ", findSubArrays(nums, target))

}
