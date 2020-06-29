package main
import (
  "fmt"
  "sort"
  "strings"
)


func isIn(str1 string, str2 string) bool {
  if len(str1) == 0 {
    return false
  }

  for _, s := range str1 {
    if string(s) == str2 {
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


func removeSymbol(str string, c string) (string, bool) {
  removed := strings.Replace(str, c, "", 1)

  if len(removed) == len(str) {
    return removed, false
  }
  return removed, true
}


// O(n^2)
func checkPermutation1(s1 string, s2 string) bool{
  if len(s1) != len(s2) {
    return false
  }

  success := false

  for i := 0; i < len(s1); i++ { // O(n)
    if !isIn(s2, string(s1[i])) {
      return false
    }

    s2, success = removeSymbol(s2, string(s1[i])) // O(n)

    if !success {
      return false
    }
  }

  return true
}


// O(n * lon(n))
func checkPermutation2(s1 string, s2 string) bool{
  if len(s1) != len(s2) {
    return false
  }

  sorted1, sorted2 := SortString(s1), SortString(s2)

  if sorted1 == sorted2 {
    return true
  }

  return false
}


// O(n)
func checkPermutation3(s1 string, s2 string) bool{
  if len(s1) != len(s2) {
    return false
  }

  counts := make(map[string]int)

  for i := 0; i < len(s1); i++ {
    str := string(s1[i])
    counts[str]++
  }

  for i := 0; i < len(s2); i++ {
    str := string(s2[i])
    if counts[str] > 0 {
      counts[str]--
    } else {
      return false
    }
  }

  return true
}


func main() {
  str1 := "oloa"
  str2 := "laos"

  fmt.Println(checkPermutation3(str1, str2))
}
