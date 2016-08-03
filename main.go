package main

import (
	"github.com/biinilya/memsrvd/server"
	"log"
)

func main() {
	var cfg server.MemSrvConfig
	var srv = cfg.Bind("0.0.0.0:16379").NewMemSrv()
	log.Panic(srv.Run())
}
