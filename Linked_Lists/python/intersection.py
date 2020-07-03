from LinkedList import LinkedList, Node


def intersection(node1, node2):
    length1, length2 = 1, 1
    head1, head2 = node1, node2

    while node1.next is not None:
        last_node1 = node1.next
        node1 = node1.next
        length1 += 1

    while node2.next is not None:
        last_node2 = node2.next
        node2 = node2.next
        length2 += 1

    if last_node1 != last_node2:
        return None

    node1, node2 = head1, head2

    if length1 > length2:
        for i in range(length1 - length2):
            node1 = node1.next
    elif length1 < length2:
        for i in range(length2 - length1):
            node2 = node2.next

    while node1.next is not None:
        if node1.next == node2.next:
            return node1.next

        node1 = node1.next
        node2 = node2.next


def main():
    intersection_list, list1, list2 = LinkedList(), LinkedList(), LinkedList()

    intersection_list.append(1)
    intersection_list.append(6)
    intersection_list.append(9)

    list1.append(2)
    list1.append(3)

    node1 = list1.head
    while node1.next is not None:
        node1 = node1.next

    node1.next = intersection_list.head

    list2.append(2)

    node2 = list2.head
    while node2.next is not None:
        node2 = node2.next

    node2.next = intersection_list.head

    print(list1)
    print(list2)
    print(intersection(list1.head, list2.head))


if __name__ == '__main__':
    main()
