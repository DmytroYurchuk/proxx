package proxx

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ProxxGame struct {
	Width        int
	Height       int
	NumHoles     int
	Board        [][]int
	VisibleBoard [][]string
}

func (g *ProxxGame) getBoardSize() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the width of the board: ")
		scanner.Scan()
		width, err := strconv.Atoi(scanner.Text())
		if err == nil {
			g.Width = width
			break
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}

	for {
		fmt.Print("Enter the height of the board: ")
		scanner.Scan()
		height, err := strconv.Atoi(scanner.Text())
		if err == nil {
			g.Height = height
			break
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}
}

func (g *ProxxGame) getNumHoles() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the number of holes: ")
		scanner.Scan()
		numHoles, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if numHoles >= g.Width*g.Height {
				fmt.Println("Invalid input. Number of holes cannot exceed the total number of cells.")
			} else {
				g.NumHoles = numHoles
				break
			}
		} else {
			fmt.Println("Invalid input. Please enter a valid integer.")
		}
	}
}

func (g *ProxxGame) createBoard() {
	g.Board = make([][]int, g.Height)
	g.VisibleBoard = make([][]string, g.Height)

	for i := 0; i < g.Height; i++ {
		g.Board[i] = make([]int, g.Width)
		g.VisibleBoard[i] = make([]string, g.Width)
		for j := 0; j < g.Width; j++ {
			g.VisibleBoard[i][j] = "-"
		}
	}
}

func (g *ProxxGame) placeHoles() {
	rand.Seed(time.Now().UnixNano())

	locations := rand.Perm(g.Width * g.Height)[:g.NumHoles]
	for _, location := range locations {
		row := location / g.Width
		col := location % g.Width
		g.Board[row][col] = -1
	}
}

func (g *ProxxGame) calculateNumbers() {
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			if g.Board[row][col] != -1 {
				count := 0
				for r := int(math.Max(0, float64(row)-1)); r <= int(math.Min(float64(row)+1, float64(g.Height)-1)); r++ {
					for c := int(math.Max(0, float64(col)-1)); c <= int(math.Min(float64(col)+1, float64(g.Width)-1)); c++ {
						if g.Board[r][c] == -1 {
							count++
						}
					}
				}
				g.Board[row][col] = count
			}
		}
	}
}

func (g *ProxxGame) printBoard() {
	for _, row := range g.VisibleBoard {
		fmt.Println(strings.Join(row, " "))
	}
}

func (g *ProxxGame) revealCell(row, col int) bool {
	if g.Board[row][col] == -1 {
		g.VisibleBoard[row][col] = "H"
		return false
	}

	g.VisibleBoard[row][col] = strconv.Itoa(g.Board[row][col])

	if g.Board[row][col] == 0 {
		for r := int(math.Max(0, float64(row)-1)); r <= int(math.Min(float64(row)+1, float64(g.Height-1))); r++ {
			for c := int(math.Max(0, float64(col)-1)); c <= int(math.Min(float64(col)+1, float64(g.Width-1))); c++ {
				if g.VisibleBoard[r][c] == "-" {
					g.revealCell(r, c)
				}
			}
		}
	}
	return true
}

func (g *ProxxGame) checkWin() bool {
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			if g.Board[row][col] != -1 && g.VisibleBoard[row][col] == "-" {
				return false
			}
		}
	}
	return true
}

func (g *ProxxGame) Play() {
	g.getBoardSize()
	g.getNumHoles()
	g.createBoard()
	g.placeHoles()
	g.calculateNumbers()

	fmt.Println("Initial board:")
	g.printBoard()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter row (0 to ", g.Height-1, "): ")
		scanner.Scan()
		row, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		fmt.Print("Enter column (0 to ", g.Width-1, "): ")
		scanner.Scan()
		col, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		if row < 0 || row >= g.Height || col < 0 || col >= g.Width {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		if !g.revealCell(row, col) {
			fmt.Println("Game over! You hit a hole.")
			break
		}

		if g.checkWin() {
			fmt.Println("Congratulations! You won the game!")
			break
		}

		fmt.Println("--------------------")
		g.printBoard()
	}

	fmt.Println("Final board:")
	g.printBoard()
}
