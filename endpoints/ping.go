package endpoints

import (
	"github.com/valyala/fasthttp"
)

// Ping handles /ping
func Ping(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/plain")
	ctx.SetBody([]byte("PONG"))
}
