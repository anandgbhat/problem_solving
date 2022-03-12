package main

import(
  "fmt"
  "os"
)


func max(i,j int) int {
  if i>j {
    return i
  }
  return j
}

func lcs(mem [][]int, str1 []byte, str2 []byte,  n int, m int) int {
  // Base condition
  if mem[n][m] != -1 {
    // Already evaluated. Return stored value
    return mem[n][m]
  }

  if n <= 0 || m <= 0 {
    // No LCS possible. Return 0
    return 0
  }
  if str1[n-1] == str2[m-1] {
    // Last characters match. So the length should be 1 + rest of the string
    mem[n][m] = 1+lcs(mem, str1, str2, n-1, m-1)
    return 1+lcs(mem, str1, str2, n-1, m-1)
  } else {
    // Last characters are not same.
    // They may match if last character of either of the string is eliminated
    mem[n][m] = max(lcs(mem, str1,str2,n-1,m), lcs(mem, str1,str2,n, m-1))
    return max(lcs(mem, str1,str2,n-1,m), lcs(mem, str1,str2,n, m-1))
  }
}

func main() {
  str1 := os.Args[1]
  str2 := os.Args[2]
  // Initialize memoization matrix
  mem := make([][]int, len(str1)+1)
  for i := range mem{
    mem[i] = make([]int, len(str2)+1)
  }

  // Initialize matrix to -1
  for i:=0; i <= len(str1); i++ {
    for j:=0; j <= len(str2); j++ {
      mem[i][j] = -1
    }
  }
  fmt.Printf("lcs len of %s and %s = %d\n", str1, str2,
   lcs(mem, []byte(str1), []byte(str2), len(str1), len(str2)))
}

