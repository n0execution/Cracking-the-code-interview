from Tree import *


def get_next(node):
    if node.left_child is not None:
        return get_next(node.left_child)

    return node


def successor(node):
    if node.right_child is None:
        parent = node.parent
        while parent.value < node.value:
            parent = parent.parent

            if parent is None:
                return None

        return parent

    return get_next(node.right_child)


def main():
    array = list(range(7))
    tree = Tree(array)
    tree.print_tree()

    print(successor(tree.root.right_child.right_child))


if __name__ == '__main__':
    main()
