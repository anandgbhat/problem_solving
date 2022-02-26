package main

import (
  "fmt"
)

type treeNode struct {
  val string
  left *treeNode
  right *treeNode
}

func (root *treeNode) contains(str string) bool {
  if root == nil {
    return false
  }
  if root.val == str {
    return true
  }

  return (*treeNode).contains(root.left, str) || (*treeNode).contains(root.right, str)
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

  if (*treeNode).contains(root, "e") {
    fmt.Println("tree contains e")
  }
}
