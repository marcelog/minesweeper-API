package game

import (
	"encoding/json"
	"math/rand"
	"time"
)

const (
	// GameStarted the game is started
	GameStarted = "started"

	// GameLost terminal state, the player lost
	GameLost = "lost"

	// GameWon terminal state, the player won
	GameWon = "won"
)

const (
	// CellUnvisited A blank cell
	CellUnvisited = 0

	// CellAdjMines0 has 0 adjacent mines
	CellAdjMines0 = 1

	// CellAdjMines1 has 1 adjacent mines
	CellAdjMines1 = 2

	// CellAdjMines2 has 2 adjacent mines
	CellAdjMines2 = 3

	// CellAdjMines3 has 3 adjacent mines
	CellAdjMines3 = 4

	// CellAdjMines4 has 4 adjacent mines
	CellAdjMines4 = 5

	// CellAdjMines5 has 5 adjacent mines
	CellAdjMines5 = 6

	// CellAdjMines6 has 6 adjacent mines
	CellAdjMines6 = 7

	// CellAdjMines7 has 7 adjacent mines
	CellAdjMines7 = 8

	// CellAdjMines8 has 8 adjacent mines
	CellAdjMines8 = 9

	// CellFlagged A cell that has been flagged by the user
	CellFlagged = 10
)

// Game represents a game.
type Game struct {
	ID         int          `json:"id"`
	OwnerID    int          `json:"-"`
	Width      int          `json:"width"`
	Height     int          `json:"height"`
	TotalMines int          `json:"mines"`
	State      string       `json:"state"`
	Board      []int        `json:"board"`
	Mines      map[int]bool `json:"-"` // Mines location, key is an int, cell number.
}

// New creates a new game.
func New(id int, ownerID int, width int, height int, mines int) *Game {
	totalCells := width * height

	g := &Game{
		ID:         id,
		OwnerID:    ownerID,
		Width:      width,
		Height:     height,
		TotalMines: mines,
		Board:      make([]int, totalCells),
		Mines:      map[int]bool{},
		State:      GameStarted,
	}

	// Initialize board
	for i := 0; i < totalCells; i++ {
		g.Board[i] = CellUnvisited
	}

	// Choose where to put mines
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate N mines. Save each cell number with a mine in our Mines map.
	for i := 0; i < mines; i++ {
		// We loop here to avoid duplicated cell numbers.
		for true {
			cell := r.Intn(totalCells)
			if _, ok := g.Mines[cell]; !ok {
				g.Mines[cell] = true
				break
			}
		}
	}
	return g
}

// JSON serializes this user as a json string.
func (g *Game) JSON() string {
	j, _ := json.Marshal(g)
	return string(j)
}
