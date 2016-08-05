package ctrie_mem

import (
	"time"

	"github.com/biinilya/memsrvd/mem"
)

type keyRef struct {
	ref      interface{}
	expireAt time.Time
}

func newKeyRef(value interface{}, ttl time.Duration) *keyRef {
	var ref = &keyRef{ref: value}
	ref.expire(ttl)
	return ref
}

func (ref *keyRef) expire(ttl time.Duration) {
	if ttl > 0 {
		ref.expireAt = time.Now().Add(ttl)
	} else {
		ref.expireAt = time.Time{}
	}
}

func (ref *keyRef) up2date() bool {
	return ref.expireAt.IsZero() || ref.expireAt.After(time.Now())
}

func toHash(h interface{}, found bool) (mem.HashMap, bool, error) {
	if found {
		var ref = h.(*keyRef)
		if ref.up2date() {
			var hash, hOk = ref.ref.(*hashMap)
			if !hOk {
				return nil, false, mem.ErrWrongType
			}
			return hash, true, nil
		}
	}
	return nil, false, nil
}

func toString(h interface{}, found bool) (string, bool, error) {
	if found {
		var ref = h.(*keyRef)
		if ref.up2date() {
			var hash, hOk = ref.ref.(string)
			if !hOk {
				return "", false, mem.ErrWrongType
			}
			return hash, true, nil
		}
	}
	return "", false, nil
}
