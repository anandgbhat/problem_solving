package main

import (
  "fmt"
)

func longestSubstrOfSameCharAfterKReplacements(str string, rep int) int {
  // Logic:
  // Sliding window technique. Note down number of occurrences of each
  // charcter in a map. Subtract the number of occurrences of the most
  // frequent character from the substring length. If that's less than k,
  // note down the length.
  // Repeat till we get max length.

  left := 0
  right := 1
  charsSeen := make(map[string]int)
  longestSubstrLen := 0
  currentLen := 0
  strByte := []byte(str)
  maxOccurringCharLen := 0

  for left < len(strByte) && right < len(strByte) {
    for i:=left; i <= right; i++ {
      if _, ok := charsSeen[string(strByte[i])]; !ok {
        charsSeen[string(strByte[i])] = 1
      } else {
        charsSeen[string(strByte[i])] ++
      }
      currentLen ++
      // Get maximum element stored in the map
      maxOccurringCharLen = maxLen(charsSeen)
      if currentLen - maxOccurringCharLen  > rep {
        // We overshoot the maximum characters that can be replaced
        // Reset the map, currentLen, reset left etc.
        left ++
      } else {
        if currentLen > longestSubstrLen {
          longestSubstrLen = currentLen
        }
      }
    }
        cleanupMap(charsSeen)
        currentLen = 0
    right ++
  }

  return longestSubstrLen
}

func maxLen(seen map[string]int) int {
  maxValue := 0
  for _, value := range seen {
    if value > maxValue {
      maxValue = value
    }
  }

  return maxValue
}

func cleanupMap(seen map[string]int) {
  for key, _ := range seen {
    delete(seen, key)
  }
}

func main() {
  fmt.Printf("Maximum length of substring where %d chars can be " +
  "replaced = %d in aabccbb\n", 2,
  longestSubstrOfSameCharAfterKReplacements("aabccbb", 2))
  fmt.Printf("Maximum length of substring where %d chars can be " +
  "replaced = %d in abbcb\n", 1,
  longestSubstrOfSameCharAfterKReplacements("abbcb", 1))
  fmt.Printf("Maximum length of substring where %d chars can be " +
  "replaced = %d in abbccde\n", 1,
  longestSubstrOfSameCharAfterKReplacements("abccde", 1))
}
