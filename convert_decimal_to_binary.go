package main

import (
  "fmt"
  "strconv"
)

var result []byte
// Divide the number by 2 repeatedly and note down the remainder
// in a byte slice.
func convertToBinary(num int, result []byte) []byte {
  if num == 0 {
    return result
  }

  result = append([]byte(strconv.Itoa(num % 2)), result...)
  return convertToBinary(num/2, result)
}

func main() {
  fmt.Println("233 in binary =", string(convertToBinary(233, result)))
}

