from Tree import *


def get_height(node):
    if node.left_child is None and node.right_child is None:
        return True, 1

    if node.left_child is None:
        left = 1
    else:
        balanced_left, left = get_height(node.left_child)

    if node.right_child is None:
        right = 1
    else:
        balanced_right, right = get_height(node.left_child)

    if balanced_left and balanced_right:
        return True, max(left + 1, right + 1)


def is_balanced(tree):
    root = tree.root

    balanced_left = get_height(root.left_child)[0]
    balanced_right = get_height(root.right_child)[0]

    if not balanced_left or not balanced_right:
        return False

    return True


def main():
    array = list(range(7))
    tree = Tree(array)
    tree.print_tree()

    print(is_balanced(tree))


if __name__ == '__main__':
    main()
