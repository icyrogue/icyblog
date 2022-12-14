package storage

import (
	"io/fs"
	"os"
)

type storage struct {
	fs fs.FS
}

func New() *storage {
	return &storage{fs: os.DirFS("./src/cache")}
}

func (s *storage) Open(name string) (fs.File, error) {
	return s.fs.Open(name)
}
