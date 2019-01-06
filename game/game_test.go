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
