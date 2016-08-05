package mem

import (
	"errors"
	"time"
)

var (
	ErrWrongType = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
)

type HashMap interface {
	Get(key []byte) ([]byte, bool)
	SetEx(key []byte, value []byte, ttl time.Duration)
	Delete(key []byte) bool
	Expire(key []byte, ttl time.Duration)

	Iter() HashIterator
}

type HashIterator interface {
	Next() (key []byte, value []byte, eof bool)
	Close()
}

type List interface {
	Get(idx int) ([]byte, bool)
	Insert(idx int, value []byte) bool
	Delete(idx int) ([]byte, bool)

	Iter() ListIterator
}

type ListIterator interface {
	Next() (idx int, value []byte, eof bool)
	Close()
}

type Set interface {
	Exists(key []byte) bool
	Insert(key []byte) bool

	Iter() SetIterator
}

type SetIterator interface {
	Next() (value []byte, eof bool)
	Close()
}

type MemCtrl interface {
	Get(key []byte) ([]byte, bool)
	SetEx(key []byte, value []byte, ttl time.Duration)
	Delete(key []byte) bool
	Expire(key []byte, ttl time.Duration)

	Hash(key []byte) (HashMap, error)
}
