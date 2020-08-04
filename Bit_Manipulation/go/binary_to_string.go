package main


import (
  "fmt"
)


func getBinaryRepresentation(num float64) string {
  var binary string = "."
  var frac float64 = 0.5
  if num >= 1 || num <= 0 {
    return "ERROR"
  }

  for num > 0 {
    if len(binary) > 32 {
      return "ERROR"
    }

    if num >= frac {
      binary += "1"
      num -= frac
    } else {
      binary += "0"
    }

    frac /= 2
  }

  return binary
}


func main() {
  var num float64 = 0.72

  fmt.Println(getBinaryRepresentation(num))
}
