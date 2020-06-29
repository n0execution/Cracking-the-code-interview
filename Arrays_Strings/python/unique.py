# O(n^2)
def is_unique1(string):
    unique = ""

    for c in string: # O(n)
        if c not in unique: # O(n)
            unique += c
        else:
            return False
    return True


# O(n^2)
def is_unique2(string):
    for i, ch1 in enumerate(string): # O(n)
        for ch2 in string[i + 1:]: # O(n)
            if ch1 == ch2:
                return False
    return True


# O(n)
def is_unique3(string):
    uniques_dict = {}

    for c in string: # O(n)
        if uniques_dict.get(c): # O(1)
            return False
        else:
            uniques_dict[c] = True
    return True


# O(n^2*log(n))
def is_unique4(string):
    sorted_string = sorted(string) # O(n*log(n))
    for i in range(len(sorted_string) - 1): #  O(n)
        if sorted_string[i] == sorted_string[i + 1]:
            return False
    return True


def main():
    string = "vfzabcde"
    print(is_unique4(string))


if __name__ == '__main__':
    main()
