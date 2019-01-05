package endpoints

import (
	"github.com/valyala/fasthttp"

	"github.com/marcelog/minesweeper-API/state"
)

// CreateUser handles POST /users
func CreateUser(ctx *fasthttp.RequestCtx, state *state.State) {
	u := state.AddUser()
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBody([]byte(u.JSON()))
}
