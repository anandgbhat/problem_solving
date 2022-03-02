package main

import (
  "fmt"
)

func longestSubstr(str string) int {
  if len(str) == 0 {
    return 0
  }

  seen := make(map[byte]bool)

  // convert string into []byte
  byteStr := []byte(str)
  // walk through all characters and save the bytes in map
  // If any byte is already seen, start the count again
  longestLen := 0
  currentCount := 0
  for _, v := range byteStr {
    if _, ok := seen[v]; ok {
      // Byte already seen
      // Reset longestLen
     currentCount = 0
    } else {
      seen[v] = true
      currentCount ++
    }
    if currentCount > longestLen {
      longestLen = currentCount
    }
  }

  return longestLen
}

func main() {
  fmt.Println("Longest substring len in abccbadd = ", longestSubstr("abccbadd"))
  fmt.Println("Longest substring len in accccc = ", longestSubstr("accccc"))
  fmt.Println("Longest substring len in cccc = ", longestSubstr("cccc"))
}


