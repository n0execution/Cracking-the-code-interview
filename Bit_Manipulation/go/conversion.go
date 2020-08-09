package main

import (
  "fmt"
)


func getBit(num uint32, i int) bool {
  return ((num & (1 << i)) != 0)
}


func getOnesCount(num uint32) int {
  var count int

  for i := 0; i < 32; i++ {
    if getBit(num, i) {
      count++
    }
  }

  return count
}


func determineNumber(num1 uint32, num2 uint32) int {
  return getOnesCount(num1 | num2) - getOnesCount(num1 & num2)
}


func main() {
  var a, b uint32
  a = 342
  b = 330

  fmt.Printf("%032b\n", a)
  fmt.Printf("%032b\n\n", b)

  fmt.Println(determineNumber(a, b))
}
