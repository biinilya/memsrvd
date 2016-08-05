package server

import (
	"github.com/biinilya/memsrvd/mem/ctrie_mem"
	"github.com/bsm/redeo"
)

type MemSrvConfig struct {
	bindAddr string
}

func (cfg *MemSrvConfig) Bind(bindAddr string) *MemSrvConfig {
	cfg.bindAddr = bindAddr
	return cfg
}

func (cfg *MemSrvConfig) NewMemSrv() MemSrv {
	var srv = &memsrv{}
	if cfg.bindAddr == "" {
		cfg.bindAddr = "127.0.0.1:6379"
	}
	srv.redeoSrv = redeo.NewServer(&redeo.Config{
		Addr: cfg.bindAddr,
	})
	srv.redeoSrv.HandleFunc("ping", srv.Ping)

	srv.redeoSrv.HandleFunc("set", srv.Set)
	srv.redeoSrv.HandleFunc("get", srv.Get)
	srv.redeoSrv.HandleFunc("del", srv.Del)

	srv.redeoSrv.HandleFunc("hset", srv.HSet)
	srv.redeoSrv.HandleFunc("hget", srv.HGet)
	srv.redeoSrv.HandleFunc("hdel", srv.HDel)
	srv.redeoSrv.HandleFunc("hkeys", srv.HKeys)
	srv.ctrl = ctrie_mem.Mem()
	return srv
}
