package main

import (
  "fmt"
  "strings"
)


func normalizeString(s string) string{
  return strings.ToLower(strings.Replace(s, " ", "", -1))
}


func isPalindrome(s string) bool {
  s = normalizeString(s)
  counts := make(map[string]int)
  var oddsCount int

  for i := 0; i < len(s); i++ {
    counts[string(s[i])]++
  }

  for _, v := range counts {
    if v % 2 == 1 {
      oddsCount++
    }

    if oddsCount > 1 {
      return false
    }
  }

  return true
}


func main() {
  str := "Te Nte"

  fmt.Println(isPalindrome(str))
}
