package main

import (
  "fmt"
)

type coordinates struct {
  x int
  y int
}

var cArray []coordinates

func drawCircle(radius int) []coordinates{
  x:=1
  y:=0

  for x*x <= radius {
    for y=0; y<x+1; y++ {
      if x*x+y*y == radius {
        cArray = append(cArray, coordinates{x:x,y:y})
        cArray = append(cArray, coordinates{x:x,y:-y})
        cArray = append(cArray, coordinates{x:-x,y:-y})
        cArray = append(cArray, coordinates{x:-x,y:y})
        cArray = append(cArray, coordinates{x:y,y:x})
        cArray = append(cArray, coordinates{x:y,y:-x})
        cArray = append(cArray, coordinates{x:-y,y:-x})
        cArray = append(cArray, coordinates{x:-y,y:x})
      }
    }
    x+= 1
  }

  return cArray
}

func main() {
  fmt.Println(drawCircle(100000))
}
