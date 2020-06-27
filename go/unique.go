package main
import (
  "fmt"
  "sort"
  "strings"
)


func isIn(unique string, c string) bool {
  if len(unique) == 0 {
    return false
  }

  fmt.Println(unique, c)

  for _, s := range unique {
    if string(s) == c {
      return true
    }
  }
  return false
}


// SortString implementation
func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}


// O(n^2)
func isUnique1(s string) bool{
  unique := ""

  for i := 0; i < len(s); i++ {
    if !isIn(unique, string(s[i])) {
      unique += string(s[i])
    } else {
      return false
    }
  }
  return true
}


// O(n^2)
func isUnique2(s string) bool{
  for i := 0; i < len(s); i++ {
    for j := i + 1; j < len(s); j++ {
      if s[j] == s[i] {
        return false
      }
    }
  }
  return true
}


// O(n)
func isUnique3(s string) bool{
  uniques := make(map[byte]bool)

  for i := 0; i < len(s); i++ {
      if uniques[s[i]] {
        return false
      }
      uniques[s[i]] = true
  }

  return true
}


// O(n^2*log(n))
func isUnique4(s string) bool{
  sortedString := SortString(s)

  for i := 0; i < len(sortedString) - 1; i++ {
      if sortedString[i] == sortedString[i + 1] {
        return false
      }
    }
  return true
}


func main() {
  str := "vfzabcde"

  fmt.Println(isUnique4(str))
}
