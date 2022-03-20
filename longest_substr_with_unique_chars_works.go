package main
import (
  "fmt"
)

func max(i,j int) int {
  if i > j {
    return i
  }
  return j
}

func longestSubstrWithKuniqueChars(str string, k int) int {
  if len(str) == 0 || len(str) < k {
    return 0
  }

  i := 0
  j := 0
  maxLen := 0
  charMap := make(map[byte]int, 32)
  strSlice := []byte(str)
  for j < len(strSlice) {
    // Add the element at j to map
    charMap[strSlice[j]]++
    // If number of unique characters is > k, remove elements at
    // strSlice[i] till map size becomes k. Increment i each time
    for len(charMap) > k {
      charMap[strSlice[i]]--
      if charMap[strSlice[i]] == 0 {
        delete(charMap, strSlice[i])
      }
      i++
    }
    maxLen = max(maxLen, (j-i)+1)
    j++
  }

  return maxLen
}

func main() {
  fmt.Println("maximum len substring with 2 uniwue characters in ", "araaci",
    "is = ", longestSubstrWithKuniqueChars("araaci", 2))
}



