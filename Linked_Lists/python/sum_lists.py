from LinkedList import  LinkedList, Node


def generate_list(n, reverse=False):
    result_list = LinkedList()

    while n // 10 != 0 or n % 10 != 0:
        remainder = n % 10

        if reverse:
            result_list.append(remainder)
        else:
            result_list.insert_at_index(0, remainder)

        n = n / 10

    return result_list


def sum_list_reverse(node, tens):
    if node.next is None:
        return node.data * pow(10, tens)

    return sum_list_reverse(node.next, tens + 1) + node.data * pow(10, tens)


def sum_lists_reverse(node1, node2):
    return generate_list(sum_list_reverse(node1, 0) + sum_list_reverse(node2, 0),
                         reverse=True)


def sum_list(node):
    if node.next is None:
        return node.data, 1

    value, tens = sum_list(node.next)

    return node.data * pow(10, tens) + value, tens + 1


def sum_lists(node1, node2):
    first_value, _ = sum_list(node1)
    second_value, _ = sum_list(node2)

    return generate_list(first_value + second_value)


def main():
    list1, list2 = LinkedList(), LinkedList()
    list1.append(7)
    list1.append(1)
    list1.append(6)

    list2.append(5)
    list2.append(9)
    list2.append(2)

    print(sum_lists_reverse(list1.head, list2.head))
    print(sum_lists(list1.head, list2.head))

if __name__ == '__main__':
    main()
