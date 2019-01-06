package endpoints

import (
	"encoding/json"

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

	g, err := state.AddGame(u, dto.Width, dto.Height, dto.Mines)
	if err != nil {
		return err
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBody([]byte(g.JSON()))

	return nil
}

// FlagCell handles POST /games/:game_id/cells/:cell_id/flag
func FlagCell(ctx *fasthttp.RequestCtx, u *user.User, state *state.State) error {
	return nil
}

// UnflagCell handles DELETE /games/:game_id/cells/:cell_id/flag
func UnflagCell(ctx *fasthttp.RequestCtx, u *user.User, state *state.State) error {
	return nil
}
