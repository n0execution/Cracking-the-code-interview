from LinkedList import  LinkedList, Node
from random import randint


def remove_duplicates1(linked_list: LinkedList):
    nums = []
    offset = 0

    for i, node in enumerate(linked_list):
        if node.data in nums:
            linked_list.remove_at_index(i - offset)
            offset += 1
        else:
            nums.append(node.data)

    return linked_list


def remove_duplicates2(linked_list: LinkedList):
    nums = []

    for i, node in enumerate(linked_list):
        if node.data in nums:
            prev.next = node.next
        else:
            nums.append(node.data)
            prev = node

    return linked_list


def main():
    linked_list = LinkedList()
    for i in range(10):
        linked_list.append(randint(0, 5))

    print(linked_list)

    print(remove_duplicates1(linked_list))
    print(remove_duplicates2(linked_list))


if __name__ == '__main__':
    main()
