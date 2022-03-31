package main

import (
	"log"

	"github.com/notedit/relay-server/relay"
	"github.com/sevlyar/go-daemon"
)

var configFile = "../../config.yaml"

func main() {

	cntxt := &daemon.Context{
		PidFileName: "relayserver.pid",
		PidFilePerm: 0644,
		LogFileName: "relayserver.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	server, err := relay.NewRelayServer(configFile)
	if err != nil {
		panic(err)
	}

	server.Run()
}
