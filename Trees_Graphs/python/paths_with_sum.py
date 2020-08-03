from Tree import Tree


def count_nums(node, sum, value, passed):
    counter = 0

    if node.left_child:
        if sum + node.left_child.value == value and not passed.get(node.left_child):
            passed[node.left_child] = True
            counter += 1

        counter += count_nums(node.left_child, node.left_child.value, value, passed)
        counter += count_nums(node.left_child, sum + node.left_child.value, value, passed)

    if node.right_child:
        if sum + node.right_child.value == value and not passed.get(node.right_child):
            passed[node.right_child] = True
            counter += 1

        counter += count_nums(node.right_child, node.right_child.value, value, passed)
        counter += count_nums(node.right_child, sum + node.right_child.value, value, passed)

    return counter


def get_nums(tree, value):
    passed = {}

    return count_nums(tree.root, tree.root.value, value, passed)


def main():
    array = list(range(7))
    tree = Tree(array)
    tree.print_tree()

    print(get_nums(tree, 6))


if __name__ == '__main__':
    main()
