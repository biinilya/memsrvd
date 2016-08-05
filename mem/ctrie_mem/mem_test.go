package ctrie_mem

import (
	"testing"

	"time"

	"github.com/Workiva/go-datastructures/trie/ctrie"
)

func TestMemFlow(t *testing.T) {
	var m = &ctrieMem{c: ctrie.New(nil)}
	m.c = ctrie.New(nil)

	var v, found, _ = m.Get("xxx")
	if found {
		t.Error("Get::xxx", v)
	}

	m.SetEx("xxx", "yyy", 0)

	v, found, _ = m.Get("xxx")
	if !found || v != "yyy" {
		t.Error("Get::xxx", v)
	}

	found = m.Delete("xxx")
	if !found {
		t.Error("Delete::xxx")
	}

	v, found, _ = m.Get("xxx")
	if found {
		t.Error("Get::xxx", v)
	}

	m.SetEx("xxx", "yyy", time.Millisecond*10)

	v, found, _ = m.Get("xxx")
	if !found || v != "yyy" {
		t.Error("Get::xxx", v)
	}

	time.Sleep(time.Millisecond*10)

	v, found, _ = m.Get("xxx")
	if found {
		t.Error("Get::xxx", v)
	}

	var llen = m.rawLen()
	if llen != 1 {
		t.Error("rawLen", llen)
	}

	m.gc()

	llen = m.rawLen()
	if llen != 0 {
		t.Error("rawLen", llen)
	}
}
