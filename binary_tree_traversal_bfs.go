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
  value string
  left *treeNode
  right *treeNode
}

func (root *treeNode) bfsTraversalIterative() {
  var queue Queue
  // Exit condition
  if root == nil {
    return
  }
  queue.Enqueue(root)
  fmt.Println(root.value)
  node := queue.Dequeue()

  // Walk through the tree and push root.left and root.right to the queue
  // Dequeue one element and push its left and right children
  // Repeat
  for node != nil {
    if node.left != nil {
       queue.Enqueue(node.left)
    }
    if node.right != nil {
       queue.Enqueue(node.right)
    }
    node = queue.Dequeue()
    if node != nil {
      fmt.Println(node.value)
    }
  }
}

func main() {
  root := new(treeNode)
  root.value = "a"

  root.right = new(treeNode)
  root.right.value = "c"

  root.left = new(treeNode)
  root.left.value = "b"

  root.left.left = new(treeNode)
  root.left.left.value = "d"

  root.left.right = new(treeNode)
  root.left.right.value = "e"

  root.right.right = new(treeNode)
  root.right.right.value = "f"

  (*treeNode).bfsTraversalIterative(root)
}
