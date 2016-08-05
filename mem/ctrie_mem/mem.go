package ctrie_mem

import (
	"time"

	"sync"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/biinilya/memsrvd/mem"
)

var nsGlobals = []byte("globals")

type ctrieMem struct {
	hm *hashMap
	hashLock sync.RWMutex
}

func Mem() mem.MemCtrl {
	var m = &hashMap{}
	m.c = ctrie.New(nil)
	return &ctrieMem{hm: m}
}

func (ctrl *ctrieMem) hash(key []byte) (*hashMap, error) {
	// Fastpath
	ctrl.hashLock.RLock()
	var h, hFound = ctrl.hm.c.Lookup(key)
	ctrl.hashLock.RUnlock()
	if hFound {
		return toHash(h)
	}

	// Slowpath
	ctrl.hashLock.Lock()
	h, hFound = ctrl.hm.c.Lookup(key)
	if !hFound {
		h = HashMap()
		ctrl.hm.c.Insert(key, h)
	}
	ctrl.hashLock.Unlock()
	return toHash(h)
}

func (ctrl *ctrieMem) Hash(key []byte) (mem.HashMap, error) {
	var h, hErr = ctrl.hash(key)
	return h, hErr
}

func (ctrl *ctrieMem) Get(key []byte) ([]byte, bool) {
	var globals, gErr = ctrl.hash(nsGlobals)
	if gErr != nil {
		return nil, false
	}
	var r, found = globals.Get(key)
	return r, found
}

func (ctrl *ctrieMem) Expire(key []byte, ttl time.Duration) {
}

func (ctrl *ctrieMem) SetEx(key []byte, value []byte, ttl time.Duration) {
	var globals, gErr = ctrl.hash(nsGlobals)
	if gErr != nil {
		return
	}
	globals.SetEx(key, value, ttl)
}

func (ctrl *ctrieMem) Delete(key []byte) bool {
	var globals, gErr = ctrl.hash(nsGlobals)
	if gErr != nil {
		return false
	}
	return globals.Delete(key)
}

func (ctrl *ctrieMem) Iter(key []byte) bool {
	var globals, gErr = ctrl.hash(nsGlobals)
	if gErr != nil {
		return false
	}
	return globals.Delete(key)
}

