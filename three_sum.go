// Find the triplet that sums upto a given number
// Time complexity: O(n^2)
// Space complexity: O(n)

package main

import (
  "fmt"
)

func threeSum(array []int, target int) []int {
  m := make(map[int]int)

  // Create the hash table
  for _, val := range array {
    m[val] = val
  }

  // Run two loops and check if target - m[i] - m[j] exists in the hash
 // If so, return the numbers
  for i:=0; i< len(array); i ++ {
    for j := i+1 ; j < len(array); j ++ {
      valueToFind := target - array[i] - array[j]
      if _, ok := m[valueToFind]; ok {
        return []int {array[i], array[j], valueToFind}
      }
    }
  }

  return make([]int, 0)
}

func main() {
  arr := []int {5,1,7,4,8,3}
  target := 20
  fmt.Println("three sum = ", threeSum(arr, target))

  arr = []int {12, 3, 4, 1, 6, 9}
  target = 24
  fmt.Println("three sum = ", threeSum(arr, target))

  arr = []int {1, 2, 3, 4, 5}
  target = 9
  fmt.Println("three sum = ", threeSum(arr, target))
}
