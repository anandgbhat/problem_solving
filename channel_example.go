package main

import (
    "fmt"
    "time"
    "sync"
)

var wg sync.WaitGroup

func pinger(c chan string) {
    for i := 0; i<10 ; i++ {
        c <- "ping"
    }
    wg.Done()
}
func ponger(c chan string) {
    for i := 0;i<10 ; i++ {
        c <- "pong"
    }
    wg.Done()
}
func printer(c chan string) {
    for i:=0; i<20; i++{
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
    wg.Done()
}
func main() {
    var c chan string = make(chan string)
    wg.Add(1)
    go pinger(c)
    wg.Add(1)
    go ponger(c)
    wg.Add(1)
    go printer(c)
    wg.Wait()
}
