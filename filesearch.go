package main

import (
  "fmt"
  "strings"
  "sync"
  "io/ioutil"
  "path/filepath"
)

var matches []string

var wg sync.WaitGroup

func fileSearch(root string, fileName string) {
  defer wg.Done()
  if len(root) == 0 || len(fileName) == 0 {
    return
  }
  // Read the directory
  files, err := ioutil.ReadDir(root)
  if err != nil {
    // fmt.Println("Failed to read dir ", root, ". Error = ", err)
    return
  }

  // Walk through results and check if the file is a directory.
  // If so, recursively call the function.
  // If not, check if the filename matches. If so, append to matches slice.
  for _, file := range files {
    if strings.Contains(file.Name(), fileName) {
      matches = append(matches, filepath.Join(root, file.Name()))
    }
    if file.IsDir() {
      wg.Add(1)
      go fileSearch(filepath.Join(root, file.Name()), fileName)
    }
  }
}

func main() {
  fileSearch("/Users/anand.bhat", "x.go")
  fmt.Println(matches)
  wg.Wait()
}
