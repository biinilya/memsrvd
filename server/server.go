package server

import "github.com/bsm/redeo"

type MemSrv interface {
	Run() error
}

type memsrv struct {
	redeoSrv *redeo.Server
}

func (srv *memsrv) Run() error {
	return srv.redeoSrv.ListenAndServe()
}
