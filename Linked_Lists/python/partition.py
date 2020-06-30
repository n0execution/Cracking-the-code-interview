from LinkedList import  LinkedList, Node
from random import randint


def perform_partition(linked_list: LinkedList, partition: int):
    for node in linked_list:
        if node.data < partition and node != linked_list.head:
            linked_list.insert_at_index(0, node.data)

            if not node.next:
                linked_list.delete_last()
            else:
                delete(node)


def delete(node):
    node.data = node.next.data
    node.next = node.next.next


def main():
    linked_list = LinkedList()
    for i in range(10):
        linked_list.append(randint(0, 9))

    print(linked_list)
    perform_partition(linked_list, 6)
    print(linked_list)


if __name__ == '__main__':
    main()
