package ctrie_mem

import (
	"time"

	"github.com/Workiva/go-datastructures/trie/ctrie"
	"github.com/biinilya/memsrvd/mem"
)

type hashMapIterator struct {
	entries <-chan *ctrie.Entry
	closer  chan struct{}
}

func (iter hashMapIterator) Next() (key []byte, value []byte, eof bool) {
	var entry, opened = <-iter.entries
	if !opened {
		eof = true
		return
	}
	key = entry.Key
	value = entry.Value.([]byte)
	return
}
func (iter hashMapIterator) Close() {
	defer func() {
		recover()
	}()
	close(iter.closer)
}

type hashMap struct {
	c *ctrie.Ctrie
}

// HashMap creates an empty hash map
func HashMap() mem.HashMap {
	var m = &hashMap{}
	m.c = ctrie.New(nil)
	return m
}

func (hash *hashMap) Iter() mem.HashIterator {
	var sn = hash.c.ReadOnlySnapshot()
	var snCloser = make(chan struct{})
	return &hashMapIterator{
		entries: sn.Iterator(snCloser),
		closer:  snCloser,
	}
}

// Remove deletes the value for the associated key, returning true if it was
// removed or false if the entry doesn't exist.
func (hash *hashMap) Delete(key []byte) (found bool) {
	_, found = hash.c.Remove(key)
	return
}

// Get returns the value for the associated key or returns false if the key
// doesn't exist.
func (hash *hashMap) Get(key []byte) (value []byte, found bool) {
	var entry interface{}
	entry, found = hash.c.Lookup(key)
	if !found {
		return
	}
	value = entry.([]byte)
	return
}

// Set adds the key-value pair to the hash map, replacing the existing value if
// the key already exists.
func (hash *hashMap) SetEx(key []byte, value []byte, ttl time.Duration) {
	hash.c.Insert(key, value)
	return
}

func (hash *hashMap) Expire(key []byte, ttl time.Duration) {
	return
}
