def get_binary_representation(num):
    binary = "."
    frac = 0.5

    if 0 <= num <= 1:
        return "ERROR"

    while num > 0:
        if len(binary) > 32:
            return "ERROR"

        if num >= frac:
            binary += "1"
        else:
            binary += "0"

        frac /= 2

    return binary


def main():
    num = 0.72

    print(get_binary_representation(num))


if __name__ == '__main__':
    main()
