package main

import (
  "fmt"
)

func swap(nums []int, i,j int) {
  tmp := nums[i]
  nums[i] = nums[j]
  nums[j] = tmp
}

func cyclicSort(nums []int)  []int {
  var j int
  // Cyclic sort is to place the element at the correct spot
  // If the element is not at the correct place, swap the elements
  // If array given is [2,6,4,3,1,5], then first element 2 is not at the
  // right place. Swap it with element at position 2 which will make
  // the array [6,2,4,3,1,5]. Now 6 is not in correct place. swap it with 6th
  // element and so on till we get the right value. Then move to next value.

  for i:=0; i< len(nums); i++ {
    // Check if the value at nums[i] is at correct place. Else, swap it
    // with element at index nums[i]
    for nums[i] != i+1 { // i+1 to account for 0 starting element
      j = nums[i]-1
      swap(nums, i, j)
    }
  }

  return nums
}

func main() {
  nums := []int {2,6,4,3,1,5}
  fmt.Println(nums)
  fmt.Println("Sorted array  = ", cyclicSort(nums))
}
