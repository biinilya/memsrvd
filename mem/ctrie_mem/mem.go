package ctrie_mem

import (
	"time"

	"sync"

	"log"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/biinilya/memsrvd/mem"
)

type ctrieMem struct {
	c        *ctrie.Ctrie
	hashLock sync.RWMutex
	running  chan struct{}
}

func Mem() mem.MemCtrl {
	var mem = &ctrieMem{
		c:       ctrie.New(nil),
		running: make(chan struct{}),
	}
	go func() {
		// cleanup loop
		for {
			select {
			case <-mem.running:
				return
			default:
				var start = time.Now()
				mem.gc()
				var done = time.Now().Sub(start)
				var sleep = done * 10
				if sleep < time.Minute {
					sleep = time.Minute
				}
				time.Sleep(sleep)
			}
		}
	}()
	return mem
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

func (ctrl *ctrieMem) Expire(key string, ttl time.Duration) bool {
	var ref, found = ctrl.c.Lookup([]byte(key))
	if !found {
		return false
	}
	var kr = ref.(*keyRef)
	if kr.up2date() {
		ref.(*keyRef).expire(ttl)
		return true
	}
	return false
}

func (ctrl *ctrieMem) SetEx(key string, value string, ttl time.Duration) {
	ctrl.c.Insert([]byte(key), newKeyRef(value, ttl))
}

func (ctrl *ctrieMem) Delete(key string) bool {
	var _, found = ctrl.c.Remove([]byte(key))
	return found
}

func (ctrl *ctrieMem) gc() {
	var snapshot = ctrl.c.ReadOnlySnapshot()
	var iter = snapshot.Iterator(nil)
	var cleaned = 0
	for item := range iter {
		var key = string(item.Key)
		var _, found, _ = ctrl.Get(key)
		if !found {
			ctrl.Delete(key)
			cleaned++
		}
	}
	log.Println("expired", cleaned)
}

func (ctrl *ctrieMem) rawLen() int {
	var snapshot = ctrl.c.ReadOnlySnapshot()
	var iter = snapshot.Iterator(nil)
	var llen = 0
	for range iter {
		llen++
	}
	return llen
}

func (ctrl *ctrieMem) Close() {
	defer func() {
		recover()
	}()
	close(ctrl.running)
}
