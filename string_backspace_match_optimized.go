package main

import (
  "fmt"
)

func strMatch(src,dst string) bool {
  // Convert strings into byte slices
  srcByte := []byte(src)
  dstByte := []byte(dst)

  // Logic: Use two pointers - one for each string
  // Both pointers point to last element in the slice
  // If the last element is a '#', decrement the pointer by 1
  // and check if the previous element is also a #.
  p1 := len(srcByte)-1
  p2 := len(dstByte)-1
  var p1dec int
  var p2dec int
  // abc  abcde##
  for p1 >= 0 || p2 >= 0{
    // Skip all '#' characters and note down how many are seen
    for p1 >= 0 && srcByte[p1] == '#' {
      p1dec ++
      p1--
    }
    p1 -= p1dec
    p1dec = 0
    // p1 still has characters and p2 is done, return false
    if p1 >= 0 && p2 < 0 {
      return false
    }
    // Skip all '#' characters in dstByte
    for p2 > 0 && dstByte[p2] == '#' {
      p2dec ++ // Number of # characters seen
      p2 --
    }
    p2 -= p2dec
    p2dec = 0
    if p2 < 0 {
      break
    }
    // p2 still has characters and p1 is done, return false
    if p2 >= 0 && p1 < 0 {
      return false
    }
    if srcByte[p1] != dstByte[p2] {
      return false
    } else {
      p1 --
      p2 --
    }
  }
  return true
}

func main() {
  fmt.Println("Does abc abcde## match? ", strMatch("abc", "abcde##"))
  fmt.Println("Does ###abc abcde## match? ", strMatch("###abc", "abcde##"))
  fmt.Println("Does baaa###abc abcde## match? ", strMatch("baaa###abc", "abcde##"))
  fmt.Println("Does abcde## baaa###abc match? ", strMatch("abcde##", "baaa###abc"))
  fmt.Println("Does abcde## aaa###abc match? ", strMatch("abcde##", "aaa###abc"))
}
