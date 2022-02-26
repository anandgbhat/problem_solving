package main

import (
  "fmt"
)

func reverseStr(str []byte) []byte {
  if len(str) == 0 {
    return []byte("")
  }

  return append(reverseStr(str[1:]), str[0])
}

func main() {
  stringToReverse := "hello world"
  strSlice := []byte(stringToReverse)
  fmt.Println("Reversed string ", string(reverseStr(strSlice)))
}
