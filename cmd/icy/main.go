package main

import (
	"log"

	"icy.com/internal/devutils"
	"icy.com/internal/server"
)

func main() {
	serverMan := server.New()
	serverMan.Options = &server.Options{Port: ":8080"}
	serverMan.Init()
	watcher := devutils.NewWatcher()
	watcher.Start()
	log.Fatal(serverMan.Run())

}
