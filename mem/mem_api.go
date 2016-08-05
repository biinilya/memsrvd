package mem

import (
	"errors"
	"time"
)

var (
	ErrWrongType = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
)

type HashMap interface {
	Get(key string) (string, bool)
	Set(key string, value string)
	Delete(key string) bool

	Iter() HashIterator
}

type HashIterator interface {
	Next() (key string, value string, eof bool)
	Close()
}

type List interface {
	Get(idx int) (string, bool)
	Insert(idx int, value string) bool
	Delete(idx int) (string, bool)

	Iter() ListIterator
}

type ListIterator interface {
	Next() (idx int, value string, eof bool)
	Close()
}

type Set interface {
	Exists(key string) bool
	Insert(key string) bool

	Iter() SetIterator
}

type SetIterator interface {
	Next() (value string, eof bool)
	Close()
}

type MemCtrl interface {
	Get(key string) (string, bool, error)
	SetEx(key string, value string, ttl time.Duration)
	Delete(key string) bool
	Expire(key string, ttl time.Duration) bool

	Hash(key string) (HashMap, error)
	Close()
}

//go:generate mockgen -source mem_api.go -destination mem_api_mock.go -package mem
