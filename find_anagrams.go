package main

import (
  "fmt"
  "sort"
)

func isEqual(slice1, slice2 []byte) bool {
  if len(slice1) != len(slice2) {
    return false
  }

  for idx := 0; idx < len(slice1); idx++ {
    if slice1[idx] != slice2[idx] {
      return false
    }
  }

  return true
}

func findAnagrams(str1, str2 string) int {
  // Basic checks
  if len(str1) == 0 || len(str2) == 0 || len(str1) < len(str2) {
    return 0
  }

  // Need to check if the first string has anagrams of second string
  // Two strings are anagrams if their sorted strings are same.
  str1Byte := []byte(str1)
  str2Byte := []byte(str2)

  // Sort str2Byte
  sort.Slice(str2Byte, func(i int, j int) bool { return str2Byte[i] < str2Byte[j] })

  // Create a sliding window of size len(str2)
  start := 0
  end := 0
  num_anagrams := 0
  var tmpSlice []byte

  for idx:=0; idx < len(str2); idx++ {
    end++
  }

  for end < len(str1Byte) {
    // Check if str1Byte[start:end] is an anagram of str2
    tmpSlice = append([]byte(nil), str1Byte[start:end]...)
    sort.Slice(tmpSlice, func(i int, j int) bool { return tmpSlice[i] < tmpSlice[j] })
    if isEqual(tmpSlice,str2Byte) == true {
      num_anagrams ++
    }
    start ++
    end ++
  }

  return num_anagrams
}

func main() {
  str1 := "foxxofoxf"
  str2 := "fox"

  fmt.Println("Input : ", str1)
  fmt.Println("Second string : ", str2)
  fmt.Println("Number of anagrams = ", findAnagrams(str1, str2))
}

