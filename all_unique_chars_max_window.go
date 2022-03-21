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

func maxAllUniqueChars(str string) int {
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
    winSize := (end - start) +1
    // If map size < winSize, there is at least one character that's
    // duplicate. Eliminate that character.
    if len(m) <  winSize {
      winSize = (end - start) +1
      // Keep reducing element from strSlice[start] till the condition is false
      m[string(strSlice[start])]--
      if m[string(strSlice[start])] == 0 {
        delete(m, string(strSlice[start]))
      }
      fmt.Println(len(m), winSize, m)
      start ++
    }
    max = maximum(max, (end-start)+1)
    end++
  }

  return max
}

func main() {
  fmt.Println("Maximum unique char subarray len in pwwkenw is = ",
    maxAllUniqueChars("pwwkenw"))
}



