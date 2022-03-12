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

var output []byte

func reverse(str string) string {
  result :=make([]byte, len(str))

  strByte := []byte(str)
  j:=0
  for i:=len(strByte)-1; i>=0; i-- {
    result[j] = strByte[i]
    j++
  }

  return string(result)
}


func lcs(mem [][]int, str1 []byte, str2 []byte,  m int, n int) []byte {

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


  // Walk through matrix and check how we arrived at the value
  i := m
  j := n
  for i > 0 && j > 0 {
    if str1[i-1] == str2[j-1] {
      // Values are same
      // So the value is got by 1+mem[i-1][j-1]
      // Note down the currrent value and reset i and j
      output = append(output, str1[i-1])
      i --
      j --
    } else {
      // Values are not equal.
      // So the value is got by max(mem[m-1][n], mem[m][n-1])
      // Check which of them is bigger and traverse there
      if mem[i-1][j] > mem[i][j-1] {
        i--
      } else {
        j--
      }
    }
  }

  return output
}

func main() {
  str1 := os.Args[1]
  str2 := reverse(str1)
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

  lps := reverse(string(lcs(mem, []byte(str1), []byte(str2), len(str1), len(str2))))
  fmt.Printf("minimum number of deletions of %s and %s to make palindrome = %d\n", str1, str2,
    len(str1)-len(lps))
}

