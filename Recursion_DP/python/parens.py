def get_all_parens(n):
    parens = []

    if n == 1:
        return ['()']
    else:
        for p in get_all_parens(n - 1):
            parens.append(f"({p})")

            if f"(){p}" not in parens:
                parens.append(f"(){p}")

            if f"{p}()" not in parens:
                parens.append(f"{p}()")

    return parens



def main():
    print(get_all_parens(3))


if __name__ == '__main__':
    main()
