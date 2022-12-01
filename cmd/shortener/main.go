package main

import (
	"github.com/fkocharli/easyurl/internal/app/handlers"
	"github.com/fkocharli/easyurl/internal/app/server"
)

func main() {
	h := handlers.NewRouteHandler()
	m := server.NewMux(h)
	s := server.NewServer(m)
	for {
		s.Run()
	}
}
