package main

import (
  "fmt"
)


func setBit(num *uint32, i int) {
  *num = *num | (1 << i)
}


func getBit(num uint32, i int) bool {
  return ((num & (1 << i)) != 0)
}


func clearMiddle(num *uint32, i int, j int) {
  var mask uint32

  for shift := 0; shift < (j - i); shift++ {
    mask = mask + (1 << uint32(shift + i))
  }

  *num = *num & ^mask
}


func insert(n *uint32, m uint32, i int, j int) {
  clearMiddle(n, i, j)

  for shift := 0; shift < (j - i + 1); shift++ {
    if getBit(m, shift) {
      setBit(n, shift + i)
    }
  }
}


func main() {
  var a, b uint32
  a = 1024
  b = 18

  fmt.Printf("%032b\n", a)
  fmt.Printf("%032b\n\n", b)

  insert(&a, b, 2, 6)
  fmt.Printf("%032b\n", a)
}
