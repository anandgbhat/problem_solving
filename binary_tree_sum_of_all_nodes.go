package main

import (
  "fmt"
)

type treeNode struct {
  val int
  left *treeNode
  right *treeNode
}

var total int

func (root *treeNode) sumOfAllValues() int {
  if root == nil {
    return 0
  }
  return root.val + (*treeNode).sumOfAllValues(root.left) +
    (*treeNode).sumOfAllValues(root.right)
}

func main() {
  root := new(treeNode)
  root.val = 3

  root.right = new(treeNode)
  root.right.val = 4

  root.left = new(treeNode)
  root.left.val = 11

  root.left.left = new(treeNode)
  root.left.left.val = 4

  root.left.right = new(treeNode)
  root.left.right.val = 2

  root.right.right = new(treeNode)
  root.right.right.val = 1

  fmt.Println("sum = ", (*treeNode).sumOfAllValues(root))
}
