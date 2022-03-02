// A simple program to spawn multiple goroutines and pass tasks to each goroutine using buffered channels
// This program reads URL from STDIN and runs HTTP GET on the URL and dumps body on screen
// Basic program to build crawler would be to read robot.txt and parse all the allowed URLs in there and start indexing them.


package main

import (
  "fmt"
  "io"
  "sync"
  "net/http"
)


const (
  MAXCONCURRENCY int = 32
)

var wg sync.WaitGroup

type task struct {
  url string
}

func worker(ch chan task) {
  defer wg.Done()
  for {
    job := <-ch
    fmt.Println("Executing GET of ", job.url)
    resp, _ := http.Get(job.url)
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
      bodyBytes, _ := io.ReadAll(resp.Body)
      bodyString := string(bodyBytes)
      fmt.Println(bodyString)
    }
    fmt.Println("Done executing job ", job.url)
  }
}

func main() {
  var input string
  ch := make(chan task, MAXCONCURRENCY)

  for i:=0; i< MAXCONCURRENCY; i++ {
    go worker(ch)
  }
  wg.Add(MAXCONCURRENCY)

  for {
    // TODO:
    // Take input from a file instead of reading from stdin
    fmt.Printf("URL to read:")
    fmt.Scanln(&input)
    t := task { url: input}
    ch <- t
  }
  wg.Wait()
}


