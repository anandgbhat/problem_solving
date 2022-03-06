package main

import (
  "fmt"
)

func logestDistinctSubstr(str string) int {
  // Base case
  if len(str) <= 1 {
    return 1
  }
  // Logic: sliding window to find all substrings.
  // Store chars seen in map and break if the char is seen already
  // Reset left and increement right until we hit end of string or
  // duplicate character.
  strByte := []byte(str)
  left := 0
  right := 1
  charsSeen := make(map[string]bool)
  longestLen := 0
  currentLen := 0

  for left < len(strByte) && right < len(strByte) {
    for i:= left; i <= right; i++ {
      if _, ok := charsSeen[string(strByte[i])]; !ok {
        currentLen ++
        charsSeen[string(strByte[i])] = true
        if longestLen < currentLen {
          longestLen = currentLen
        }
      } else {
        // Repeat char
        // Cleanup map and reset currentLen
        for key, _ := range charsSeen {
          delete(charsSeen, key)
        }
        currentLen = 0
        left = i
        break
      }
    }
    right ++
  }


  return longestLen
}

func main() {
  fmt.Println("Longest distint substr len in araaci = ",
    logestDistinctSubstr("araaci"))
  fmt.Println("Longest distint substr len in jjjj = ",
    logestDistinctSubstr("jjjj"))
}
