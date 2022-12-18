package server

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/access"
	"github.com/go-ozzo/ozzo-routing/v2/fault"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	"github.com/go-ozzo/ozzo-routing/v2/slash"
)

type server struct {
	router   *routing.Router
	logger   logger
	fs       fs.FS
	reqChan  chan string
	template *template.Template
	Options  *Options
}

type logger interface {
	Infof(template string, args ...interface{})
	Errorln(args ...interface{})
}

type Options struct {
	Port         string
	TemplatePath string
}

func New(logger logger, fs fs.FS) *server {
	return &server{logger: logger, fs: fs}
}

func (s *server) Init() {
	s.router = routing.New()
	s.router.Use(
		access.Logger(s.logger.Infof),
		slash.Remover(http.StatusMovedPermanently),
		fault.Recovery(s.logger.Infof),
	)
	s.router.Get("/", file.Content("./src/index.html"))
	s.router.Get("/article/name=<name>", s.handleArticle)
	s.router.Get("/src/*", file.Server(file.PathMap{"/src/": "/src/"}))

	http.Handle("/", s.router)
	template, err := template.ParseGlob(s.Options.TemplatePath)
	if err != nil {
		s.logger.Errorln(err.Error())
		return
	}
	s.template = template
}

func (s *server) Run() error {
	return http.ListenAndServe(s.Options.Port, nil)
}

func (s *server) handleArticle(c *routing.Context) error {
	var err error
	name := c.Param("name")
	if s.template.Lookup(name) == nil {
		s.template, err = s.template.ParseFS(s.fs, name)
	}
	if err != nil {
		s.logger.Errorln(err.Error())
		c.WriteWithStatus(err.Error, http.StatusNotFound)
		return nil
	}
	err = s.template.ExecuteTemplate(c.Response, "header.html", nil)
	if err != nil {
		s.logger.Errorln(err.Error())
		c.WriteWithStatus(err.Error, http.StatusInternalServerError)
		return nil
	}
	err = s.template.ExecuteTemplate(c.Response, name, name)
	if err != nil {
		s.logger.Errorln(err.Error())
		c.WriteWithStatus(err.Error, http.StatusInternalServerError)
		return nil
	}
	return s.template.ExecuteTemplate(c.Response, "footer.html", name)
}
