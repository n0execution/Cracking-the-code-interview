package main
import (
  "fmt"
  "strconv"
)


func compress(s string) string {
  var newString string
  currentLetter := string(s[0])
  var count int
  for i := 0; i < len(s); i++ {
    if string(s[i]) == currentLetter {
      count++
    } else {
      newString += currentLetter
      newString += strconv.Itoa(count)

      currentLetter = string(s[i])
      count = 1
    }
  }

  newString += currentLetter
  newString += strconv.Itoa(count)

  if len(newString) < len(s) {
    return newString
  }

  return s
}


func main() {
  str := "aabbbaccc"

  fmt.Println(compress(str))
}
