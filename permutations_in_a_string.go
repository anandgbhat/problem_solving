package main

import (
  "fmt"
  "reflect"
)

func checkPermutation(strMap map[string]int, patternMap map[string]int) bool {
  return reflect.DeepEqual(strMap, patternMap)
}


func isPermutationPresent(str string, pattern string) bool {
  // Basic validations
  if len(str) == 0 || len(pattern) == 0 || len(str) < len(pattern) {
    return false
  }

  // Create a sliding window and figure out whether number of each characters in
  // the sliding window same as the one in pattern.

  patternMap := make(map[string]int) // For pattern
  strMap := make(map[string]int) // For input string

  // Populate the patternMap
  strByte := []byte(str)
  patternByte := []byte(pattern)
  for _,v := range patternByte {
    patternMap[string(v)] ++
  }

  // Create a sliding window of size len(pattern) and check if char map
  // in any window has same chars in patternMap. If so, return true.
  start := 0
  end := len(pattern) -1
  // Create initial strMap
  for i:=0; i < len(pattern) ; i++ {
    strMap[string(strByte[i])]++
  }

  for end < len(str) {
    if checkPermutation(strMap, patternMap) == true {
      return true
    }
    // Slide the window right
    end ++
    // Remove the character at start from the map
    strMap[string(strByte[start])]--
    if strMap[string(strByte[start])] == 0 {
      delete(strMap, string(strByte[start]))
    }
    // Add the char at end into the map
    if end < len(str) {
      strMap[string(strByte[end])]++
    }
    start ++
  }

  return false
}

func main() {
  fmt.Println(isPermutationPresent("oidbcaf", "abc"))
  fmt.Println(isPermutationPresent("", "abc"))
  fmt.Println(isPermutationPresent("a", "abc"))
  fmt.Println(isPermutationPresent("a", ""))
  fmt.Println(isPermutationPresent("odicf", "dc"))
  fmt.Println(isPermutationPresent("abcd", "cdb"))
}
