package server

import "github.com/bsm/redeo"

func (srv *memsrv) HSet(out *redeo.Responder, in *redeo.Request) error {
	switch len(in.Args) {
	case 3:
		var hash, hErr = srv.ctrl.Hash((in.Args[0]))
		if hErr != nil {
			out.WriteErrorString(hErr.Error())
			return nil
		}
		var key, value = (in.Args[1]), (in.Args[2])
		var found = hash.Delete(key)
		hash.Set(key, value)
		if found {
			out.WriteInt(1)
		} else {
			out.WriteInt(0)
		}
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'hget' command")
	}
	return nil
}

func (srv *memsrv) HGet(out *redeo.Responder, in *redeo.Request) error {
	switch len(in.Args) {
	case 2:
		var hash, hErr = srv.ctrl.Hash((in.Args[0]))
		if hErr != nil {
			out.WriteErrorString(hErr.Error())
			return nil
		}
		var r, rFound = hash.Get((in.Args[1]))
		if !rFound {
			out.WriteNil()
		} else {
			out.WriteString(r)
		}
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'hget' command")
	}
	return nil
}

func (srv *memsrv) HDel(out *redeo.Responder, in *redeo.Request) error {
	switch {
	case len(in.Args) > 1:
		var hash, hErr = srv.ctrl.Hash((in.Args[0]))
		if hErr != nil {
			out.WriteErrorString(hErr.Error())
			return nil
		}
		var delCount = 0
		for _, key := range in.Args[1:] {
			if hash.Delete((key)) {
				delCount++
			}
		}
		out.WriteInt(delCount)
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'hget' command")
	}
	return nil
}

func (srv *memsrv) HKeys(out *redeo.Responder, in *redeo.Request) error {
	switch {
	case len(in.Args) == 1:
		var hash, hErr = srv.ctrl.Hash((in.Args[0]))
		if hErr != nil {
			out.WriteErrorString(hErr.Error())
			return nil
		}
		var iter = hash.Iter()
		var keys = []string{}
		// FIXME: Not effective, should implement a kind of io.Reader
		for {
			var key, _, last = iter.Next()
			if last {
				break
			}
			keys = append(keys, key)
		}
		out.WriteStringBulk(keys)
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'hget' command")
	}
	return nil
}
