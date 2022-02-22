package main

import (
  "fmt"
  "math"
)

func maxProfit(prices []int) int {
  profit := 0
  minLeft := math.MaxInt
  for _, val := range prices {
    if val < minLeft {
      minLeft = val
    }
    curProfit := val - minLeft
    if profit < curProfit {
      profit = curProfit
    }
  }
  return profit
}

func main() {
  nums := []int{7,1,5,3,6,4}
  fmt.Println("Max profit = ", maxProfit(nums))
}
