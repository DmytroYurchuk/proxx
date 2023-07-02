import unittest
from io import StringIO
from unittest.mock import patch
from src.proxx import ProxxGame


class TestProxxGame(unittest.TestCase):
    def setUp(self):
        self.game = ProxxGame()

    def test_reveal_cell_hit_hole(self):
        self.game.width = 3
        self.game.height = 3
        self.game.board = [[0, -1, 0],
                            [0, 0, 0],
                            [0, 0, 0]]
        self.game.visible_board = [["-", "-", "-"],
                                    ["-", "-", "-"],
                                    ["-", "-", "-"]]

        result = self.game.reveal_cell(0, 1)

        self.assertFalse(result)
        self.assertEqual(self.game.visible_board, [["-", "H", "-"],
                                                    ["-", "-", "-"],
                                                    ["-", "-", "-"]])

    def test_one_reveal_cell_empty_cell(self):
        self.game.width = 3
        self.game.height = 3
        self.game.board = [[0, -1, 0],
                            [0, 0, 0],
                            [0, 0, 0]]
        self.game.visible_board = [["-", "-", "-"],
                                    ["-", "-", "-"],
                                    ["-", "-", "-"]]
        
        self.game.calculate_numbers()

        result = self.game.reveal_cell(0, 0)

        self.assertTrue(result)
        self.assertEqual(self.game.visible_board, [["1", "-", "-"],
                                                    ["-", "-", "-"],
                                                    ["-", "-", "-"]])
        
    def test_multiply_reveal_cell_empty_cell(self):
        self.game.width = 3
        self.game.height = 3
        self.game.board = [[-1, 0, 0],
                            [0, 0, 0],
                            [0, 0, 0]]
        self.game.visible_board = [["-", "-", "-"],
                                    ["-", "-", "-"],
                                    ["-", "-", "-"]]
        
        self.game.calculate_numbers()

        result = self.game.reveal_cell(2, 2)

        self.assertTrue(result)
        self.assertEqual(self.game.visible_board, [["-", "1", "0"],
                                                    ["1", "1", "0"],
                                                    ["0", "0", "0"]])
        
    def test_check_win_no_win(self):
        self.game.width = 3
        self.game.height = 3
        self.game.board = [[0, -1, 0],
                            [0, 0, 0],
                            [0, 0, 0]]
        self.game.visible_board = [["-", "-", "-"],
                                    ["-", "-", "-"],
                                    ["-", "-", "-"]]

        result = self.game.check_win()

        self.assertFalse(result)

    def test_check_win_win(self):
        self.game.width = 3
        self.game.height = 3
        self.game.board = [[0, -1, 0],
                            [0, 0, 0],
                            [0, 0, 0]]
        self.game.visible_board = [["1", "H", "1"],
                                    ["1", "1", "1"],
                                    ["0", "0", "0"]]

        result = self.game.check_win()

        self.assertTrue(result)

    @patch("sys.stdout", new_callable=StringIO)
    def test_print_board(self, mock_stdout):
        self.game.visible_board = [["-", "1", "-"],
                                    ["2", "-", "3"],
                                    ["-", "4", "-"]]

        self.game.print_board()

        expected_output = "- 1 -\n2 - 3\n- 4 -\n"
        self.assertEqual(mock_stdout.getvalue(), expected_output)

    def test_get_board_size_valid_input(self):
        with patch("builtins.input", side_effect=["4", "5"]):
            self.game.get_board_size()
        self.assertEqual(self.game.width, 4)
        self.assertEqual(self.game.height, 5)
    
    @patch("sys.stdout", new_callable=StringIO)
    def test_get_board_size_invalid_input(self, _):
        with patch("builtins.input", side_effect=["invalid", "3", "6"]):
            self.game.get_board_size()
        self.assertEqual(self.game.width, 3)
        self.assertEqual(self.game.height, 6)

    def test_get_num_holes_valid_input(self):
        self.game.width = 5
        self.game.height = 5
        with patch("builtins.input", side_effect=["8"]):
            self.game.get_num_holes()
        self.assertEqual(self.game.num_holes, 8)

    @patch("sys.stdout", new_callable=StringIO)
    def test_get_num_holes_invalid_input(self, _):
        self.game.width = 5
        self.game.height = 5
        with patch("builtins.input", side_effect=["invalid", "100", "12"]):
            self.game.get_num_holes()
        self.assertEqual(self.game.num_holes, 12)
