package main
import (
  "fmt"
)


func urlify(s []byte, length int) string {
  spaces := 0
  for i := 0; i < length; i++ {
    if string(s[i]) == " "{
      spaces++
    }
  }

  index := length + spaces * 2
  for i := length - 1; i >= 0; i-- {
    if string(s[i]) != " " {
        s[index - 1] = s[i]
        index--

    } else {
      s[index - 1] = '0'
      s[index - 2] = '2'
      s[index - 3] = '%'

      index -= 3
    }
  }

  return string(s)
}


func main() {
  var str []byte = []byte("Mr John Smith    ")

  fmt.Println(urlify(str, 13))
}
