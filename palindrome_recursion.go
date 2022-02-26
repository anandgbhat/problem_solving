package main

import (
  "fmt"
)

func isPalindromeRecursive(str []byte) bool {
  // base case
  if len(str) == 0 || len(str) == 1 {
    return true
  }

  if str[0] == str[len(str)-1] {
    // shrink the strink by 1 character both sides
    return isPalindromeRecursive(str[1:len(str)-1])
  }

  return false
}

func main() {
  fmt.Println("Is racecar a palindrome? ", isPalindromeRecursive([]byte("racecar")))
  fmt.Println("Is abcd a palindrome? ", isPalindromeRecursive([]byte("abcd")))
  fmt.Println("Is nullstr palindrome? ", isPalindromeRecursive([]byte("")))
  fmt.Println("Is abcba palindrome? ", isPalindromeRecursive([]byte("abcba")))
}

