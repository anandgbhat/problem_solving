package main

import (
  "fmt"
  "sync"
  "time"
)

var (
  lock = sync.Mutex{}
  wg = sync.WaitGroup{}
)
var cond *sync.Cond

var balance int

func stingy() {
  defer wg.Done()
  for {
    // Add 10 to balance.
    lock.Lock()
    balance += 10
    fmt.Println("stingy ", balance)
    cond.Signal()
    lock.Unlock()
    time.Sleep(10 * time.Second)
  }
}

func spendy() {
  defer wg.Done()
  for {
    lock.Lock()
     for  balance < 20 {
        cond.Wait()
     }
    balance -= 20
    fmt.Println("spendy ", balance)
    lock.Unlock()
    time.Sleep(5 * time.Second)
  }
}

func main() {
  cond = sync.NewCond(&lock)
  wg.Add(2)
  go stingy()
  go spendy()
  wg.Wait()
}



