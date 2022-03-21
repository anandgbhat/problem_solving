package main

import (
  "fmt"
)

const (
  MAXALPHABET int = 26
)

func max(i,j int) int {
  if i > j {
    return i
  }
  return j
}

func partitionLabel(str string) []int {
  // Basic validation
  if len(str) == 0 {
    return make([]int, 0)
  }

  // Create a map for each character. Map contains last occurrence of
  // all characters in the slice
  strSlice := []byte(str)
  m := make(map[string]int, MAXALPHABET)
  for i, _ := range strSlice {
    m[string(strSlice[i])] = i
  }
  fmt.Println(m)
  result := make([]int, 0)

  // Walk through string and keep track of highest index seen so far with
  // current set of characters. Split when i == highest index
  end := 0
  size := 0

  for i:=0; i < len(strSlice); i++ {
    size ++
    end = max(end, m[string(strSlice[i])])
    fmt.Println(end, string(strSlice[i]), size)
    if i == end {
      result = append(result, size)
      size = 0
    }
  }

  return result
}


func main() {
  str :="ababcbacadefegdehijhklij"
  fmt.Println("Partitons of ", str, " at ", partitionLabel(str))
}

