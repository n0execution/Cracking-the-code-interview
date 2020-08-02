from Tree import *


def get_sequences(node):
    sequences, list_of_sequences = [], []

    if node.left_child is None and node.right_child is None:
        sequences.append(node)
        list_of_sequences.append(sequences)
    elif node.left_child is None:
        right_sequences = get_sequences(node.right_child)

        for sequence in right_sequences:
            s = [node, *sequence]
            list_of_sequences.append(s)
    elif node.right_child is None:
        left_sequences = get_sequences(node.left_child)

        for sequence in left_sequences:
            s = [node, *sequence]
            list_of_sequences.append(s)
    else:
        left_sequences = get_sequences(node.left_child)
        right_sequences = get_sequences(node.right_child)

        for sequence in left_sequences:
            begin = [node, *sequence]

            for sequence in right_sequences:
                s = [*begin, *sequence]
                list_of_sequences.append(s)

        for sequence in right_sequences:
            begin = [node, *sequence]

            for sequence in left_sequences:
                s = [*begin, *sequence]
                list_of_sequences.append(s)

    return list_of_sequences

def main():
    array = list(range(7))
    tree = Tree(array)
    tree.print_tree()

    print(get_sequences(tree.root))


if __name__ == '__main__':
    main()
