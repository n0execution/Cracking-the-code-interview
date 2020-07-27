from Tree import *


def write_to_array(array, node):
    if node.left_child is not None:
        write_to_array(array, node.left_child)

    array.append(node)

    if node.right_child is not None:
        write_to_array(array, node.right_child)


def validate_BST(node):
    array = []
    write_to_array(array, node.left_child)
    array.append(node)
    write_to_array(array, node.right_child)

    for i in range(len(array) - 1):
        if array[i].value > array[i + 1].value:
            return False
        elif array[i].value == array[i + 1].value and node.right_child is not None:
            return False

    return True


def main():
    array = list(range(7))
    array = [1, 2, 4, 4, 5, 6, 7, 8]
    tree = Tree(array)
    tree.print_tree()

    print(validate_BST(tree.root))


if __name__ == '__main__':
    main()
