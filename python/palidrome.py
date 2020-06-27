def normalize_string(string):
    return string.replace(" ", "").lower()


def is_palindrome(string):
    counts_dict = {}
    odds_count = 0

    string = normalize_string(string)

    for s in string:
        if counts_dict.get(s):
            counts_dict[s] += 1
        else:
            counts_dict[s] = 1

    for v in counts_dict.values():
        if v % 2 == 1:
            odds_count += 1

        if odds_count > 1:
            return False

    return True


def main():
    string = "So patient a nurse to nurse a patient so"

    print(is_palindrome(string))


if __name__ == '__main__':
    main()
