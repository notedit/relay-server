package main

import (
	"log"

	"github.com/notedit/relay-server/relay"
)

var configFile = "../../config.yaml"

func main() {

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	server, err := relay.NewRelayServer(configFile)
	if err != nil {
		panic(err)
	}

	server.Run()
}
