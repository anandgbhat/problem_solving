package main

import (
  "fmt"
)

type Queue []treeNode

func (queue *Queue) IsEmpty() bool {
  return len(*queue) == 0
}

func (queue *Queue) Enqueue (node *treeNode) {
  *queue = append(*queue, *node)
}

func (queue *Queue) Dequeue() *treeNode {
  if queue.IsEmpty() == true {
    return nil
  }
  // Return the first element and reinitialize the slice
  // to [1:]
  element := (*queue)[0]
  *queue = (*queue)[1:]
  return &element
}

type treeNode struct {
  val string
  left *treeNode
  right *treeNode
}

var maxWidth int

func (root *treeNode) maxTreeWidthBFS() int {
  // basic validation
  if root == nil {
    return 0
  }

  var queue Queue
  queue.Enqueue(root)
  maxWidth = 1
  countAtLevel := len(queue)
  var node *treeNode

  for queue.IsEmpty() != true {
    countAtLevel = len(queue)
    if countAtLevel > maxWidth {
      maxWidth = countAtLevel
    }

    // Dequeue all elements at this level
    for countAtLevel != 0 {
      node = queue.Dequeue()
      countAtLevel --
      if node.left != nil {
        queue.Enqueue(node.left)
      }
      if node.right != nil {
        queue.Enqueue(node.right)
      }
    }
  }

  return maxWidth
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

  fmt.Println("Maximum width of the tree = ",
    (*treeNode).maxTreeWidthBFS(root))
}
