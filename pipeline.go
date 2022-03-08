// Toy pipeline implementation with multiple channels
// First channel receives []int and sorts them and send the
// sorted []int to channel two. Next function reveives sorted array
// from channel two and sums it up and sends the result on channel three.
// Next function in the pipeline reads the sum from channel three and squares
// the value and writes to channel 4. Main function inputs the initial
// unsorted array on channel 1 and receives the final o/p from channel 4.

package main

import (
  "fmt"
  "sort"
)

func main() {
  chanOne := make(chan []int)
  chanTwo := make(chan []int)
  chanThree := make(chan int)
  chanFour := make(chan int)

  go sortElements(chanOne, chanTwo)
  go addValues(chanTwo, chanThree)
  go squareResult(chanThree, chanFour)

  nums := []int{1, 5, 3, 2, 4}
  chanOne <- nums
  finalVal := <-chanFour
  fmt.Println(finalVal)

}


func sortElements(one chan []int, two chan []int) {
  // Receive array from channel one
  values := <-one
  // Soft the array and send it to channel two
  sort.Ints(values)
  // Send the result to channel two
  two <- values
}

func addValues(two chan []int, three chan int) {
  sum := 0
  sortedValues := <-two
  // Add the values
  for _, val := range sortedValues {
    sum += val
  }
  three <- sum
}

func squareResult(three chan int, four chan int) {
  sum := <-three
  four <-(sum * sum)
}
