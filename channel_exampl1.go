package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func display(words chan string) {
  defer wg.Done()
  str := <- words
  fmt.Println(str)
  words <- "received"
}

func main() {
  words := make(chan string)
  wg.Add(1)
  go display(words)
  words <- "Sending encoded messages"
  fmt.Println(<-words)
  close(words)
  wg.Wait()
}

