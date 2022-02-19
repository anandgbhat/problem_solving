package main

import (
  "fmt"
)

//
// P1---------P2
// |
// |
// |
// | P3       P4
//
// If the points form a square, then distP1P2 should be same as
// distP1P3 and distP1P4 = sqrt(2) * distP1P2

type Point struct {
  x int
  y int
}

func distance(p1,p2 Point) int {
  return (p2.x-p1.x) * (p2.x-p1.x) + (p2.y-p1.y) * (p2.y - p1.y)
}

func isSquare(p1, p2, p3, p4 Point) bool {
  distP1P2 := distance(p1,p2)
  distP1P3 := distance(p1,p3)
  distP1P4 := distance(p1,p4)
  fmt.Println(distP1P2, distP1P3, distP1P4)

  // Check if distP1P2 is same as distP1P3
  if distP1P2 == distP1P3 {
    // distanceP1P4  should be equal to
    // 2 * distanceP1P2
    if distP1P4  == 2 * distP1P2  {
      return true
    }
  }

  return false
}


func main() {
  p1 := Point {
      x: 10,
      y: 20,
    }

  p2 := Point {
      x: 20,
      y: 20,
  }

  p3 := Point {
      x: 10,
      y: 10,
  }

  p4 := Point {
      x: 20,
      y: 10,
  }

  fmt.Println("Is square ", isSquare(p1, p2, p3, p4))
}
