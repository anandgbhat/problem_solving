package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func display(words chan []string) {
  defer wg.Done()
  str := <- words
  fmt.Println(str)
  words <- []string {"received"}
}

func main() {
  words := make(chan []string)
  wg.Add(1)
  go display(words)
  s := []string {"sending", "encoded", "messages"}
  words <- s
  fmt.Println(<-words)
  close(words)
  wg.Wait()
}

