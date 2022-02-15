// Program to calculate maximum area that can hold water in a bar chart
// Logic is to start with left most and right most sides of the array and
// calculate the area. Move the left pointer if value at left is <= right.
// Otherwize move right.
// Loop through all values till we hit end of array
// Store max value in each iteration.
// Area = min(value_on_left, value_on_right) *
//                                      abs(index_on_right - index_on_left)

package main

import (
  "fmt"
)

func min(i,j int) int {
  if i <= j {
    return i
  } else {
    return j
  }
}

func abs(i int) int {
  if i < 0 {
    return -i
  } else {
    return i
  }
}



func minMaxArea(numbers []int) (int, int) {
  // base case
  if len(numbers) <=1 {
    return 0, 0
  }

  // Temporary variables to hold current values and maximum area
  left := 0
  right := len(numbers) - 1
  var maxArea int
  var minArea int

  // Loop through all the numbers in the array and save maximum area
  for {
    if left >= len(numbers) -1  || right <= 0 {
      break
    }
    area := min(numbers[left], numbers[right])  * abs(right - left)
    if area > maxArea {
      maxArea = area
    } else if area <= minArea {
      minArea = area
    }
    if numbers[left] <= numbers[right] {
      left ++
    } else {
      right --
    }
  }

  return minArea, maxArea
}

func main() {
  arr := []int {6, 9, 3, 4, 5, 8}
  min, max := minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {7, 1, 2, 3, 9}
  min, max = minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {}
  min, max = minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {1}
  min, max = minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {3, 4}
  min, max = minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {9, 8, 7, 6, 5, 4, 3, 2, 1}
  min, max = minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
  min, max = minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)

  arr = []int {4, 8, 1, 2, 3, 9}
  min, max = minMaxArea(arr)
  fmt.Println("min and max areas = ", min, max)
}
