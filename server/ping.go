package server

import "github.com/bsm/redeo"

func (srv *memsrv) Ping(out *redeo.Responder, in *redeo.Request) error {
	switch len(in.Args) {
	case 0:
		out.WriteInlineString("PONG")
	case 1:
		out.WriteString(in.Args[0])
	default:
		out.WriteErrorString("ERR wrong number of arguments for 'ping' command")
	}
	return nil
}
