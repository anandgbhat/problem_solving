// Find the total water that can be trapped given an elevation graph with each value width being 1

package main

import (
  "fmt"
)

func min(i, j int) int {
  if i < j {
    return i
  }

  return j
}

func rainWaterTrapping(arr []int) int {
  // For each value, calculate maxLeft and maxRight.
  // Water that can be held on top of this index value is
  // min(maxRight, maxLeft) - arr[idx]
  // Repeat this for each value
  fmt.Println(arr)

  maxLeft := 0
  maxRight := 0
  idx := 0
  total := 0
  for ; idx <= len(arr) -1 ; idx ++ {
    maxLeft , maxRight = 0, 0
    // Calculate maxRight
    for j := idx+1; j <= len(arr) - 1; j++ {
      if arr[j] > maxRight {
        maxRight = arr[j]
      }
    }
    // Calculate maxLeft
    for k := idx-1; k >= 0 ; k-- {
      if arr[k] > maxLeft {
        maxLeft = arr[k]
      }
    }
    // Calculate the water that can be held on top of this block
    waterOnTheBlock := min(maxRight, maxLeft) - arr[idx]
    if waterOnTheBlock > 0 {
      total += waterOnTheBlock
    }
  }

  return total
}

func main() {
  arr := []int {0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
  fmt.Printf("Water trapping of %v is %d\n", arr, rainWaterTrapping(arr))

  arr = []int {5, 0, 3, 0, 0, 0, 2, 3, 4, 2, 1}
  fmt.Printf("Water trapping of %v is %d\n", arr, rainWaterTrapping(arr))
}
