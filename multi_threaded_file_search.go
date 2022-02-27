package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
  "sync"
)

// wait group
var (
  wg = sync.WaitGroup{}
)

var matches []string
var lock sync.Mutex


func findAllMatches(dirName, fileName string) {
  // Find all files in the directory
  files, errFiles := ioutil.ReadDir(dirName)
  if errFiles != nil {
    fmt.Println("Failed to read the directory", errFiles)
    return
  }

  // Check for the presence of the file in this directory
  for _, file := range files {
    // If the file is a directory, recursively call the function passing
    // the canonical path.
    if file.IsDir() {
      // Add to wait group
      wg.Add(1)
      go findAllMatches(filepath.Join(dirName, file.Name()), fileName)
    }
    // It is a file. Check if the filename matches
    if strings.Contains(file.Name(), fileName) {
      // Add the whole path to matches slice
      lock.Lock()
      matches = append(matches, filepath.Join(dirName, file.Name()))
      lock.Unlock()
    }
  }
  // All go routines are done.
  wg.Done()
}

func main() {
  // Find all matches of a given file in a given directory
  wg.Add(1)
  go findAllMatches(os.Args[1], os.Args[2])
  wg.Wait()
  fmt.Println(matches)
}
