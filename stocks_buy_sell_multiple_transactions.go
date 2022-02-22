package main

import (
  "fmt"
)

func maxProfitMultipleTransactions(prices []int) int {
  profit := 0

  for i:=1; i<len(prices); i++ {
    if prices[i] > prices[i-1] {
      profit += (prices[i] - prices[i-1])
    }
  }

  return profit
}

func main() {
  prices := []int {7,1,5,3,6,4}
  fmt.Println("Max profit with multiple transactions = ",
    maxProfitMultipleTransactions(prices))
}
