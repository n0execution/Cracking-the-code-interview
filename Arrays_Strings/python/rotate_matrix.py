def rotate(matrix):
    length = len(matrix)

    for layer in range(length // 2):
        first = layer
        last = length - 1 - layer

        for i in range(first, last):
            top = matrix[first][i]

            # left to top
            matrix[first][i] = matrix[last - i + first][first]

            # bottom to left
            matrix[last - i + first][first] = matrix[last][last - i + first]

            # right to bottom
            matrix[last][last - i + first] = matrix[i][last]

            # top to right
            matrix[i][last] = top

    return matrix


def main():
    matrix = [[0]*5 for i in range(5)]

    counter = 1

    for i in range(5):
        for j in range(5):
            matrix[i][j] = counter
            counter += 1

    print(rotate(matrix))


if __name__ == '__main__':
    main()
