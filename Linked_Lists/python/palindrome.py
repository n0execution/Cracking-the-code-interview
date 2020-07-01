from LinkedList import  LinkedList


def compare_lists(node1, node2):
    while node1.next is not None:
        if node1.data != node2.data:
            return False

        node1 = node1.next
        node2 = node2.next

    return True


def reverse_linked_list(node):
    if node.next is None:
        result_list = LinkedList()
        result_list.append(node.data)

        return result_list

    result_list = reverse_linked_list(node.next)
    result_list.append(node.data)

    return result_list


def is_palindrome(linked_list):
    reversed_list = reverse_linked_list(linked_list.head)

    return compare_lists(linked_list.head, reversed_list.head)


def main():
    linked_list = LinkedList()
    linked_list.append(1)
    linked_list.append(1)
    linked_list.append(2)
    linked_list.append(3)
    linked_list.append(2)
    linked_list.append(1)
    linked_list.append(1)

    print(is_palindrome(linked_list))


if __name__ == '__main__':
    main()
