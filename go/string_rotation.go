package main

import (
  "fmt"
)


func isRotation1(str1 string, str2 string) bool {
  if len(str1) != len(str2) {
    return false
  }

  index := 0
  var result string

  for _, letter := range str2 {
    if string(letter) == string(str1[index]) {
      index++
      result += string(letter)
    } else {
      if index != 0 {
        result = ""
        index = 0
      }
    }
  }

  if result == str1{
    return true
  }

  for _, letter := range str2 {
    if string(letter) == string(str1[index]) {
      index++
      result += string(letter)

      if result == str1 {
        return true
      }
    } else {
      fmt.Println(result, index)
      return false
    }
  }

  return true
}


func isRotation2(str1 string, str2 string) bool {
  if len(str1) != len(str2) {
    return false
  }

  for index, letter := range str2 {
    if string(letter) == string(str1[0]) {
      left, right := str2[:index], str2[index:]

      fmt.Println(right + left)
      if right + left == str1 {
        return true
      }
    }
  }
  return false
}


func isSubstring(str1 string, str2 string) bool {
  if len(str1) != len(str2) / 2 {
    return false
  }

  index := 0
  var result string

  for _, letter := range str2 {
    if result == str1{
      return true
    }

    if string(letter) == string(str1[index]) {
      index++
      result += string(letter)
    } else {
      if index != 0 {
        result = ""
        index = 0
      }
    }
  }

  if result == str1{
    return true
  }

  return false

}


func main() {
  str1 := "waterbottle"
  str2 := "erbottlewat"

  fmt.Println(isSubstring(str1, str2 + str2))
}
