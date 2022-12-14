package processor

import (
	"os"
	"path"
	"strings"
	"sync"

	"github.com/Depado/bfchroma/v2"
	bf "github.com/russross/blackfriday/v2"
	"icy.com/internal/models"
)

type processor struct {
	logger   Logger
	cacheMap map[string]string
	renderer *bfchroma.Renderer
	mtx      *sync.RWMutex
	Options  *Options
}

type Options struct {
	CacheDir string
}

type Logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}

func New(logger Logger) *processor {
	return &processor{logger: logger, mtx: &sync.RWMutex{}}
}

func (p *processor) Init() {
	renderer := bfchroma.NewRenderer(bfchroma.Style("dracula"))
	p.renderer = renderer
}

func (p *processor) Start(input chan models.Article) {
	go func() {
	loop:
		for {
			select {
			case a := <-input:
				p.logger.Infoln("got a new input", a.Name)
				err := p.parse(a)
				if err != nil {
					p.logger.Errorln(err.Error())
					continue loop
				}
			}
		}
	}()
}

func (p *processor) parse(article models.Article) error {
	html := bf.Run(article.Content, bf.WithRenderer(p.renderer))
	return p.cache(article.Name, html)
}

func (p *processor) cache(name string, data []byte) error {
	p.mtx.Lock()
	p.mtx.Unlock()
	name = strings.Split(name, ".")[0]
	err := os.WriteFile(path.Join(p.Options.CacheDir, name), data, 0666)
	if err != nil {
		return err
	}
	return nil
}
