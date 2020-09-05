from exceptions import NoPathException


class Cell(object):
    def __init__(self, value, row, column, no_entry=False):
        self.value = value
        self.row = row
        self.column = column
        self.no_entry = no_entry

    def __str__(self):
        return f"{self.value} {self.no_entry}"

    def __repr__(self):
        return f"{self.value} {self.no_entry}"


class Grid(object):
    def __init__(self, rows, cols, no_entries=[]):
        self.cells = []
        self.rows = rows
        self.cols = cols
        self.no_entries = no_entries
        self.paths = {}

    def initialize_grid(self):
        index = 1
        for i in range(self.rows):
            row = []
            for j in range(self.cols):
                cell = Cell(value=index,
                            row=i,
                            column=j,
                            no_entry=index in self.no_entries)
                row.append(cell)

                index += 1
            self.cells.append(row)

        print(self.cells)

    def get_cell(self, row, col):
        if row < 0 or col < 0:
            return None

        cell = self.cells[row][col]

        if cell.no_entry:
            return None

        return cell

    def find_path(self, from_cell, to_cell):
        print(from_cell)
        if not from_cell:
            return []

        if from_cell in self.paths.keys():
            print('l')
            return self.paths[from_cell]

        if from_cell == to_cell:
            return [from_cell]

        left_cell = self.get_cell(from_cell.row, from_cell.column-1)
        upper_cell = self.get_cell(from_cell.row-1, from_cell.column)

        if not left_cell and not upper_cell:
            return []

        left_path = self.find_path(left_cell, to_cell)
        upper_path = self.find_path(upper_cell, to_cell)

        if left_path:
            self.paths[from_cell] = [*left_path, from_cell]
            return self.paths[from_cell]

        if upper_path:
            self.paths[from_cell] = [*upper_path, from_cell]
            return self.paths[from_cell]

        raise NoPathException()

    def find(self, from_cell, to_cell):
        return self.find_path(to_cell, from_cell)


def main():
    rows = 4
    cols = 5

    grid = Grid(rows, cols, no_entries=[7, 14, 17, 19])
    # grid = Grid(rows, cols, no_entries=[3, 6, 9, 12, 15, 18])
    # grid = Grid(rows, cols, no_entries=[8, 14, 18])
    grid.initialize_grid()

    path = grid.find(from_cell=grid.cells[0][0],
                     to_cell=grid.cells[rows-1][cols-1])

    print(path)

if __name__ == '__main__':
    main()
