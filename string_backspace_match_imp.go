package main

import (
  "fmt"
)

// # represents a backspace
// So ab#c is same as ac
// abc##d is same as ad

// Both strings can contains #s

func sanitizeStr(s string) string {
  // Validate
  if len(s) == 0 {
    return ""
  }
  // Convert the string into rune
  sRune := []rune(s)
  var str []rune
  j := 0
  var validCharFound bool  // If any valid char found, will be set to true

  // # character is 35
  // Walk through the rune slice and get rid of previous characters
  for _, val := range sRune {
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

  return string(str)
}

func isMatch(s1, s2 string) bool {
  return s1 == s2
}


func main() {
  fmt.Println("Match abc##c and ac = ",
    isMatch(sanitizeStr("abc##c"), sanitizeStr("ac")))
  fmt.Println("Match a###b##c and a#####c = ",
    isMatch(sanitizeStr("a###b##c"), sanitizeStr("a#####c")))
}
