
package main

import (
  "fmt"
  "os"
)

func removeSubstr(s1, s2 string) string {
  i := 0
  j := 0

  s1Byte := []byte(s1)
  s2Byte := []byte(s2)

  for {
    if i > len(s1) -1  || j > len(s2) -1 {
      break
    }
    if s1Byte[i] == s2Byte[j] {
      s1Byte = append(s1Byte[:i], s1Byte[i+1:]...)
      j++
      continue
    }
    i++
  }

  return string(s1Byte)
}


func main() {
  fmt.Printf("Removing %s from %s = %s\n",
    os.Args[1], os.Args[2], removeSubstr(os.Args[1], os.Args[2]))
}
