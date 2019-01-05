package endpoints

import (
	"github.com/valyala/fasthttp"

	"github.com/marcelog/minesweeper-API/state"
)

// Ping handles GET /ping
func Ping(ctx *fasthttp.RequestCtx, state *state.State) error {
	ctx.SetContentType("text/plain")
	ctx.SetBody([]byte("PONG"))
	return nil
}
