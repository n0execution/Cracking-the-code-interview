from numpy import uint32


def get_bit(num, i):
    return ((num & (1 << i)) != 0)


def set_bit(num, i):
    return num | (1 << i)


def clear_bit(num, i):
    mask = ~(1 << i)

    return num & mask


def clear_middle(num, i, j):
    mask = 0
    for shift in range(j - i):
        mask += (1 << shift + i)
    return num & ~mask


def insert(n, m, i, j):
    n = clear_middle(n, i, j)

    for shift in range(j - i + 1):
        if get_bit(m, shift):
            n = set_bit(n, shift + i)

    return n

def main():
    num1 = uint32(1024)
    num2 = uint32(19)
    print("{0:b}".format(num1))
    print("{0:b}".format(num2))

    result = insert(num1, num2, 2, 6)
    print("{0:b}".format(result))


if __name__ == '__main__':
    main()
