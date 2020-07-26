from Tree import TreeNode


def get_child(array):
    length = len(array)
    node = TreeNode(array[length / 2])

    if length == 1:
        return node

    if length == 2:
        node.left_child = get_child(array[:length / 2])
    else:
        node.left_child, node.right_child = get_child(array[:length / 2]), get_child(array[length / 2  + 1:])

    return node

def create_tree(array):
    length = len(array)
    root = TreeNode(array[length / 2])

    root.left_child, root.right_child = get_child(array[:length / 2]), get_child(array[length / 2  + 1:])
    return root


def main():
    array = list(range(7))
    print(array)
    root = create_tree(array)
    print(root.value)
    print(root.left_child.value, root.right_child.value)
    print(root.left_child.left_child.value, root.left_child.right_child.value, root.right_child.left_child.value, root.right_child.right_child.value)


if __name__ == '__main__':
    main()
