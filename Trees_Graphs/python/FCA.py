from Tree import *
from Stack import Stack


def get_first_common_ancestor(node1, node2):
    print(node1, node2)
    path1, path2 = Stack(), Stack()
    parent1, parent2 = node1, node2

    while parent1.parent is not None:
        if parent1.parent.left_child == parent1:
            path1.push("left")
        else:
            path1.push("right")

        parent1 = parent1.parent

    while parent2.parent is not None:
        if parent2.parent.left_child == parent2:
            path2.push("left")
        else:
            path2.push("right")

        parent2 = parent2.parent

    node = parent1

    while True:
        way1, way2 = path1.pop(), path2.pop()

        if way1 == way2:
            if way1 == "left":
                node = node.left_child
            else:
                node = node.right_child
        elif node1 == node or node2 == node:
            return node.parent
        else:
            break

    return node


def main():
    array = list(range(7))
    tree = Tree(array)
    tree.print_tree()

    print(get_first_common_ancestor(tree.root.left_child.left_child, tree.root.right_child))


if __name__ == '__main__':
    main()
