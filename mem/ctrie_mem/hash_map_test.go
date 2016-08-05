package ctrie_mem

import (
	"testing"

	"github.com/Workiva/go-datastructures/trie/ctrie"
)

func TestHashMap(t *testing.T) {
	var m = &hashMap{}
	m.c = ctrie.New(nil)

	var v, found = m.Get("xxx")
	if found {
		t.Error("Get::xxx", v)
	}

	m.Set("xxx", "yyy")

	v, found = m.Get("xxx")
	if !found || v != "yyy" {
		t.Error("Get::xxx", v)
	}

	found = m.Delete("xxx")
	if !found {
		t.Error("Delete::xxx")
	}

	v, found = m.Get("xxx")
	if found {
		t.Error("Get::xxx", v)
	}
}
