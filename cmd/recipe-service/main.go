package main

import (
	"log"

	"github.com/ajwinebrenner/recipe-service/internal/config"
	"github.com/ajwinebrenner/recipe-service/internal/server"
)

func main() {
	c := mustLoadConfig() //must implies 'or fatal error'
	server.Start(c)
}

// seperate out server config in one file to parse into server
func mustLoadConfig() *config.Config {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}
	return conf
}
