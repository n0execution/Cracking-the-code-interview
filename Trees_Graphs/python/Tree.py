from LinkedList import *


class TreeNode(object):
    def __init__(self, value, parent=None):
        self.value = value
        self.left_child = None
        self.right_child = None
        self.parent = parent

    def __str__(self):
        return str(self.value)

    def __repr__(self):
        return str(self.value)


class Tree():
    def __init__(self, array):
        length = len(array)
        root = TreeNode(array[int(length / 2)])

        root.left_child, root.right_child = self.get_child(array[:int(length / 2)], root), self.get_child(array[int(length / 2)  + 1:], root)
        self.root = root

    def get_child(self, array, parent):
        length = len(array)
        node = TreeNode(array[int(length / 2)], parent)

        if length == 1:
            return node

        if length == 2:
            node.left_child = self.get_child(array[:int(length / 2)], node)
        else:
            node.left_child, node.right_child = self.get_child(array[:int(length / 2)], node), self.get_child(array[int(length / 2)  + 1:], node)

        return node

    def write_depth(self, node, depth_num, depths):
        depth_list = LinkedList()

        if depth_num == len(depths):
            depth_list.append(node)
            depths.append(depth_list)
        else:
            depth_list = depths[depth_num]
            depth_list.append(node)
            depths[depth_num] = depth_list

        if node.left_child:
            self.write_depth(node.left_child, depth_num + 1, depths)

        if node.right_child:
            self.write_depth(node.right_child, depth_num + 1, depths)

    def print_tree(self):
        depths = []
        depth_list = LinkedList()

        depth_list.append(self.root)

        depths.append(depth_list)

        self.write_depth(self.root.left_child, 1, depths)
        self.write_depth(self.root.right_child, 1, depths)

        for depth in depths:
            print(depth)


class LinkedListTreeNode(object):
    def __init__(self):
        self.node = None
        self.next = None


class LinkedListOfTreeNodes(object):
    def __init__(self):
        self.head = None
        self.tail = None
