package main

import (
  "fmt"
  "os"
)

func maxRepeatingSubstrTimes(str string, substr string) int {
  if len(str) == 0 || len(substr) == 0 {
    return 0
  }

  // Convert str and substr into []byte
  strByte := []byte(str)
  substrByte := []byte(substr)
  i := 0
  count := 0

  for i < len(strByte) {
    // Walk till we find the first character of substr
    for  i < len(strByte)  && strByte[i] != substrByte[0] {
      i++
    }
    // Reached end of str but substr is not found
    if i == len(strByte) {
      return count
    }
    // First char of substr matches with str[i]
    // Check if rest of the characters match
    j := 0
    for j < len(substrByte) {
      if strByte[i] == substrByte[j] {
        i++
        j++
      } else {
        break
      }
    }
    if j == len(substrByte) {
      count ++
    }
    j=0
  }
  return count
}

func main() {
  fmt.Printf("How many time %s occurs in %s = %d\n",
    os.Args[1], os.Args[2], maxRepeatingSubstrTimes(os.Args[1], os.Args[2]))
}
