from LinkedList import LinkedList


def detect(node1):
    node2 = node1.next

    while node1 is not None:
        if node1 == node2:
            return node1

        node1 = node1.next
        node2 = node2.next.next

    return None


def main():
    linked_list = LinkedList()
    linked_list.append(1)
    linked_list.append(2)
    linked_list.append(3)
    linked_list.append(4)
    node = linked_list.head

    while node.next is not None:
        node = node.next
        loop_begin = node

    linked_list.append(5)
    linked_list.append(6)
    linked_list.append(7)

    node = linked_list.head

    while node.next is not None:
        node = node.next

    node.next = loop_begin

    print(detect(linked_list.head))


if __name__ == '__main__':
    main()
