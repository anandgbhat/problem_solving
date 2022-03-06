package main

import (
  "fmt"
)

func max(i,j int) int {
  if i > j {
    return i
  }
  return j
}

func knapsack(weight []int, value []int, capacity int, n int) int {
  // Base condition
  // Smallest array size as base condition
  if capacity == 0 || n == 0 {
    return 0
  }

  // Choice tree
  // If weight[n-1] < capacity, then it can either be put in knapsack or
  // ignored. Choice is based on what yields max profit.
  // If weight[n-1] > capacity, item cannot be included

  if weight[n] <= capacity {
    fmt.Println(capacity)
    return max(value[n]+knapsack(weight, value, capacity-weight[n], n-1),
      knapsack(weight, value, capacity, n-1))
  } else {
    return knapsack(weight, value, capacity, n-1)
  }
}


func main() {
  weight := []int {30, 20, 20}
  value := []int {10, 20, 40}
  capacity := 40


  fmt.Println("Maximum profit = ", knapsack(weight, value, capacity, len(weight)-1))
}
