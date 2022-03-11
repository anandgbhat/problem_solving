// Given an array and a sum, return number of subsets that sum upto the value

package main

import (
  "fmt"
)



func numSubsets(nums []int, sum int, dp [][]int) int {
  // Main logic:
  // We can either include the value into the subset or not
  // If we include, then deduct the value from the currenr value.
  for i:=1; i < len(nums)+1; i++ {
    for j:=1; j < (sum+1); j++ {
      if nums[i-1] <= j {
        val := j-nums[i-1]
        dp[i][j] = dp[i-1][j] + dp[i-1][val]
      } else {
        dp[i][j] = dp[i-1][j]
      }
    }
  }

  return dp[len(nums)][sum]
}

func main() {
  nums := []int {2, 3, 5, 6, 8, 10}
  sum := 10
  dp := make([][]int, len(nums)+1)
  for i := range dp {
    dp[i] = make([]int, sum+1)
  }

  // Initialize the dp matrix
  // Matrix size is len(nums)+1 X sum+1
  // Initialize first row and column
  // First row is when number of sum is 0. There is no way we can have
  // any subsets with 0 sum. So full row is 0
  // First coulumn is when array len is 0 . There is always
  // {} that results in sum 0. So full column is 1
  for i:=0; i < len(nums)+1; i++ {
    for j:=0; j < sum+1; j++ {
      if j==0 {
        dp[i][j] = 1
      } else if i==0 {
        dp[i][j] = 0
      }
    }
  }

  fmt.Println(numSubsets(nums, sum, dp))
}


