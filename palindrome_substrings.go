package main

import (
  "fmt"
)

func palindromeSubstringsCount(str string) int {
  // Convert string to byte slice
  strByte := []byte(str)

  // Calculate palindromes with odd lengths
  // Logic: Use two pointers, left and right both pointing to same element
  // Check if the elements around the middle forms a palindrome.
  // If so, incerement the count.
  result := 0
  for i:=0; i < len(strByte); i++ {
    // Calculate odd sized palindromes
    // Logic: left and right point to same char and they spread to left and right. If char at left is not same as right, break.
    result += checkPalindrome(strByte, i, i)
    // Calculate even sized palindromes
    // Logic: left and right point to adjacent elements. left and right spread in either direction. If char at left is not same as right, break.
    result += checkPalindrome(strByte, i, i+1)
  }

  return result
}

func checkPalindrome(strByte []byte, left,right int) int {
  result := 0
  for right < len(strByte) && left >= 0 && strByte[left] == strByte[right] {
    // Increment the number of palindromes
    result ++
    // increment right and decrement left
    left --
    right ++
  }
  return result

}


func main() {
  fmt.Println("Number of palindromic substrings in abac = ", palindromeSubstringsCount("abac"))
}
