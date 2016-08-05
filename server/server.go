package server

import (
	"github.com/biinilya/memsrvd/mem"
	"github.com/bsm/redeo"
)

type MemSrv interface {
	Run() error
}

type memsrv struct {
	redeoSrv *redeo.Server
	ctrl     mem.MemCtrl
}

func (srv *memsrv) Run() error {
	return srv.redeoSrv.ListenAndServe()
}
