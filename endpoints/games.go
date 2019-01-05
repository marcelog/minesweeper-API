package endpoints

import (
	"github.com/valyala/fasthttp"

	"github.com/marcelog/minesweeper-API/state"
	"github.com/marcelog/minesweeper-API/user"
)

// CreateGame handles POST /games
func CreateGame(ctx *fasthttp.RequestCtx, u *user.User, state *state.State) {
	g := state.AddGame(u)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBody([]byte(g.JSON()))
}
