package main

import (
  "fmt"
  "sort"
)

func removeDupliates(nums []int) int {
  if len(nums) == 0 {
    return 0
  }

  m := make(map[int]int)
  for _, val := range nums {
    m[val] = val
  }

  // Walk through the map and store the keys in a slice
  var sortedArr []int
  for key, _ := range m {
    sortedArr = append(sortedArr, key)
  }
  sort.Ints(sortedArr)
  return len(sortedArr)
}

func main() {
  nums := []int {1, 1, 2, 2, 2, 3, 3, 4, 5}
  fmt.Println("Length of unique array = ", removeDupliates(nums))
}
