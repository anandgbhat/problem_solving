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

func lcs(mem [][]int, str1 []byte, str2 []byte,  m int, n int) int {

  for i:=1; i<m+1; i++ {
    for j:=1; j<n+1; j++ {
      if str1[i-1] == str2[j-1] {
        mem[i][j] = 1+mem[i-1][j-1]
      } else {
        // Last characters are not same.
        // They may match if last character of either of the
        // string is eliminated
        mem[i][j] =  max(mem[i-1][j], mem[i][j-1])
      }
    }
  }

  return mem[m][n]
}

func main() {
  str1 := os.Args[1]
  str2 := os.Args[2]
  // Initialize memoization matrix
  mem := make([][]int, len(str1)+1)
  for i := range mem{
    mem[i] = make([]int, len(str2)+1)
  }

  // Initialize first two and first column to 0
  for i:=0; i <= len(str1); i++ {
    for j:=0; j <= len(str2); j++ {
      if i==0 {
        mem[i][j] = 0
      } else if j== 0 {
        mem[i][j] = 0
      }
    }
  }
  fmt.Printf("lcs len of %s and %s = %d\n", str1, str2,
   lcs(mem, []byte(str1), []byte(str2), len(str1), len(str2)))
}

