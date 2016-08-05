package ctrie_mem

import (
	"time"

	"sync"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/biinilya/memsrvd/mem"
)

var nsGlobals = string("globals")

type ctrieMem struct {
	c        *ctrie.Ctrie
	hashLock sync.RWMutex
}

func Mem() mem.MemCtrl {
	return &ctrieMem{c: ctrie.New(nil)}
}

func (ctrl *ctrieMem) Hash(key string) (mem.HashMap, error) {
	// Fastpath
	ctrl.hashLock.RLock()
	var h, hFound, _ = toHash(ctrl.c.Lookup([]byte(key)))
	ctrl.hashLock.RUnlock()
	if hFound {
		return h, nil
	}

	// Slowpath
	ctrl.hashLock.Lock()
	h, hFound, _ = toHash(ctrl.c.Lookup([]byte(key)))
	if !hFound {
		h = HashMap()
		ctrl.c.Insert([]byte(key), newKeyRef(h, 0))
	}
	ctrl.hashLock.Unlock()
	return h, nil
}

func (ctrl *ctrieMem) Get(key string) (string, bool, error) {
	return toString(ctrl.c.Lookup([]byte(key)))
}

func (ctrl *ctrieMem) Expire(key string, ttl time.Duration) {
	var ref, found = ctrl.c.Lookup([]byte(key))
	if !found {
		return
	}
	ref.(*keyRef).expire(ttl)
}

func (ctrl *ctrieMem) SetEx(key string, value string, ttl time.Duration) {
	ctrl.c.Insert([]byte(key), newKeyRef(value, ttl))
}

func (ctrl *ctrieMem) Delete(key string) bool {
	var _, found = ctrl.c.Remove([]byte(key))
	return found
}

func (ctrl *ctrieMem) Close() {

}