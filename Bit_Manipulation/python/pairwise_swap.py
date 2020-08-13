from numpy import uint32


def rshift(val, n):
    return val>>n if val >= 0 else (val+0x100000000)>>n


def swap(num):
    mask1 = sum([2 ** i for i in range(1, 32, 2)])
    mask2 = sum([2 ** i for i in range(0, 32, 2)])

    return (rshift(num & mask1, 1)) | ((num & mask2) << 1)


def main():
    num = uint32(350)
    print("{0:b}".format(num))

    print("{0:b}".format(swap(num)))


if __name__ == '__main__':
    main()
