package main

import (
  "fmt"
)


func getBit(num uint32, i int) bool {
  return ((num & (1 << i)) != 0)
}


func setBit(num uint32, i int) uint32 {
  return num | (1 << i)
}


func getOnesCount(num uint32) int {
  var count, length int

  for i := 0; i < 32; i++ {
    if getBit(num, i) {
      count++
    } else {
      if count > length {
        length = count
      }

      count = 0
    }
  }

  return length
}


func flip(num uint32) int {
  var length, ones int
  var flippedNum uint32

  for i := 0; i < 32; i++ {
    if !getBit(num, i) {
      flippedNum = setBit(num, i)

      ones = getOnesCount(flippedNum)
      if ones > length {
        length = ones
      }
    }
  }
  return length
}


func main() {
  var a uint32
  a = 383
  fmt.Printf("%032b\n", a)

  fmt.Println(flip(a))
}
