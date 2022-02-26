package main

import (
  "fmt"
  "sort"
)

type treeNode struct {
  val int
  left *treeNode
  right *treeNode
}


func max(a,b,c int) int {
  var arr []int

  arr = append(arr, a)
  arr = append(arr, b)
  arr = append(arr, c)
  sort.Ints(arr)
  return arr[2]
}

var maximumValue int
var leftMax, rightMax int

func (root *treeNode) maximumValue() int {
  if root == nil {
    return 0
  }

  if root.left != nil {
    leftMax = (*treeNode).maximumValue(root.left)
  }
  if root.right != nil {
    rightMax= (*treeNode).maximumValue(root.right)
  }
  return  max(root.val, leftMax, rightMax)
}

func main() {

  maximumValue = 0

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
  fmt.Println("Maximum value = ", (*treeNode).maximumValue(root))

}
