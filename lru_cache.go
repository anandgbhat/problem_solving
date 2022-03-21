// Logic: Create a linked list node and assign the actual value in it
// Create a hashmap key and point to the linked list node address
// Insert the node at the head of the doubly linked list
// When there is a need to remove the node for eviction, simply remove
// the last node.
// If there is a need to update the existing key, remove the node from the
// linked list, update its value and readd the node.

// For easier removal, linked list node will store both key and value
// whereas hashmap will store only the node pointer.


package main

import (
  "fmt"
)

const (
  cacheCapacity int = 10
)

type Node struct {
  value int
  key int
  prev *Node
  next *Node
}

var m map[int]*Node
var head *Node
var tail *Node
var hashSize int

func get(key int) int {
  if val, ok := m[key]; ok {
    // val points to a node . Return its value
    return val.value
  }
  return -1
}

func put(key, val int) {
  // If key is already found, reinsert the node at the head and update the
  // value in the linked list
  if node, ok := m[key]; ok {
    removeNode(node)
    node.value = val
    addNode(node)
  } else {
    // value is not found
    // Check if we reached the cache capacity
    // If so, remove the last node
    if hashSize == cacheCapacity {
      delete(m, tail.prev.key)
      removeNode(tail.prev)
    }
    // Add the new node
    node := new(Node)
    node.value = val
    node.key = key
    addNode(node) // Add the node at the beginning of the linked list
    m[key] = node  // Add the node address to hashmap
    hashSize ++ // Increement the number of instances in the hashmap
  }
}

// addNode adds to the head
func addNode(node *Node) {
  // Add the node to the head
  node.next = head.next
  node.prev = head
  head.next.prev = node
  head.next = node
}

// Remove any given node
// Point previous node's next to node's next
// Point node's next's prev to node's prev
// Point previous node's next to node's next
// node gets garbage collected
func removeNode(node *Node) {
  // Remove the node
  prevNode := node.prev
  node.next.prev = node.prev
  prevNode.next = node.next
}
