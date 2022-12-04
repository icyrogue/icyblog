//go:build js && wasm
// +build js,wasm

package main

import (
	"log"

	"icy.com/internal/eventhandler"
	"icy.com/internal/loader"
	"icy.com/internal/models"
	"icy.com/internal/renderer"
)

func main() {
	log.Println("hello there")
	resizeEevent := make(chan struct{}, 5)
	drawLineEvent := make(chan models.Line, 16)

	handler := eventhandler.New(resizeEevent, drawLineEvent)
	handler.Start()

	loaderMan := loader.New(handler)
	loaderMan.Scan()

	rendererMan := renderer.New()
	rendererMan.Init()
	rendererMan.Start(drawLineEvent, resizeEevent)

}
