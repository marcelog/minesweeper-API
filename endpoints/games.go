package endpoints

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"

	"github.com/marcelog/minesweeper-API/state"
	"github.com/marcelog/minesweeper-API/user"
)

// GameCreationDTO is what the user POSTs to create a game.
type GameCreationDTO struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Mines  int `json:"mines"`
}

// CreateGame handles POST /games
func CreateGame(ctx *fasthttp.RequestCtx, u *user.User, state *state.State) error {
	dto := &GameCreationDTO{}
	err := json.Unmarshal(ctx.PostBody(), dto)
	if err != nil {
		return err
	}

	// Arbitrary values.
	if dto.Width < 8 {
		return errors.New("width too low")
	}
	if dto.Width > 64 {
		return errors.New("width too high")
	}

	if dto.Height < 8 {
		return errors.New("height too low")
	}
	if dto.Height > 64 {
		return errors.New("height too high")
	}
	// At least 1 mine and less than 51% of cells with mines.
	if dto.Mines < 1 {
		return errors.New("mine number too low")
	}

	if dto.Mines > ((dto.Width * dto.Height) / 2) {
		return errors.New("mine number too high")
	}

	g := state.AddGame(u)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBody([]byte(g.JSON()))

	return nil
}
