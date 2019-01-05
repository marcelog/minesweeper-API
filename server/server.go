package server

import (
	"fmt"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/marcelog/minesweeper-API/endpoints"
	"github.com/marcelog/minesweeper-API/state"
)

// IServer is a generic interface to servers.
type IServer interface {
	Run() error
	Stop()
}

// Args is used to create a new Server.
type Args struct {
	Address string
	Port    int
}

// Server is our current server implementation.
type Server struct {
	args   *Args
	server *fasthttp.Server
	state  *state.State
}

// New returns a new *Server.
func New(args *Args) *Server {
	return &Server{
		args:  args,
		state: state.New(),
	}
}

// Run starts the server.
func (s *Server) Run() error {
	// ListenAndServer sadly blocks forever, so this trick allow us to return
	// a success condition by trusting that a listen operation will not take more
	// than 100ms to complete (99.9999% a "sure bet" :shrug:). The trick is as
	// follows:
	//  * Create a channel
	//  * Start the server in a go routine so it doesn't block the main thread
	//  * On error, the listen operation should return immediatly, and we send
	//    that error through the channel.
	//  * In the main thread, we "select" on that channel OR by a timeout.
	//  * If we receive something on the channel, it means an error.
	//  * If we hit the timeout, we _assume_ no errors are present and return
	//    success.
	c := make(chan error)
	go func() {
		router := s.createRoutes()
		s.server = &fasthttp.Server{
			Handler:     router.Handler,
			ReadTimeout: (100 * time.Millisecond),
		}
		addr := fmt.Sprintf("%s:%d", s.args.Address, s.args.Port)
		fmt.Println("Starting a server in", addr)
		if err := s.server.ListenAndServe(addr); err != nil {
			c <- err
		}
	}()

	select {
	case err := <-c:
		return err
	case <-time.After(100 * time.Millisecond):
		return nil
	}
}

// Stop closes the listening socket and shuts down the server.
func (s *Server) Stop() {
	fmt.Println("Shutting down the server")
	s.server.Shutdown()
}

func (s *Server) createRoutes() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.GET("/ping", s.createHandler(endpoints.Ping))
	router.POST("/users", s.createHandler(endpoints.CreateUser))
	return router
}

func (s *Server) createHandler(realHandler func(*fasthttp.RequestCtx, *state.State)) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		realHandler(ctx, s.state)
	}
}
