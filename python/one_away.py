def remove_symbol(string, index):
    return string[:index] + string[index + 1:]


def check_edit(str1, str2):
    if abs(len(str1) -  len(str2)) >= 2:
        return False

    i = 0
    length = min(len(str1), len(str2))

    while i < length:
        if str1[i] == str2[i]:
            i += 1
        else:
            if len(str1) == len(str2):
                str1, str2 = remove_symbol(str1, i), remove_symbol(str2, i)
                length -= 1
            elif len(str1) > len(str2):
                str1 = remove_symbol(str1, i)
            elif len(str1) < len(str2):
                str2 = remove_symbol(str2, i)

            if str1 != str2:
                return False
    return True


def main():
    str1 = "spale"
    str2 = "spales"

    print(check_edit(str1, str2))


if __name__ == '__main__':
    main()
