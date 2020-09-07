def get_index_distinct(array, start, end):
    index = (start + end) / 2

    if array[index] == index:
        return index
    elif array[index] < index:
        return get_index_distinct(array, index + 1, end)
    elif array[index] > index:
        return get_index_distinct(array, start, index - 1)


def get_index(array, start, end):
    index = (start + end) / 2

    if end < start:
        return -1

    if array[index] == index:
        return index

    left = get_index(array, start, index - 1)
    if left >= 0:
        return left

    right = get_index(array, index + 1, end)
    if right >= 0:
        return right


def main():
    array = [-10, -8, 0, 2, 4, 8, 10]
    array = [-1, 1, 2, 4, 30, 44]
    array = [-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 12]
    # array = [-10, -5, 1, 3, 30, 40, 50, 60, 90, 110, 120]
    # print(get_index(array, len(array) / 2))
    print(get_index(array, 0, len(array)))


if __name__ == '__main__':
    main()
