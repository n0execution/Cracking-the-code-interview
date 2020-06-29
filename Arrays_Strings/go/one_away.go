package main

import (
  "fmt"
  "math"
)


func removeSymbol(s string, index int) string {
  return s[:index] + s[index + 1:]
}


func checkEdit(s1 string, s2 string) bool {
  length1 := float64(len(s1))
  length2 := float64(len(s2))

  if math.Abs(length1 - length2) >= 2 {
    return false
  }

  length := math.Min(length1, length2)

  i := 0
  for i < int(length) {
    if s1[i] == s2[i] {
      i++
    } else {
      if length1 == length2 {
        s1, s2 = removeSymbol(s1, i), removeSymbol(s2, i)
        length--
      } else if length1 > length2{
        s1 = removeSymbol(s1, i)
      } else if length1 < length2{
        s2 = removeSymbol(s2, i)
      }

      if s1 != s2{
        return false
      }
    }
  }

  return true
}


func main() {
  str1 := "palinr"
  str2 := "palidr"

  fmt.Println(checkEdit(str1, str2))
}
