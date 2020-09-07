def slow_multiply(a, b):
    if b > 0:
        return a + multiply(a, b - 1)
    return 0


def multiply(a, b, memo={}):
    if not a or not b:
        return 0

    if (a, b) in memo.keys():
        return memo[(a, b)]

    if b > a:
        a, b = b, a # a always bigger than b

    if b == 1:
        return a

    if a % 2 == 0:
        memo[(a, b)] = multiply(a / 2, b) + multiply(a / 2, b)
    else:
        memo[(a, b)] = multiply(a / 2, b) + multiply(a / 2, b) + b

    return memo[(a, b)]


def main():
    print(multiply(32, 1))


if __name__ == '__main__':
    main()
