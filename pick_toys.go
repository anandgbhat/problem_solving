package main

import (
  "fmt"
)

var m map[string]int

func maximum(i,j int) int {
  if i > j {
    return i
  }
  return j
}

func pickToys(str string, k int) int {
  // Convert string to []byte
  strSlice := []byte(str)
  m = make(map[string]int, 10)
  // Create a variable sliding window
  start := 0
  end := 0
  max := 0

  for end < len(strSlice) {
    m[string(strSlice[end])]++
    // Check if m length is greater than current window size
    // If map size < winSize, there is at least one toy that appears twice.

    if len(m) >  k {
      // Keep reducing element from strSlice[start] till the condition is false
      m[string(strSlice[start])]--
      if m[string(strSlice[start])] == 0 {
        delete(m, string(strSlice[start]))
      }
      fmt.Println(m)
      start ++
    }
    max = maximum(max, (end-start)+1)
    end++
  }

  return max
}

func main() {
  fmt.Println("Maximum number of of toys of two types in pwwwwwkew is = ",
    pickToys("pwwwwwkenw", 2))
}



