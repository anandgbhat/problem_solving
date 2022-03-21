package main

import (
  "fmt"
  "math"
)

func maxSubArraySum(nums []int) int {
  maxSoFar := math.MinInt
  maxEndingHere := 0

  for i:=0; i< len(nums); i++ {
    maxEndingHere = maxEndingHere + nums[i]
    if maxSoFar < maxEndingHere {
      maxSoFar = maxEndingHere
    }

    if maxEndingHere < 0 {
      maxEndingHere = 0
    }
  }

  return maxSoFar
}


func main() {
  fmt.Println("Kadane's algorithm")
  nums := []int {-1, -3, 4, -1, -2, 1, 5, -3}
  fmt.Println("max subarray sum of ", nums, " = ", maxSubArraySum(nums))
}
