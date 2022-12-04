package devutils

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type watcher struct {
}

func NewWatcher() *watcher {
	return &watcher{}
}

func (w *watcher) Start() {
	go func() {
		var prevSize int64
		for {
			time.Sleep(5 * time.Second)
			var size int64
			err := filepath.Walk("internal", func(_ string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					size += info.Size()
				}
				return err
			})
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if size != prevSize {
				log.Println("started build")
				if err = w.build(); err != nil {
					log.Println(err.Error())
				}
				prevSize = size
			}
		}
	}()
}

func (w *watcher) build() error {
	cmd := exec.Command("go", "build", "-o", "../../main.wasm")
	cmd.Dir = "cmd/renderer"

	cmd.Env = append(cmd.Environ(), "GOOS=js", "GOARCH=wasm")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
