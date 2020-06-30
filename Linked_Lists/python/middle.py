from LinkedList import  LinkedList, Node
from random import randint


class StaticVar(object):
    length = 0


def delete_middle(node):
    node.data = node.next.data
    node.next = node.next.next


def get_middle(head: LinkedList, s):
    if head.next is None:
        s.length += 1
        return s.length

    s.length += 1
    index = get_middle(head.next, s) - 1

    if index == s.length // 2:
        delete_middle(head)

    return index


def main():
    linked_list = LinkedList()
    for i in range(5):
        linked_list.append(randint(0, 5))
    print(linked_list)

    s = StaticVar
    get_middle(linked_list.head, s)

    print(linked_list)



if __name__ == '__main__':
    main()
