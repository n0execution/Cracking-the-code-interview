from Stack import Stack


class Tower(object):
    def __init__(self, number):
        self.disks = Stack()
        self.number = number

    def initialize_disks(self, disks):
        for disk in sorted(disks, reverse=True):
            self.disks.push(disk)

    def move_disks(self, n, dest, buffer):
        if n > 0:
            self.move_disks(n - 1, buffer, dest)
            dest.disks.push(self.disks.pop())
            buffer.move_disks(n - 1, dest, self)


def main():
    towers = [Tower(1), Tower(2), Tower(3)]

    disks = list(range(1, 6))
    towers[0].initialize_disks(disks)

    print(towers[0].disks)
    print(towers[0].disks.peek())

    towers[0].move_disks(towers[0].disks.size, towers[2], towers[1])

    for tower in towers:
        print(tower.number, tower.disks)


if __name__ == '__main__':
    main()
