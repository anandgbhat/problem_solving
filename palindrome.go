package main

import (
  "fmt"
)

func isPalindrome(str string) bool {
  strSlice := []byte(str)
  // Two pointer method
  left := 0
  right := len(strSlice) - 1
  for left != right {
    if strSlice[left] != strSlice[right] {
      return false
    }
    left ++
    right --
  }
  if left == right {
    return true
  }
  return false
}

func main() {
  fmt.Println("Is kayak palindrome? ", isPalindrome("kayak"))
  fmt.Println("Is car palindrome? ", isPalindrome("car"))
  fmt.Println("Is Aka Aka  palindrome? ", isPalindrome("AKA AKA"))
}
