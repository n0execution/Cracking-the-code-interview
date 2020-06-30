from LinkedList import  LinkedList, Node
from random import randint


def k_to_last(head: Node, k):
    if head is None:
        return 0

    index  = k_to_last(head.next, k) + 1

    if index == k:
        print(k, head)

    return index


def k_to_last2(head: Node, k):
    p1 = head
    p2 = head

    for i in range(k):
        p2 = p2.next

    while p2.next:
        p1 = p1.next
        p2 = p2.next

    print(k, p1)


def main():
    linked_list = LinkedList()
    for i in range(10):
        linked_list.append(randint(0, 5))
    print(linked_list)

    k_to_last2(linked_list.head, 5)


if __name__ == '__main__':
    main()
