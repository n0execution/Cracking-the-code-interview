def compress(string):
    new_string = ""
    current_letter = string[0]
    count = 0

    for letter in string:
        if letter == current_letter:
            count += 1
        else:
            new_string += current_letter + str(count)

            current_letter = letter
            count = 1

    new_string += letter + str(count)

    if len(new_string) < len(string):
        return new_string
    return string


def main():
    string = "aabbbaccc"
    print(compress(string))


if __name__ == '__main__':
    main()
