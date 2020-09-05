def count_steps(num):
    if num < 0:
        return 0
    elif num == 0:
        return 1

    return count_steps(num - 3) + count_steps(num - 2) + count_steps(num - 1)


def cached_count_steps(num, cache):
    if cache.get(num):
        return cache.get(num)

    if num < 0:
        cache[num] = 0
    elif num == 0:
        cache[num] = 1
    else:
        cache[num] = cached_count_steps(num - 3, cache) \
                     + cached_count_steps(num - 2, cache) \
                     + cached_count_steps(num - 1, cache)

    return cache[num]


def main():
    cache = {}
    print(count_steps(4))
    print(cached_count_steps(4, cache))


if __name__ == '__main__':
    main()
