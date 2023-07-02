package proxx

import (
	"os"
	"io"
	"io/ioutil"
	"strings"
	"testing"
	"reflect"
)

func helperStdin(t *testing.T, f func(), input ...string) string {
	if len(input) > 0 {
		reader := strings.NewReader(input[0])
		
		// Create a temporary file
		tempFile, err := ioutil.TempFile("", "stdin")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		// Write the input to the temporary file
		_, err = io.Copy(tempFile, reader)
		if err != nil {
			t.Fatalf("Failed to write input to temporary file: %v", err)
		}

		// Open the temporary file for reading
		file, err := os.Open(tempFile.Name())
		if err != nil {
			t.Fatalf("Failed to open temporary file: %v", err)
		}
		defer file.Close()

		// Replace os.Stdin temporarily
		originalStdin := os.Stdin
		defer func() { os.Stdin = originalStdin }()
		os.Stdin = file
	}

	// Capture the standard output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function that prints to stdout
	f()

	// Restore the standard output
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout
	return strings.TrimSpace(string(out))
}

func TestProxxGame_RevealCell_HitHole(t *testing.T) {
	game := ProxxGame{
		Width:         3,
		Height:        3,
		Board:         [][]int{{0, -1, 0}, {0, 0, 0}, {0, 0, 0}},
		VisibleBoard:  [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}},
	}

	result := game.revealCell(0, 1)

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}

	expectedVisibleBoard := [][]string{{"-", "H", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	if !reflect.DeepEqual(game.VisibleBoard, expectedVisibleBoard) {
		t.Errorf("Visible board is not as expected.\nExpected: %v\nActual: %v", expectedVisibleBoard, game.VisibleBoard)
	}
}

func TestProxxGame_RevealCell_EmptyCell(t *testing.T) {
	game := ProxxGame{
		Width:         3,
		Height:        3,
		Board:         [][]int{{0, -1, 0}, {0, 0, 0}, {0, 0, 0}},
		VisibleBoard:  [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}},
	}

	game.calculateNumbers()

	result := game.revealCell(0, 0)

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}

	expectedVisibleBoard := [][]string{{"1", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	if !reflect.DeepEqual(game.VisibleBoard, expectedVisibleBoard) {
		t.Errorf("Visible board is not as expected.\nExpected: %v\nActual: %v", expectedVisibleBoard, game.VisibleBoard)
	}
}

func TestProxxGame_CheckWin_NoWin(t *testing.T) {
	game := ProxxGame{
		Width:         3,
		Height:        3,
		Board:         [][]int{{0, -1, 0}, {0, 0, 0}, {0, 0, 0}},
		VisibleBoard:  [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}},
	}

	result := game.checkWin()

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestProxxGame_CheckWin_Win(t *testing.T) {
	game := ProxxGame{
		Width:         3,
		Height:        3,
		Board:         [][]int{{0, -1, 0}, {0, 0, 0}, {0, 0, 0}},
		VisibleBoard:  [][]string{{"1", "H", "1"}, {"1", "1", "1"}, {"0", "0", "0"}},
	}

	result := game.checkWin()

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestProxxGame_PrintBoard(t *testing.T) {
	game := ProxxGame{
		VisibleBoard: [][]string{{"-", "1", "-"}, {"2", "-", "3"}, {"-", "4", "-"}},
	}
	expectedOutput := "- 1 -\n2 - 3\n- 4 -"

	output := helperStdin(t, game.printBoard)

	if output != expectedOutput {
		t.Errorf("Output is not as expected.\nExpected: %v\nActual: %v", expectedOutput, output)
	}
}

func TestProxxGame_GetBoardSize_ValidInput(t *testing.T) {
	game := ProxxGame{}
	input := "4\n5"
	expectedWidth := 4
	expectedHeight := 5

	helperStdin(t, game.getBoardSize, input)

	if game.Width != expectedWidth {
		t.Errorf("Expected width %v, got %v", expectedWidth, game.Width)
	}

	if game.Height != expectedHeight {
		t.Errorf("Expected height %v, got %v", expectedHeight, game.Height)
	}
}

func TestProxxGame_GetBoardSize_InvalidInput(t *testing.T) {
	game := ProxxGame{}
	input := "invalid\n3\n6"
	expectedWidth := 3
	expectedHeight := 6

	helperStdin(t, game.getBoardSize, input)

	if game.Width != expectedWidth {
		t.Errorf("Expected width %v, got %v", expectedWidth, game.Width)
	}

	if game.Height != expectedHeight {
		t.Errorf("Expected height %v, got %v", expectedHeight, game.Height)
	}
}

func TestProxxGame_GetNumHoles_ValidInput(t *testing.T) {
	game := ProxxGame{
		Width: 5,
		Height: 5,
	}
	input := "8"
	expectedNumHoles := 8

	helperStdin(t, game.getNumHoles, input)

	if game.NumHoles != expectedNumHoles {
		t.Errorf("Expected numHoles %v, got %v", expectedNumHoles, game.NumHoles)
	}
}

func TestProxxGame_GetNumHoles_InvalidInput(t *testing.T) {
	game := ProxxGame{
		Width: 5,
		Height: 5,
	}
	input := "invalid\n100\n12"
	expectedNumHoles := 12

	helperStdin(t, game.getNumHoles, input)

	if game.NumHoles != expectedNumHoles {
		t.Errorf("Expected numHoles %v, got %v", expectedNumHoles, game.NumHoles)
	}
}
