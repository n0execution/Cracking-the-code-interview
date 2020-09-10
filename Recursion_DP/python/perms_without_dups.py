def get_all_permutations(s):
    permutations = []

    if not s:
        return [""]

    for i, c in enumerate(s):
        copy_string = s.replace(s[i], '')

        for p in get_all_permutations(copy_string):
            for index in range(len(p) + 1):
                permutation = p[:index] + s[i] + p[index:]
                if permutation not in permutations:
                    permutations.append(permutation)

    return permutations



def main():
    print(get_all_permutations("abc"))
    print(get_all_permutations("abcd"))
    # print(get_all_permutations("abcd"))


if __name__ == '__main__':
    main()
