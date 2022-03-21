package main

import (
  "fmt"
)

// Logic:
// Have head and tail pointers pointing to start and end of doubly linked list
// Each time an cache entry is put, add the entry at the end
// If cache is full, remove the entry at the head and add the new value to tail
// For every cache hit, remove the entry and add it back to head

type Node struct {
  value int
  key int
  next *Node
  prev *Node
}

var head *Node
var tail *Node
var m map[int]*Node
var hashSize int

const (
  cacheCapacity int = 128
)


func get(key int) int {
  if val, ok := m[key]; !ok {
    fmt.Println("Key ", key, "does not exists.")
    return -1
  }
  return val
}

func putKey(key, value int) {
  // Check if the value exists in the map
  if node, ok := m[key]; ok {
      // Value already exists. Remove it from the
      // current list and add it to the beginning
      // after updating the value.
      removeNode(node)
      node.value = value
      addNode(node)
    } else {
      if hashSize == cacheCapacity {
        // Remove node from the head
        delete(m, head->next->key)
        removeNode(head->next)
      }
      node := new(Node)
      node.key = key
      node.value = value
      addNode(node)
      m[key] = node
      hashSize ++
   }
}

func addNode(node *Node) {
  node.next = head.next
  node.prev = head
  head.next.prev = node
  head.next = node
}

func removeNode(node *Node) {
  prevNode := node.prev
  node.next.prev = node.prev
  prevNode.next = node.next
}
