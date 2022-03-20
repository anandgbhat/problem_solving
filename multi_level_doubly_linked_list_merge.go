
type DoublyLinkedList struct {
  value int
  next *DoublyLinkedList
  prev *DoublyLinkedList
  child *DoublyLinkedList
}

func flatten(head *DoublyLinkedList) *DoublyLinkedList {
  if head == nil {
    return head
  }

  tmp := head

  
  for tmp != nil {
    // Advance the list if there are no child pointers in the node
    if tmp.child == nil {
      tmp = tmp.next
    } else {
      // Child pointer. Flatten the child list as is.
      // Ignore any child pointers in the child list as it will get processed in
      // the next iteration of main loop
      tail := tmp.child
      // Traverse the child list till end
      for tail.next != nil {
        tail = tail.next
      }

      // Adjust all pointers to add the child list to main list
      tail.next = tmp.next
      // Handle the case wherein last node has child pinter
      if tail.next != nil {
        tail.next.prev = tail
      }
      tmp.next = tmp.child
      tmp.next.prev = tmp
      tmp.child = nil
    }
  }
}
