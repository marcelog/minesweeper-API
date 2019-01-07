package game

import (
	"testing"
)

func TestErrorWithInvalidWidth(t *testing.T) {
	g, err := New(1, 1, 0, 8, 1)
	if g != nil {
		t.Fatal("Expected nil game:", g)
	}
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "width too low" {
		t.Fatal("Unexpected error:", err.Error())
	}

	g, err = New(1, 1, 65, 8, 1)
	if g != nil {
		t.Fatal("Expected nil game:", g)
	}
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "width too high" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorWithInvalidHeight(t *testing.T) {
	g, err := New(1, 1, 8, 0, 1)
	if g != nil {
		t.Fatal("Expected nil game:", g)
	}
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "height too low" {
		t.Fatal("Unexpected error:", err.Error())
	}

	g, err = New(1, 1, 8, 65, 1)
	if g != nil {
		t.Fatal("Expected nil game:", g)
	}
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "height too high" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorWithInvalidNumberOfMines(t *testing.T) {
	g, err := New(1, 1, 8, 8, 0)
	if g != nil {
		t.Fatal("Expected nil game:", g)
	}
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "mine number too low" {
		t.Fatal("Unexpected error:", err.Error())
	}

	g, err = New(1, 1, 8, 8, 33)
	if g != nil {
		t.Fatal("Expected nil game:", g)
	}
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "mine number too high" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorFlaggingInvalidCell(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)

	err := g.Flag(-1)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "invalid cell" {
		t.Fatal("Unexpected error:", err.Error())
	}

	err = g.Flag(65)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "invalid cell" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorFlaggingInvalidGameState(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)
	g.State = GameLost

	err := g.Flag(5)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "game has finished" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorFlaggingInvalidCellState(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)
	g.Board[5] = CellAdjMines0
	err := g.Flag(5)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "cell in invalid state" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestCreatesRightNumberOfMines(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 5)
	if len(g.Mines) != 5 {
		t.Fatal("Unexpected number of mines:", len(g.Mines))
	}
}

func TestCanFlagCell(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)

	err := g.Flag(5)
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}
	if g.Board[5] != CellFlagged {
		t.Fatal("Cell wasn't flagged:", g.Board[5])
	}

	// Can reflag as well.
	err = g.Flag(5)
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorUnflaggingInvalidCell(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)

	err := g.Unflag(-1)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "invalid cell" {
		t.Fatal("Unexpected error:", err.Error())
	}

	err = g.Unflag(65)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "invalid cell" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorUnflaggingInvalidGameState(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)
	g.State = GameLost

	err := g.Unflag(5)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "game has finished" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestErrorUnflaggingInvalidCellState(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)
	g.Board[5] = CellAdjMines0
	err := g.Unflag(5)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() != "cell in invalid state" {
		t.Fatal("Unexpected error:", err.Error())
	}
}

func TestCanUnflagCell(t *testing.T) {
	g, _ := New(1, 1, 8, 8, 1)
	g.Board[5] = CellFlagged

	err := g.Unflag(5)
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}
	if g.Board[5] != CellUnvisited {
		t.Fatal("Cell wasn't unflagged:", g.Board[5])
	}

	// Can re-unflag as well.
	err = g.Unflag(5)
	if err != nil {
		t.Fatal("Unexpected error:", err.Error())
	}
}
