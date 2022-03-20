
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

  // Walk through the list till we hit nil
  for tmp != nil {
    if tmp.child == nil {
      tmp = tmp.next
    } else {
      tail := tmp.child
      for tail.next != nil {
        tail = tail.next
      }

      tail.next = tmp.next
      if tail.next != nil {
        tail.next.prev = tail
      }
      tmp.next = tmp.child
      tmp.next.prev = tmp
      tmp.child = nil
    }
  }
}
