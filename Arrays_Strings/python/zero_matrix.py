from random import randint


def zero_matrix(matrix):
    rows = []
    columns = []

    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            if matrix[i][j] == 0:
                rows.append(i)
                columns.append(j)

    for row_num in rows:
        row = matrix[row_num]

        for j in range(len(row)):
            matrix[row_num][j] = 0

    for column_num in columns:
        for i in range(len(matrix)):
            matrix[i][column_num] = 0


    return matrix


def main():
    m = 2
    n = 3

    matrix = [[0]*n for i in range(m)]

    for i in range(m):
        for j in range(n):
            matrix[i][j] = randint(0, 5)

    print(matrix)
    print(zero_matrix(matrix))

if __name__ == '__main__':
    main()
