def urlify(string, length):
    spaces = 0

    print(string)

    for s in string[:length]:
        if s == " ":
            spaces += 1

    index = length + spaces * 2

    i = length - 1

    for s in string[:length][::-1]:
        if s != " ":
            string[index - 1] = string[i]
            index -= 1
        else:
            string[index - 1] = "0"
            string[index - 2] = "2"
            string[index - 3] = "%"
            index -= 3
        i -= 1

    return ''.join(string)


def main():
    string = "Mr John Smith    "
    print(urlify(list(string), 13))


if __name__ == '__main__':
    main()
