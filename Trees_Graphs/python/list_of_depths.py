from Tree import *
from LinkedList import *


def write_depth(node, depth_num, depths):
    depth_list = LinkedList()

    if depth_num == len(depths):
        depth_list.append(node)
        depths.append(depth_list)
    else:
        depth_list = depths[depth_num]
        depth_list.append(node)
        depths[depth_num] = depth_list

    if node.left_child:
        write_depth(node.left_child, depth_num + 1, depths)

    if node.right_child:
        write_depth(node.right_child, depth_num + 1, depths)


def main():
    depths = []
    depth_list = LinkedList()
    array = list(range(7))
    tree = Tree(array)

    depth_list.append(tree.root)
    depths.append(depth_list)

    write_depth(tree.root.left_child, 1, depths)
    write_depth(tree.root.right_child, 1, depths)

    for depth in depths:
        print(depth)


if __name__ == '__main__':
    main()
