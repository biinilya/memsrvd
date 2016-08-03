package server

import "github.com/bsm/redeo"

type MemSrvConfig struct{
	bindAddr string
}

func (cfg *MemSrvConfig) Bind(bindAddr string) *MemSrvConfig{
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
	return srv
}