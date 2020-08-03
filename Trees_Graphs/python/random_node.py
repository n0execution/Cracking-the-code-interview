from AnotherTree import *


def main():
    tree = AnotherTree()

    tree.insert_in_order(5)
    tree.insert_in_order(2)
    tree.insert_in_order(10)
    tree.insert_in_order(15)
    tree.insert_in_order(11)
    tree.insert_in_order(1)

    print(tree.get_random_node())


if __name__ == '__main__':
    main()
