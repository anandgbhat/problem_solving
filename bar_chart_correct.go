// Find the maximum area formed in the bar graph

package main

import (
  "fmt"
)

func min(i,j int) int {
  if i < j {
    return i
  }

  return j
}

// Logic:
// Use two indixes: left and right
// left initialized to first element and right initialized to last element
// Iterate over the array and find the area formed between left and right
// Area is calculated as min(arr[left], arr[right]) * (right - left)
// increment left and update max if bigger area is found

func maxArea(arr []int) int {
  left := 0
  right := len(arr) - 1
  maxArea := 0

  for  {
    if left >= len(arr) - 1 {
      break
    }
    area := min(arr[left], arr[right]) * (right - left)
    if area > maxArea {
      maxArea = area
    }
    left ++
  }
  return  maxArea
}

func main() {
  arr := []int {6, 9, 3, 4, 5, 8}
  fmt.Println("max area = ", maxArea(arr))
}
