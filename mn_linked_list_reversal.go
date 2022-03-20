package main

import (
  "fmt"
)

func mnReversal(head *node, m,n int) *node {
  currentPos := 1
  current_node := head
  start := head

  // Advance currentPos to just before m
  for currentPos < m {
    start = currentNode
    currentNode = currentNode.next
    currentPos++
  }

  newList := node*(nil)
  tail := currentNode

  for currentPos >= m && currentPos <= n {
    next := currentNode.next
    currentNode.next = newList
    newList = currentNode
    currentNode = next
    currentPos++
  }

  start.next = newList
  tail.next = currentNode

  if m > 1 {
    return head
  }

  return newList
}

