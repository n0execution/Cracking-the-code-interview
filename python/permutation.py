def remove_symbol(str, c):
    removed = str.replace(c, "", 1)

    if len(removed) == len(str):
        return removed, False
    return removed, True


def check_permutation1(str1, str2):
    if len(str1) != len(str2):
        return False

    for c in str1:
        if c not in str2:
            return False

        str2, success = remove_symbol(str2, c)

        if not success:
            return False
    return True


# O(n*log(n))
def check_permutation2(str1, str2):
    if len(str1) != len(str2):
        return False

    sorted_string1, sorted_string2 = sorted(str1), sorted(str2)

    if sorted_string1 == sorted_string2:
        return True
    return False


# O(n)
def check_permutation3(str1, str2):
    if len(str1) != len(str2):
        return False

    counts_dict = {}

    for c in str1:
        if counts_dict.get(c):
            counts_dict[c] += 1
        else:
            counts_dict[c] = 1

    for c in str2:
        if counts_dict.get(c):
            counts_dict[c] -= 1
        else:
            return False

    return True


def main():
    string1 = "oloa"
    string2 = "laoo"
    print(check_permutation3(string1, string2))


if __name__ == '__main__':
    main()
