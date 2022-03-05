package main

import (
  "fmt"
)

// Logic: Sliding window along with keeping track of number of
// unique characters seen in a map
// k represents number of unique characters in the string

func longestSubstring(str string, k int) int {
  charsSeen := make(map[string]int) // Counts number of occurrences of each char
  left, right := 0, 1
  var numCharsSeen int
  maxLength := 0

  // Convert string to byte slice
  strByte := []byte(str)
  for left < len(strByte) && right < len(strByte) {
    // Check if the characters between left and right are already seen in the
    // map. If so, increment the counter.
    // If not, add the char to map and increment numCharsSeen
    for i:= left; i <= right; i++ {
      if _, ok := charsSeen[string(strByte[i])]; !ok {
        numCharsSeen ++
      }
//      fmt.Println("Adding ", string(strByte[i]) )
      charsSeen[string(strByte[i])] ++
    }
   // fmt.Println(str, maxLength, left, right, numCharsSeen, charsSeen)
    if numCharsSeen > k {
      // We have already hit more than k unique characters.
      // Pop out characters from left till we have k unique characters.
      // Keep moving left till we hit numCharsSeen <= k
      for left < right {
        charsSeen[string(strByte[left])] --
        if charsSeen[string(strByte[left])] == 0 {
          // Remove the key from map and reduce numCharsSeen
          delete(charsSeen, string(strByte[left]))
          numCharsSeen --
        }
 //       fmt.Println("Incrementing left")
        left ++
        if numCharsSeen == k {
          // We have required number of unique characters
          break
        }
      }
    }
  //  fmt.Println("Incrementing right")
    // Slide to right
    right ++
    // Cleanup the map
    for key, _ := range charsSeen {
      delete(charsSeen, key)
    }
    current_len := right - left
    if maxLength < current_len {
      maxLength = current_len
    }
    // Clenaup numCharsSeen as it will get reinitialized with the map
    numCharsSeen = 0
  }

  return maxLength
}

func main() {
  fmt.Println("Maximum length of substring with 2 unique characters in "+
    " araaci = ", longestSubstring("araaci", 2))
  fmt.Println("Maximum length of substring with 3 unique characters in "+
    " araaci = ", longestSubstring("araaci", 3))

  fmt.Println("Maximum length of substring with 4 unique characters in "+
    " araacjki = ", longestSubstring("araacjki", 4))
}




