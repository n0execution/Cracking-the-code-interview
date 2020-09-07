def find_all_subsets(s, subsets):
    if s not in subsets:
        subsets.append(s)

    for i in range(len(s)):
        copy_set = s[:]
        copy_set.remove(copy_set[i])

        if copy_set not in subsets:
            subsets.append(copy_set)

        find_all_subsets(copy_set, subsets)

    return sorted(subsets, key=len)


def find_subsets(s, subsets):
    subsets.append([])

    for i in range(len(s)):
        copy_subsets = subsets[:]
        for subset in copy_subsets:
            subsets.append([*subset, s[i]])

    return subsets


def main():
    s = ["a", "b", "c", "d", "e"]
    subsets = []
    print(find_all_subsets(s, subsets))
    subsets = []
    print(find_subsets(s, subsets))


if __name__ == '__main__':
    main()
