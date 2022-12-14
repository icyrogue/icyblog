package loader

import (
	"os"

	"icy.com/internal/models"
)

type loader struct {
	logger  logger
	Options *Options
}

type logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}

type Options struct {
	ConfigPath  string
	ContentPath string
}

func New(logger logger) *loader {
	return &loader{logger: logger}
}

func (l *loader) Init() error {
	l.logger.Infoln("loading configurations")
	return nil
}

func (l *loader) Scan(output chan models.Article) {
	files, err := os.ReadDir(l.Options.ContentPath)
	if err != nil {
		l.logger.Errorln(err.Error())
	}
	for _, file := range files {
		content, err := os.ReadFile(l.Options.ContentPath + "/" + file.Name())
		if err != nil {
			l.logger.Errorln(err.Error())
			continue
		}
		output <- models.Article{
			Content: content,
			Name:    file.Name(),
		}
	}
}
