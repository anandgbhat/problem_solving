package main

import (
  "fmt"
  "math"
)

type treeNode struct {
  val int
  left *treeNode
  right *treeNode
}

func max(i,j int) int {
  if i > j {
    return i
  }
  return j
}


var maximumPath int
var leftMaxPath int
var rightMaxPath int

func (root *treeNode) maxPath() int {
  // If a node is not present, return its value as -infinity
  if root == nil {
    return math.MinInt
  }
  // If left and right are nil, path of the subtree is the value of
  // the node itself.
  if root.left == nil && root.right == nil {
    return root.val
  }

  return  root.val + max((*treeNode).maxPath(root.left), (*treeNode).maxPath(root.right))
}

func main() {

  root := new(treeNode)
  root.val = 5

  root.right = new(treeNode)
  root.right.val = 3

  root.left = new(treeNode)
  root.left.val = 11

  root.left.left = new(treeNode)
  root.left.left.val = 4

  root.left.right = new(treeNode)
  root.left.right.val = 2

  root.right.right = new(treeNode)
  root.right.right.val = 1
  fmt.Println("Maximum path value = ", (*treeNode).maxPath(root))

}
