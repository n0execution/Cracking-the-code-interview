from numpy import uint32


def get_bit(num, i):
    return ((num & (1 << i)) != 0)


def set_bit(num, i):
    return num | (1 << i)


def get_ones_count(num):
    count, length = 0, 0

    for i in range(32):
        if get_bit(num, i):
            count += 1
        else:
            if count > length:
                length = count
            count = 0

    return length


def flip(num):
    length = 0
    for i in range(32):
        if not get_bit(num, i):
            flipped = set_bit(num, i)

            ones = get_ones_count(flipped)

            if ones > length:
                length = ones

    return length



def main():
    num = uint32(383)
    print("{0:b}".format(num))

    print(flip(num))


if __name__ == '__main__':
    main()
