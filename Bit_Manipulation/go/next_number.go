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


func clearBit(num uint32, i int) uint32 {
  var mask uint32 = ^(1 << i)
  return num & mask
}


func getNext(num uint32) uint32 {
  var trailing bool = true
  var next uint32

  for i := 0; i < 32; i++ {
    if !getBit(num, i) {
      if !trailing {
        next = setBit(num, i)
        break
      }
    } else {
      trailing = false
    }

    if i == 31 {
      return 0
    }
  }

  for i := 0; i < 32; i++ {
    if getBit(num, i) {
      next = clearBit(next, i)
      return next
    }
  }

  return 0
}


func getPrev1(num uint32) uint32 {
  var prev uint32
  var firstOneIndex int

  for i := 0; i < 32; i++ {
    if getBit(num, i) {
      firstOneIndex = i
      prev = clearBit(num, i)
      break
    }
  }

  for i := 0; i < 32; i++ {
    if i + 1 == firstOneIndex{
      prev = setBit(prev, i)
    }
  }

  return prev
}


func getPrev2(num uint32) uint32 {
  var prev uint32 = num
  var ones, index int
  var trailing bool

  for i := 0; i < 32; i++ {
    if getBit(num, i) && trailing {
      index = i
      break
    } else if getBit(num, i) {
      ones++
    } else {
      trailing = true
    }
  }

  for i := 0; i < 32; i++ {
    if i == index {
      prev = clearBit(prev, i)
      break
    } else if index - ones - 1 <= i {
      prev = setBit(prev, i)
    } else {
      prev = clearBit(prev, i)
    }
  }

  return prev
}


func getPrev(num uint32) uint32 {
  if !getBit(num, 0) {
    return getPrev1(num)
  }

  return getPrev2(num)
}


func main() {
  var a, next, prev uint32
  a = 368

  fmt.Printf("%032b\n", a)

  next = getNext(a)
  fmt.Printf("%032b\n", next)

  prev = getPrev(a)
  fmt.Printf("%032b\n", prev)
}
