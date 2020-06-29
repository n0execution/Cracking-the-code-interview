def is_rotation(str1, str2):
    if len(str1) != len(str2):
        return False

    for index, letter in enumerate(str2):
        if letter == str1[0]:
            left, right = str2[:index], str2[index:]
            print(right + left)

            if right + left == str1:
                return True

    return False


def is_substring(str1, str2):
    if len(str1) != len(str2) / 2:
        return False

    index = 0
    result = ""

    for letter in str2:
        if result == str1:
            return True

        if letter == str1[index]:
            index += 1
            result += letter
        else:
            if index != 0:
                result = ""
                index = 0

    if result == str1:
        return True

    return False



def main():
    str1 = "waterbottle"
    str2 = "erbottlewat"

    print(is_rotation(str1, str2 + str2))


if __name__ == '__main__':
    main()
