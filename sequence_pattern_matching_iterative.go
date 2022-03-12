package main

import (
  "fmt"
  "os"
)

func sequencePattenMatching(str1,str2 string) bool {
  // Characters in str1 should match a subsequence in str2

  i := 0
  j := 0
  str1Byte := []byte(str1)
  str2Byte := []byte(str2)

  for i < len(str1) && j < len(str2) {
    // Check if str1[i] is present in str2
      for j < len(str2) && str1Byte[i] != str2Byte[j] {
        j++
      }
      if j == len(str2) {
        return false
      }
      // Found
      // Increment i
      i++
  }

  return true
}

func main() {
  fmt.Printf("Sequence pattern matching of %s and %s = %v\n",
    os.Args[1], os.Args[2], sequencePattenMatching(os.Args[1], os.Args[2]))
}
