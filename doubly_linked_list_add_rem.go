
type Node struct {
  value int
  prev *Node
  next *Node
}

var head *Node

func addNodeAtHead(node *Node) {
  node.next = head.next
  node.next.prev = node
  node.prev = head
  head.next = node
}

func addNodeLast(node *Node) {
  // Walk till last node in the list
  tmp := *Node(nil)
  tmp = head
  for tmp.next != nil {
    tmp = tmp.next
  }
  // At the last node
  // Insert the node at the end
  tmp.next = node
  node.prev = tmp
  tmp.next = nil
}

func addNodeAt(node *Node, val int) {
  // Insert node after val number of nodes
  tmp := head
  for i:=0; i< val; i++ {
      if tmp.next != nil {
        tmp = tmp.next
      } else {
        fmt.Printf("Invalid offset within the list. Exiting...")
        return
      }
  }
  // tmp points to offset of val from head
  // Add the new node
  node.next = tmp.next
  node.next.prev = tmp
  tmp.next = node
  node.prev = tmp
}

func removeNode(node *Node, val int) {
  // Remove node after val number of nodes
  for i:=0; i < val; i++ {
    if tmp.next != nil {
      tmp = tmp.next
    } else {
      return
    }
  }

  // tmp pointing at the node to be removed
  cur = tmp.prev
  cur.next = tmp.next
  tmp.next.prev = cur
  // Can't delete the object in golang as its garbage collected.
  // Reset pointers instead.
  tmp.next = nil
  tmp.prev = nil
  tmp = nil

}





