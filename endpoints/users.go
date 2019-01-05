package endpoints

import (
	"github.com/valyala/fasthttp"

	"github.com/marcelog/minesweeper-API/state"
)

// CreateUser handles POST /users
func CreateUser(ctx *fasthttp.RequestCtx, state *state.State) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte("PONG"))
}
