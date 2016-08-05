package server

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/biinilya/memsrvd/mem"
	"github.com/bsm/redeo"
	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var memctrl = mem.NewMockMemCtrl(ctrl)
	var srv = &memsrv{}
	srv.ctrl = memctrl

	var test = func(in string, args ...interface{}) {
		var req = &redeo.Request{Args: strings.Split(in, " ")}
		var outBuf = bytes.NewBuffer(nil)
		var resp = redeo.NewResponder(outBuf)

		switch len(args) {
		case 5:
			var msg = args[4].(string)
			memctrl.EXPECT().Get(args[0]).Return(args[1], args[2], args[3])

			srv.Get(resp, req)
			resp.Flush()
			if msg != outBuf.String() {
				t.Errorf("'%v' != '%v'", msg, outBuf.String())
			}
		case 1:
			memctrl.EXPECT()
			srv.Set(resp, req)
			resp.Flush()
			if args[0].(string) != outBuf.String() {
				t.Errorf("'%v' != '%v'", args[0].(string), outBuf.String())
			}
		default:
			t.Log("Invalid helper usage")
			t.FailNow()
		}
	}

	test("a", "a", "b", true, nil, "$1\r\nb\r\n")
	test("a", "a", "", false, nil, "$-1\r\n")
	test("a", "a", "", false, errors.New("xxx"), "-xxx\r\n")
	test("a b EX", "-ERR wrong number of arguments for 'get' command\r\n")
}

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var memctrl = mem.NewMockMemCtrl(ctrl)
	var srv = &memsrv{}
	srv.ctrl = memctrl

	var test = func(in string, args ...interface{}) {
		var req = &redeo.Request{Args: strings.Split(in, " ")}
		var outBuf = bytes.NewBuffer(nil)
		var resp = redeo.NewResponder(outBuf)

		switch len(args) {
		case 3:
			memctrl.EXPECT().SetEx(args[0], args[1], args[2])
			srv.Set(resp, req)
		case 1:
			memctrl.EXPECT()
			srv.Set(resp, req)
			resp.Flush()
			if args[0].(string) != outBuf.String() {
				t.Errorf("'%v' != '%v'", args[0].(string), outBuf.String())
			}
		default:
			t.Log("Invalid helper usage")
			t.FailNow()
		}
	}

	test("a b", "a", "b", time.Duration(0))
	test("a b EX 1", "a", "b", time.Second)
	test("a b PX 5", "a", "b", time.Millisecond*5)
	test("a b EX 1 PX 5", "a", "b", time.Millisecond*5+time.Second)
	test("a b EX", "-ERR wrong number of arguments for 'get' command\r\n")
}

func TestDel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var memctrl = mem.NewMockMemCtrl(ctrl)
	var srv = &memsrv{}
	srv.ctrl = memctrl

	var test = func(in string, sample string, keys ...string) {
		var req = &redeo.Request{Args: strings.Split(in, " ")}
		var outBuf = bytes.NewBuffer(nil)
		var resp = redeo.NewResponder(outBuf)

		if len(keys) > 0 {
			for _, key := range keys {
				memctrl.EXPECT().Delete(key).Return(true)
			}
		} else {
			memctrl.EXPECT()
		}
		srv.Del(resp, req)
		resp.Flush()
		if sample != outBuf.String() {
			t.Errorf("'%v' != '%v'", sample, outBuf.String())
		}
	}

	test("a", ":1\r\n", "a")
	test("a b", ":2\r\n", "a", "b")
}
