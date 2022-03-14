package main

import (
  "fmt"
  "os"
  "strings"
)

type stack []string

func (stk *stack) push(val string) {
  *stk = append(*stk, val)
}

func (stk *stack) pop() string {
  if stk.isEmpty() == true {
    return ""
  }
  lastIdx := len(*stk)-1
  val := (*stk)[lastIdx]
  *stk = (*stk)[:lastIdx]
  return val
}

func (stk *stack) isEmpty() bool {
  return len(*stk) == 0
}

func simplifyPath(path string) string {
  var stk stack
  // Divide components based on delimiter '/'
  fields := strings.Split(path, "/")
  // Parse each field and push the value to stack
  for _, p := range fields {
    fmt.Println(string(p))
    if p == string(".")  || p == string("") {
      // Do nothing
    } else if p == string("..") {
          if stk.isEmpty() != true {
            // Pop the past element
            _ = stk.pop()
          }
    } else {
          stk.push(string(p))
    }
  }


  fmt.Println(stk)
  var output string
  for i:=0; i< len(stk); i++ {
    fmt.Println("stk = ", stk[i])
    output += "/"+stk[i]
  }
  if len(stk) == 0 {
    output = "/"
  }

  return output
}

func main() {
  fmt.Printf("Canonical path of %s is %s\n", os.Args[1], simplifyPath(os.Args[1]))
}
