package main

import (
  "fmt"
)

type treeNode struct {
  val string
  left *treeNode
  right *treeNode
}

type Stack []treeNode

func (stack *Stack) IsEmpty() bool {
  return len(*stack) == 0
}

func (stack *Stack) Push (node *treeNode) {
  *stack = append(*stack, *node)
}

func (stack *Stack) Pop() *treeNode {
  if stack.IsEmpty() == true {
    return nil
  }
  // Return last element and reinitalize the stack to
  // len -1
  index := len(*stack) -1
  element := (*stack)[index]
  *stack = (*stack)[:index]
  return &element
}


func (root *treeNode) preOrderTraversalIterative() {
  var stack Stack
  // Push the root to stack
  // Pop it out and check root.right and root.left
  // If they are non-NULL, push to stack in same order
  // Repeat
  if root != nil {
    stack.Push(root)
  }
  node := stack.Pop()
  fmt.Println(node.val)

  for  node != nil {
    if node.right != nil {
      stack.Push(node.right)
    }
    if node.left != nil {
      stack.Push(node.left)
    }
    node = stack.Pop()
    if node != nil {
      fmt.Println(node.val)
    }
  }
}

func main() {
  root := new(treeNode)
  root.val = "a"

  root.right = new(treeNode)
  root.right.val = "c"

  root.left = new(treeNode)
  root.left.val = "b"

  root.left.left = new(treeNode)
  root.left.left.val = "d"

  root.left.right = new(treeNode)
  root.left.right.val = "e"

  root.right.right = new(treeNode)
  root.right.right.val = "f"

  (*treeNode).preOrderTraversalIterative(root)
}
