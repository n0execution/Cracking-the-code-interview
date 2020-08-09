from numpy import uint32
from exceptions import *


def get_bit(num, i):
    return ((num & (1 << i)) != 0)


def set_bit(num, i):
    return num | (1 << i)


def clear_bit(num, i):
    mask = ~(1 << i)

    return num & mask


def get_ones_count(num):
    count = 0

    for i in range(32):
        if get_bit(num, i):
            count += 1

    return count


def get_next(num):
    trailing = True

    for i in range(32):
        if not get_bit(num, i):
            if not trailing:
                next = set_bit(num, i)
                break
        else:
            trailing = False

        if i == 31:
            raise NoOnesException(num)

    for i in range(32):
        if get_bit(num, i):
            next = clear_bit(next, i)
            return next


def get_prev1(num):
    for i in range(32):
        if get_bit(num, i):
            first_one_index = i
            prev = clear_bit(num, i)
            break

    for i in range(32):
        if i + 1 == first_one_index:
            prev = set_bit(prev, i)

    return prev


def get_prev2(num):
    ones = 0
    trailing = False

    for i in range(32):
        if get_bit(num, i) and trailing:
            index = i
            break
        elif get_bit(num, i):
            ones += 1
        else:
            trailing = True

    prev = num

    for i in range(32):
        if i == index:
            prev = clear_bit(prev, i)
            return prev
        elif index - ones - 1 <= i:
            prev = set_bit(prev, i)
        else:
            prev = clear_bit(prev, i)



def get_prev(num):
    if not get_bit(num, 0):
        return get_prev1(num)

    return get_prev2(num)


def main():
    num = uint32(369)
    print("{0:b}".format(num))

    # next = get_next(num)
    # print("{0:b}".format(next))

    prev = get_prev(num)
    print("{0:b}".format(prev))


if __name__ == '__main__':
    main()
