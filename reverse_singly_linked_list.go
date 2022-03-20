package main

import (
  "fmt"
)

type node struct {
  val int
  node *next
}

func reverseList(head *node) *node {
  // Base condition
  if head == nil || head->next == nil {
    return head
  }
  prev := nil
  current := head

  for current != nil {
    next := current.next
    current.next = prev
    prev = current
    current = next
  }

  return prev
}


