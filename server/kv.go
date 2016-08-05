package server

import (
	"strconv"
	"time"

	"github.com/bsm/redeo"
)

func (srv *memsrv) Set(out *redeo.Responder, in *redeo.Request) error {
	var key string
	var value string
	var expire time.Duration
	switch {
	case len(in.Args)%2 == 1:
		out.WriteErrorString("ERR wrong number of arguments for 'get' command")
		return nil
	case len(in.Args) > 6:
		out.WriteErrorString("ERR wrong number of arguments for 'get' command")
		return nil
	case len(in.Args) == 0:
		out.WriteErrorString("ERR wrong number of arguments for 'get' command")
		return nil
	default:
		key, value = (in.Args[0]), (in.Args[1])
		var args = in.Args[2:]

		for len(args) > 0 {
			switch args[0] {
			case "EX":
				var ttlSec, ttlErr = strconv.ParseUint(args[1], 10, 64)
				if ttlErr != nil {
					out.WriteErrorString("ERR wrong format of seconds for 'get' command: " + ttlErr.Error())
					return nil
				}
				expire += time.Duration(ttlSec) * time.Second
			case "PX":
				var ttlSec, ttlErr = strconv.ParseUint(args[1], 10, 64)
				if ttlErr != nil {
					out.WriteErrorString("ERR wrong format of milliseconds for 'get' command: " + ttlErr.Error())
					return nil
				}
				expire += time.Duration(ttlSec) * time.Millisecond
			default:
				out.WriteErrorString("ERR wrong format of arguments for 'get' command")
				return nil
			}
			args = args[2:]
		}
		srv.ctrl.SetEx(key, value, expire)
	}
	return nil
}

func (srv *memsrv) Get(out *redeo.Responder, in *redeo.Request) error {
	switch len(in.Args) {
	case 1:
		var r, rFound, rErr = srv.ctrl.Get((in.Args[0]))
		if rErr != nil {
			out.WriteErrorString(rErr.Error())
			return nil
		}
		if !rFound {
			out.WriteNil()
		} else {
			out.WriteString(r)
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
		if srv.ctrl.Delete((key)) {
			delCount++
		}
	}
	out.WriteInt(delCount)
	return nil
}

func (srv *memsrv) Expire(out *redeo.Responder, in *redeo.Request) error {
	switch len(in.Args) {
	case 2:
		var ttlSec, ttlErr = strconv.ParseUint(in.Args[1], 10, 64)
		if ttlErr != nil {
			out.WriteErrorString("ERR wrong format of seconds for 'get' command: " + ttlErr.Error())
			return nil
		}

		var ok = srv.ctrl.Expire(in.Args[0], time.Duration(ttlSec)*time.Second)
		if ok {
			out.WriteOne()
		} else {
			out.WriteZero()
		}
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'get' command")
	}
	return nil
}