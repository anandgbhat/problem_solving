package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

// NOTE: In other languages, read the entire file and parse the content to separate out lines with '\n' as delimiter.
// Reverse the array
// Alternatively, push each line onto stack and pop() and print each line.
// In go, its pretty simple due to presence of strings.Split()


func reverseFile(fileName string) []string {
  if len(fileName) == 0 {
    fmt.Println("Invalid filename")
    return make([]string, 0)
  }

  // Read the file
  data, err := ioutil.ReadFile(fileName)
  if err != nil {
    fmt.Printf("Unable to read file. Error = %v\n", err)
    return make([]string, 0)
  }

  // Split lines
  lines := strings.Split(string(data), "\n")
  output := make([]string, len(lines))
  // Walk through the slice in reverse order and append to a new slice
  for i:=len(lines)-1; i >= 0; i-- {
    output = append(output, lines[i])
    output = append(output, "\n")
  }

  return output
}

func main() {
  fmt.Println("Reversed file = ", reverseFile("/Users/anand.bhat/x.go"))
}
