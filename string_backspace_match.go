package main

import (
  "fmt"
)

// # represents a backspace
// So ab#c is same as ac
// abc##d is same as ad

func isMatch(s1,s2 string) bool {
  // Validation
  if len(s1) == 0 && len(s2) == 0 {
    fmt.Println("Invalid strings")
    return false
  }

  // Convert the string into rune
  s1Rune := []rune(s1)
  var str []rune
  j := 0
  var validCharFound bool  // If any valid char found, will be set to true

  // # character is 35
  // Walk through the rune slice and get rid of previous characters
  for _, val := range s1Rune {
    if val != 35 {
      str = append(str, val)
      validCharFound = true
      j++
    } else {
      if validCharFound != true {
        // No valid characters seen so far
        // Skip the # and move forward
        continue
      }
      if j > 0 {
        j--
      }
      str = str[0:j]
    }
  }

  return string(str) == s2
}

func main() {
  fmt.Println("Match abc##c and ac = ", isMatch("abc##c", "ac"))
  fmt.Println("Match ab#c and abc = ", isMatch("ab#c", "abc"))
  fmt.Println("Match ##ab#c and abc = ", isMatch("##abb#c", "abc"))
  fmt.Println("Match ab## and nullstr = ", isMatch("ab##", ""))
  fmt.Println("Match a###b and b = ", isMatch("a###b", "b"))
}
