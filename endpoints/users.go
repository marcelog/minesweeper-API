package endpoints

import (
	"github.com/valyala/fasthttp"

	"github.com/marcelog/minesweeper-API/state"
	"github.com/marcelog/minesweeper-API/user"
)

// CreateUser handles POST /users
func CreateUser(ctx *fasthttp.RequestCtx, state *state.State) {
	u := user.New()
	state.AddUser(u)
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte(u.JSON()))
}
