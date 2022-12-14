package main

import (
	"log"

	"icy.com/internal/loader"
	"icy.com/internal/logger"
	"icy.com/internal/models"
	"icy.com/internal/processor"
	"icy.com/internal/server"
	"icy.com/internal/storage"
)

func main() {
	loggerMan := logger.New(&logger.Options{})

	output := make(chan models.Article, 10)

	loaderMan := loader.New(loggerMan)
	loaderMan.Options = &loader.Options{ConfigPath: "./config/config.toml", ContentPath: "./content"}
	loaderMan.Init()
	loaderMan.Scan(output)
	processorMan := processor.New(loggerMan)
	processorMan.Options = &processor.Options{CacheDir: "./src/cache"}
	processorMan.Init()
	processorMan.Start(output)
	storageMan := storage.New()
	serverMan := server.New(loggerMan, storageMan)
	serverMan.Options = &server.Options{Port: ":8080", TemplatePath: "./config/*.html"}
	serverMan.Init()
	log.Fatal(serverMan.Run())
}
