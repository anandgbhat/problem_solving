package main

import (
  "fmt"
)

type treeNode struct {
  val int
  left *treeNode
  right *treeNode
}


var result []int

func leftView(root *treeNode, level int) {
  if root == nil {
    return
  }
  // If the result array has right most element of last level
  // append this value. If not, then the right at this level
  // is already present in the result.
  if len(result) < level {
    result = append(result, root.val)
  }
  // Travrse right subtree first this time
  leftView(root.right, level+1)
  leftView(root.left, level+1)
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
  fmt.Println(result)
}

