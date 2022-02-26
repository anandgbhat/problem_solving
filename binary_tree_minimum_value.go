package main

import (
  "fmt"
  "math"
  "sort"
)

type treeNode struct {
  val int
  left *treeNode
  right *treeNode
}


func min(a,b,c int) int {
  var arr []int

  arr = append(arr, a)
  arr = append(arr, b)
  arr = append(arr, c)
  sort.Ints(arr)
  return arr[0]
}

var minimumValue int
var leftMin, rightMin int

func (root *treeNode) minimumValue() int {
  if root == nil {
    return math.MaxInt
  }

  leftMin = (*treeNode).minimumValue(root.left)
  rightMin= (*treeNode).minimumValue(root.right)

  return  min(root.val, leftMin, rightMin)
}

func main() {

  minimumValue = math.MaxInt

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
  fmt.Println("Minimum value = ", (*treeNode).minimumValue(root))

}
