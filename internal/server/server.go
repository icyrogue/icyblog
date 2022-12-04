package server

import (
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/access"
	"github.com/go-ozzo/ozzo-routing/v2/fault"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	"github.com/go-ozzo/ozzo-routing/v2/slash"
)

type server struct {
	router  *routing.Router
	Options *Options
}

type Options struct {
	Port string
}

func New() *server {
	return &server{}
}

func (s *server) Init() {
	s.router = routing.New()
	s.router.Use(
		access.Logger(log.Printf),
		slash.Remover(http.StatusMovedPermanently),
		fault.Recovery(log.Printf),
	)
	//main := s.router.Group("/main")
	s.router.Get("/", file.Content("./index.html"))
	s.router.Get("/main.wasm", file.Content("./main.wasm"))
	s.router.Get("/src/*", file.Server(file.PathMap{"/src/": "/src/"}))

	http.Handle("/", s.router)
}

func (s *server) Run() error {
	return http.ListenAndServe(s.Options.Port, nil)
}
