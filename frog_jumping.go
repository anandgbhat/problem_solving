package main

import (
  "fmt"
  "math"
)

func min(i,j int) int {
  if i> j {
    return j
  }
  return i
}

func frogJumpMinEnergy(energy []int, idx int) int {
  // Base condition
  if idx == 0 {
    return 0
  }

  // Frog can jump either one step or two steps at each point
  // Calculate all possible combinations and take the minimum of the
  // energy spent.
  var twoSteps int
  oneStep := frogJumpMinEnergy(energy, idx-1) +
    int(math.Abs(float64(energy[idx] - energy[idx-1])))
  if idx > 1 {
    twoSteps = frogJumpMinEnergy(energy, idx-2) +
      int(math.Abs(float64(energy[idx] - energy[idx-2])))
  }
  return min(oneStep, twoSteps)
}

func main() {
  nums := []int {30, 10, 60, 10, 60, 50}
  fmt.Println("minimum energy spent in reaching from 0 to 5 = ",
    frogJumpMinEnergy(nums, len(nums)-1))
}
