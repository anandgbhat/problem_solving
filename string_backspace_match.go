package main

import (
  "fmt"
)

// # represents a backspace
// So ab#c is same as ac
// abc##d is same as ad

func isMatch(s1,s2 string) bool {
  // Convert the string into rune
  s1Rune := []rune(s1)
  var str []rune
  j := 0

  // # character is 35
  // Walk through the rune slice and get rid of previous characters
  for _, val := range s1Rune {
    if val != 35 {
      str = append(str, val)
      j++
    } else {
      j--
      str = str[0:j]
    }
  }

  return string(str) == s2
}

func main() {
  fmt.Println("Match abc##c and ac = ", isMatch("abc##c", "ac"))
  fmt.Println("Match ab#c and abc = ", isMatch("ab#c", "abc"))
}
