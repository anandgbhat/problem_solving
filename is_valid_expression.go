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

func (stack *Stack) Peak() string {
  if stack.IsEmpty() == true {
    return ""
  }
  return (*stack)[len(*stack)-1]
}

func isValidExp(str string) bool {
  // Basic validations
  if len(str) == 0 {
    return true
  }

  var stack Stack

  // Convert str to []byte
  strSlice := []byte(str)
  for _, val := range strSlice {
    // Only valid expressions include ( followed by ) and
    // they should match in numbers. If not, its an invalid expression.
    if string(val) == "(" {
      stack.Push(string(val)) // Store in stack to match later
    } else if string(val) == ")" {
      // Check if stack already contains matching "("
      // If not, its an invalid expression
      if stack.IsEmpty() == true {
        return false
      }
      top := stack.Pop()
      if top != "(" {
        return false
      }
    }
  }
  // If anything is left in the stack, expression is not valid
  if stack.IsEmpty() != true {
    return false
  }
  return true
}

func main() {
  fmt.Println("Is expression )abc(pqr valid ? ", isValidExp(")abc(pqr"))
  fmt.Println("Is expression (abc(pqr)) valid ? ", isValidExp("(abc(pqr))"))
  fmt.Println("Is expression (a)bc(p)q((r)) valid ? ", isValidExp("(a)bc(p)q((r))"))
  fmt.Println("Is expression (a)bc(p)q((r) valid ? ", isValidExp("(a)bc(p)q((r)"))
  fmt.Println("Is expression () valid ? ", isValidExp("()"))
  fmt.Println("Is expression )( valid ? ", isValidExp(")("))
}
