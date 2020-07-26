class TreeNode(object):
    def __init__(self, value):
        self.value = value
        self.left_child = None
        self.right_child = None

    def __str__(self):
        return str(self.value)

    def __repr__(self):
        return str(self.value)


class Tree():
    def __init__(self, array):
        length = len(array)
        root = TreeNode(array[length / 2])

        root.left_child, root.right_child = self.get_child(array[:length / 2]), self.get_child(array[length / 2  + 1:])
        self.root = root

    def get_child(self, array):
        length = len(array)
        node = TreeNode(array[length / 2])

        if length == 1:
            return node

        if length == 2:
            node.left_child = self.get_child(array[:length / 2])
        else:
            node.left_child, node.right_child = self.get_child(array[:length / 2]), self.get_child(array[length / 2  + 1:])

        return node


class LinkedListTreeNode(object):
    def __init__(self):
        self.node = None
        self.next = None


class LinkedListOfTreeNodes(object):
    def __init__(self):
        self.head = None
        self.tail = None
