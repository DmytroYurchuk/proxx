import random
from typing import List

class ProxxGame:
    def __init__(self):
        """Initialize the Proxx game."""
        self.width: int = 0
        self.height: int = 0
        self.num_holes: int = 0
        self.board: List[List[int]] = []
        self.visible_board: List[List[str]] = []

    def get_board_size(self) -> None:
        """Prompt the player to enter the width and height of the board."""
        while True:
            try:
                self.width = int(input("Enter the width of the board: "))
                self.height = int(input("Enter the height of the board: "))
                break
            except ValueError:
                print("Invalid input. Please enter a valid integer.")

    def get_num_holes(self) -> None:
        """Prompt the player to enter the number of holes."""
        while True:
            try:
                self.num_holes = int(input("Enter the number of holes: "))
                if self.num_holes >= self.width * self.height:
                    print("Invalid input. Number of holes cannot exceed the total number of cells.")
                else:
                    break
            except ValueError:
                print("Invalid input. Please enter a valid integer.")

    def create_board(self) -> None:
        """Create the game board and the visible board."""
        self.board = [[0 for _ in range(self.width)] for _ in range(self.height)]
        self.visible_board = [["-" for _ in range(self.width)] for _ in range(self.height)]

    def place_holes(self) -> None:
        """Randomly place the holes on the game board."""
        locations = random.sample(range(self.width * self.height), self.num_holes)
        for location in locations:
            row = location // self.width
            col = location % self.width
            self.board[row][col] = -1

    def calculate_numbers(self) -> None:
        """Calculate the numbers indicating the neighboring holes for each cell."""
        for row in range(self.height):
            for col in range(self.width):
                if self.board[row][col] != -1:
                    count = 0
                    for r in range(max(0, row - 1), min(row + 2, self.height)):
                        for c in range(max(0, col - 1), min(col + 2, self.width)):
                            if self.board[r][c] == -1:
                                count += 1
                    self.board[row][col] = count

    def print_board(self) -> None:
        """Print the visible board."""
        for row in self.visible_board:
            print(" ".join(str(cell) for cell in row))

    def reveal_cell(self, row: int, col: int) -> bool:
        """
        Reveal the selected cell.

        Args:
            row (int): The row index of the cell.
            col (int): The column index of the cell.

        Returns:
            bool: True if the cell is successfully revealed, False if it contains a hole.

        """
        if self.board[row][col] == -1:
            self.visible_board[row][col] = "H"
            return False
        
        self.visible_board[row][col] = str(self.board[row][col])
        
        if self.board[row][col] == 0:
            for r in range(max(0, row - 1), min(row + 2, self.height)):
                for c in range(max(0, col - 1), min(col + 2, self.width)):
                    if self.visible_board[r][c] == "-":
                        self.reveal_cell(r, c)
        return True

    def check_win(self) -> bool:
        """
        Check if the player has won the game.

        Returns:
            bool: True if the player has won, False otherwise.

        """
        for row in range(self.height):
            for col in range(self.width):
                if self.board[row][col] != -1 and self.visible_board[row][col] == "-":
                    return False
        return True

    def play(self) -> None:
        """Play the Proxx game."""
        self.get_board_size()
        self.get_num_holes()
        self.create_board()
        self.place_holes()
        self.calculate_numbers()

        print("Initial board:")
        self.print_board()

        while True:
            try:
                row = int(input("Enter row (0 to {}): ".format(self.height - 1)))
                col = int(input("Enter column (0 to {}): ".format(self.width - 1)))

                if row < 0 or row >= self.height or col < 0 or col >= self.width:
                    print("Invalid move. Try again.")
                    continue

                if not self.reveal_cell(row, col):
                    print("Game over! You hit a hole.")
                    break

                if self.check_win():
                    print("Congratulations! You won the game!")
                    break

                self.print_board()
            except ValueError:
                print("Invalid input. Please enter a valid integer.")

        print("Final board:")
        self.print_board()