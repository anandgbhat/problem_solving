package main

import (
  "fmt"
)

type Queue []string

func (queue *Queue) IsEmpty() bool {
  return len(*queue) == 0
}

func (queue *Queue) Enqueue (str string) {
  *queue = append(*queue, str)
}

func (queue *Queue) Dequeue() string {
  if queue.IsEmpty() == true {
    return ""
  }
  // Return the first element and reinitialize the slice
  // to [1:]
  element := (*queue)[0]
  *queue = (*queue)[1:]
  return element
}

func main() {
  var queue Queue

  queue.Enqueue("This")
  queue.Enqueue("is")
  queue.Enqueue("very")
  queue.Enqueue("good!")

  for len(queue) > 0 {
    x := queue.Dequeue()
    if x != "" {
      fmt.Println(x)
    }
  }
}

