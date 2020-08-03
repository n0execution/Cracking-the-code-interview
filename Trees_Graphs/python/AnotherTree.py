from random import randint


class AnotherTreeNode(object):
    def __init__(self, value):
        self.value = value
        self.left_child = None
        self.right_child = None

    def __str__(self):
        return str(self.value)

    def __repr__(self):
        return str(self.value)


class AnotherTree(object):
    def __init__(self):
        self.size = 0
        self.index = 0
        self.root = None

    def insert_in_order(self, value):
        if not self.size:
            self.root = AnotherTreeNode(value)
        else:
            self.insert(self.root, value)

        self.size += 1

    def insert(self, node, value):
        if value < node.value:
            if node.left_child:
                self.insert(node.left_child, value)
            else:
                node.left_child = AnotherTreeNode(value)
        elif value >= node.value:
            if node.right_child:
                self.insert(node.right_child, value)
            else:
                node.right_child = AnotherTreeNode(value)

    def find(self, node, value):
        if node.value == value:
            return node

        if node.left_child and node.right_child:
            left_node = self.find(node.left_child, value)
            right_node = self.find(node.right_child, value)

            if left_node:
                return left_node
            elif right_node:
                return right_node
        elif node.left_child:
            return self.find(node.left_child, value)
        elif node.right_child:
            return self.find(node.right_child, value)

        return None

    def get_node(self, node, num):
        if num == self.index:
            self.index = 0
            return node
        else:
            self.index += 1

            if not node.left_child and not node.right_child:
                return None
            elif node.left_child and node.right_child:
                left_node = self.get_node(node.left_child, num)
                right_node = self.get_node(node.right_child, num)

                if left_node:
                    return left_node
                elif right_node:
                    return right_node
            elif node.left_child:
                return self.get_node(node.left_child, num)
            elif node.right_child:
                return self.get_node(node.right_child, num)

        return None

    def get_random_node(self):
        num = randint(0, self.size - 1)
        print(num)

        return self.get_node(self.root, num)
