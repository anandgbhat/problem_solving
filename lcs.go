package main

import(
  "fmt"
)

func max(i,j int) int {
  if i>j {
    return i
  }
  return j
}

func lcs(str1 []byte, str2 []byte,  n int, m int) int {
  // Base condition
  if n <= 0 || m <= 0 {
    // No LCS possible. Return 0
    return 0
  }
  if str1[n-1] == str2[m-1] {
    // Last characters match. So the length should be 1 + rest of the string
    return 1+lcs(str1, str2, n-1, m-1)
  } else {
    // Last characters are not same.
    // They may match if last character of either of the string is eliminated
    return max(lcs(str1,str2,n-1,m), lcs(str1,str2,n, m-1))
  }
}

func main() {
  str1 := "abcdgh"
  str2 := "abedfhr"

  fmt.Printf("lcs len of %s and %s = %d\n", str1, str2,
   lcs([]byte(str1), []byte(str2), len(str1), len(str2)))

  str1 = "aaaaabcdhh"
  str2 = "a"
  fmt.Printf("lcs len of %s and %s = %d\n", str1, str2,
   lcs([]byte(str1), []byte(str2), len(str1), len(str2)))

  str1 = "abcdef"
  str2 = ""
  fmt.Printf("lcs len of %s and %s = %d\n", str1, str2,
   lcs([]byte(str1), []byte(str2), len(str1), len(str2)))


}

