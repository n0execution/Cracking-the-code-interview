from numpy import uint32


def get_bit(num, i):
    return ((num & (1 << i)) != 0)


def get_ones_count(num):
    count = 0

    for i in range(32):
        if get_bit(num, i):
            count += 1

    return count


def determine_number(num1, num2):
    return get_ones_count(num1 | num2) - get_ones_count(num1 & num2)


def main():
    num1 = uint32(244)
    num2 = uint32(249)

    print("{0:b}".format(num1))
    print("{0:b}".format(num2))

    print(determine_number(num1, num2))


if __name__ == '__main__':
    main()
