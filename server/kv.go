package server

import (
	"strconv"
	"time"

	"github.com/bsm/redeo"
)

func (srv *memsrv) Set(out *redeo.Responder, in *redeo.Request) error {
	var key []byte
	var value []byte
	var expire time.Duration
	switch {
	case len(in.Args) >= 2:
		key, value = []byte(in.Args[0]), []byte(in.Args[1])
	case len(in.Args) >= 4 && in.Args[2] == "EX":
		var ttlSec, ttlErr = strconv.ParseUint(in.Args[3], 10, 64)
		if ttlErr != nil {
			out.WriteErrorString("ERR wrong format of seconds for 'get' command: " + ttlErr.Error())
			return nil
		}
		expire += time.Duration(ttlSec) * time.Second
	case len(in.Args) >= 4 && in.Args[2] == "PX":
		var ttlSec, ttlErr = strconv.ParseUint(in.Args[3], 10, 64)
		if ttlErr != nil {
			out.WriteErrorString("ERR wrong format of milliseconds for 'get' command: " + ttlErr.Error())
			return nil
		}
		expire += time.Duration(ttlSec) * time.Millisecond
	case len(in.Args) == 6 && in.Args[4] == "PX":
		var ttlSec, ttlErr = strconv.ParseUint(in.Args[5], 10, 64)
		if ttlErr != nil {
			out.WriteErrorString("ERR wrong format of seconds for 'get' command: " + ttlErr.Error())
			return nil
		}
		expire += time.Duration(ttlSec) * time.Second
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'get' command")
	}
	srv.ctrl.SetEx(key, value, expire)
	return nil
}

func (srv *memsrv) Get(out *redeo.Responder, in *redeo.Request) error {
	switch len(in.Args) {
	case 1:
		var r, _ = srv.ctrl.Get([]byte(in.Args[0]))
		if r == nil {
			out.WriteNil()
		} else {
			out.WriteBytes(r)
		}
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'get' command")
	}
	return nil
}

func (srv *memsrv) Del(out *redeo.Responder, in *redeo.Request) error {
	switch len(in.Args) {
	case 0:
		out.WriteErrorString("ERR wrong number of arguments for 'del' command")
		return nil
	}

	var delCount = 0
	for _, key := range in.Args {
		if srv.ctrl.Delete([]byte(key)) {
			delCount++
		}
	}
	out.WriteInt(delCount)
	return nil
}
