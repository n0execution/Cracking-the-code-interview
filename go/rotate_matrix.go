package main

import ("fmt")


func rotate(matrix [][]int) [][]int {
  fmt.Println(matrix)

  length := len(matrix)

  for layer := 0; layer < length / 2; layer++ {
    first := layer
    last := length - 1 - layer

    for i := first; i < last; i++ {
      top := matrix[first][i]

      // left to top
      matrix[first][i] = matrix[last - i + first][first]

      // bottom to left
      matrix[last - i + first][first] = matrix[last][last - i + first]

      // right to bottom
      matrix[last][last - i + first] = matrix[i][last]

      // top to right
      matrix[i][last] = top
    }
  }

  return matrix
}


func main() {
  matrix := make([][]int, 5)

  var counter = 1

  for i := 0; i < 5; i++ {
    matrix[i] = make([]int, 5)
    for j := 0; j < 5; j++ {
        matrix[i][j] = counter
        counter++
      }
    }
    fmt.Println(rotate(matrix))
}
