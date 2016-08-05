package main

import (
	"flag"
	"log"

	"github.com/biinilya/memsrvd/server"
)

var port = flag.String("port", "16379", "port to listen command to")

func main() {
	flag.Parse()

	var cfg server.MemSrvConfig
	var srv = cfg.Bind("0.0.0.0:" + *port).NewMemSrv()
	log.Panic(srv.Run())
}
