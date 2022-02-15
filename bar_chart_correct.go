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

func minMax(arr []int) (int, int) {
  left := 0
  right := len(arr) - 1
  minArea := 0
  maxArea := 0

  for  {
    if left >= len(arr) - 1 {
      break
    }
    area := min(arr[left], arr[right]) * (right - left)
    if area <= minArea {
      minArea = area
    }
    if area > maxArea {
      maxArea = area
    }
    left ++
  }
  return minArea, maxArea
}

func main() {
  arr := []int {6, 9, 3, 4, 5, 8}
  min, max := minMax(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {7, 1, 2, 3, 9}
  min, max = minMax(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {}
  min, max = minMax(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {1}
  min, max = minMax(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {3, 4}
  min, max = minMax(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {9, 8, 7, 6, 5, 4, 3, 2, 1}
  min, max = minMax(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
  min, max = minMax(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {4, 8, 1, 2, 3, 9}
  min, max = minMax(arr)
  fmt.Println("min and max areas = ", min, max)
}
