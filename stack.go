package main

import (
  "fmt"
)

type Stack []string

func (stack *Stack) IsEmpty() bool {
  return len(*stack) == 0
}

func (stack *Stack) Push (str string) {
  *stack = append(*stack, str)
}

func (stack *Stack) Pop() string {
  if stack.IsEmpty() == true {
    return ""
  }
  // Return last element and reinitalize the stack to
  // len -1
  index := len(*stack) -1
  element := (*stack)[index]
  *stack = (*stack)[:index]
  return element
}

func main() {
  var stack Stack

  stack.Push("This")
  stack.Push("is")
  stack.Push("very")
  stack.Push("good!")

  for len(stack) > 0 {
    x := stack.Pop()
    if x != "" {
      fmt.Println(x)
    }
  }
}

