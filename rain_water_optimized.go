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

func rainWater(array []int) int {
  left := 0
  right := len(array) - 1
  maxLeft := 0
  maxRight := 0
  total := 0

  for {
    // Check if there is a wall on either left or right
    // This depends on whether array[left] is smaller or array[right] is smaller
    // If array[left] is smaller, check if maxLeft is > array[left]
    // If so, update maxLeft and move left by 1
    // If current value is > maxLeft, amount of water on the current sq =
    // diff (maxLeft, currentValue)
    // Same logic applies to right side
    //

    if right <= left {
      break
    }

   fmt.Println(array[left], array[right], maxLeft, maxRight)
    if array[left] <= array[right] {
      if maxLeft < array[left] {
        maxLeft = array[left]
        left ++
      } else {
        total += maxLeft - array[left]
        left ++
      }
    } else {
      if maxRight < array[right] {
        maxRight = array[right]
        right --
      } else {
        total += maxRight - array[right]
        right --
      }
    }
  }

  return total
}

func main() {
  array := []int {0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
  fmt.Println("Amount of water = ", rainWater(array))
}

