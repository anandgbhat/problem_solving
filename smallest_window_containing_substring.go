package main

import (
  "fmt"
  "math"
)

func hasAllElements(strMap, subStrMap map[byte]int) bool {
  for k, _ := range subStrMap {
    if _, ok := strMap[byte(k)]; !ok {
      return false
    }
  }
  return true
}

func smallestWindowContainingSubstring(str string, subStr string) string {
  // basic checks
  if len(str) == 0 || len(subStr) == 0 || len(subStr) > len(str) {
    return ""
  }

  minLength := math.MaxInt
  startIdx := 0

  strByte := []byte(str)
  subStrByte := []byte(subStr)

  strMap := make(map[byte]int)
  subStrMap := make(map[byte]int)

  // Populate subStrMap
  for _, v := range subStrByte {
    subStrMap[v] ++
  }

  // Variable sliding window problem
  // Find the substring that has all characters in the subStr
  // Once it is found, try to move the window till we find the minimum length

  left, right := 0 , 0

  for ; right < len(str); right ++ {
    // Add the right char to the map
    strMap[strByte[right]] ++

    for  hasAllElements(strMap, subStrMap) == true {
      // Check and update minLength
      if (right - left +1) < minLength {
        minLength = right - left +1
        startIdx = left
      }
      // Move the window and check if anything matches
      strMap[strByte[left]]--
      if strMap[strByte[left]] == 0 {
        delete(strMap, strByte[left])
      }
      // Move the window
      left ++
    }
  }

  // Construct the string from startIdx and minLength
  strSlice := str[startIdx:startIdx+minLength]
  return string(strSlice)
}

func main() {
  fmt.Println("Smallest window containing abc in aabdec = ",
    smallestWindowContainingSubstring("aabdec", "abc"))
  fmt.Println("Smallest window containing abc in abdbca = ",
    smallestWindowContainingSubstring("abdbca", "abc"))
}


