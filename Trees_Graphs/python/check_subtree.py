from Tree import Tree


def is_subtree(node1, node2):
    if node1 == node2:
        return True

    if node1.left_child is None and node1.right_child is None:
        return False
    elif node1.left_child is None:
        return is_subtree(node1.right_child, node2)
    elif node1.right_child is None:
        return is_subtree(node1.left_child, node2)
    else:
        return is_subtree(node1.left_child, node2) or is_subtree(node1.right_child, node2)

    return False


def main():
    array = list(range(7))
    tree = Tree(array)
    tree.print_tree()

    print(is_subtree(tree.root, tree.root.left_child))


if __name__ == '__main__':
    main()
