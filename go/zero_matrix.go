package main

import (
  "fmt"
  "math/rand"
)


func valueInArray(array []int, value int) bool {
  if len(array) == 0 {
    return false
  }

  for _, num := range array {
    if num == value {
      return true
    }
  }
  return false
}


func zero(matrix [][]int) [][]int {
  rows := make([]int, 0)
  columns := make([]int, 0)

  for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[i]); j++ {
      if matrix[i][j] == 0 {
        rows = append(rows, i)
        columns = append(columns, j)
      }
    }
  }

  for _, rowNum := range rows {
    row := matrix[rowNum]

    for j := range row {
      matrix[rowNum][j] = 0
    }
  }

  for _, columnNum := range columns {
    for i := range matrix {
      matrix[i][columnNum] = 0
    }
  }

  return matrix
}


func main() {
  var M int = 6
  var N int = 8
  rand.Seed(11)

  matrix := make([][]int, M)

  for i := 0; i < M; i++ {
    matrix[i] = make([]int, N)
    for j := 0; j < N; j++ {
      matrix[i][j] = rand.Intn(30)
    }
  }

  fmt.Println(matrix)
  fmt.Println(zero(matrix))
}
