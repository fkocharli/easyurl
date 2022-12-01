package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/fkocharli/easyurl/internal/app/handlers"
)

type Serv struct {
	Server *http.Server
}

func NewServer(m *http.ServeMux) *Serv {
	return &Serv{
		Server: &http.Server{
			Handler: m,
			Addr:    "127.0.0.1:8080",
		},
	}
}

func NewMux(r []handlers.RouteHandler) *http.ServeMux {

	mux := http.NewServeMux()
	for _, v := range r {
		mux.Handle(v.Route, v.Handler)
	}
	return mux
}

func (s *Serv) Run() {

	group := &sync.WaitGroup{}
	group.Add(1)
	go func() {
		defer group.Done()
		if err := s.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	group.Wait()
}
