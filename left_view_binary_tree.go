package main

import (
  "fmt"
)

type treeNode struct {
  val int
  left *treeNode
  right *treeNode
}


var max_level int

func leftView(root *treeNode, level int) {
  if root == nil {
    return
  }
  // Check if the node is at the same level as max_level seen
  // so far. If so, skip. Otherwise print the value and update
  // max_level seen so far.
  // Since this is preorder tree traversal, algo prints left subtree
  if max_level < level {
    fmt.Println(root.val)
    max_level = level
  }

  leftView(root.left, level+1)
  leftView(root.right, level+1)
}

func main() {
  root := new(treeNode)
  root.val = 1
  root.left = new(treeNode)
  root.left.val = 2
  root.right = new(treeNode)
  root.right.val = 3
  root.right.left = new(treeNode)
  root.right.left.val = 4
  root.right.right = new(treeNode)
  root.right.right.val = 5
  root.right.left.right = new(treeNode)
  root.right.left.right.val = 6

  leftView(root, 1)
}

