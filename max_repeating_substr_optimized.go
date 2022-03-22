package main

import (
  "fmt"
  "strings"
)

// If str1 is ababc and str2 is ab, then answer is 2

func maxRepeatingSubstr(str1, str2 string) int {
  // Basic checks
  if len(str1) == 0 || len(str2) == 0 || len(str1) < len(str2) {
    return 0
  }
  var i int
  var str3 string
  for i = 0; i < len(str1); i++ {
    // Crude idea.
    // Append str2 to str2 and check if the pattern is found
    str3 = str3 + str2
    fmt.Println(str3)
    if strings.Contains(str1, str3) == false {
      fmt.Println("False ")
      break
    }
  }

  return i
}

func main() {
  fmt.Println("Max repeating characters of aaa and a = ",
    maxRepeatingSubstr("aaa", "a"))
}



