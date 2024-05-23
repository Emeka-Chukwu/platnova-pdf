package main

import (
	"log"
	"pdf-go/cmd"
	"pdf-go/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot set config: %v", err.Error())
	}
	runGinServer(config)
}

func runGinServer(config util.Config) {
	server, err := cmd.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server")
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
