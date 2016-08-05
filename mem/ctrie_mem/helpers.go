package ctrie_mem

import "github.com/biinilya/memsrvd/mem"

func toHash(h interface{}) (*hashMap, error) {
	var hash, hOk = h.(*hashMap)
	if !hOk {
		return nil, mem.ErrWrongType
	}
	return hash, nil
}

func toString(h interface{}) ([]byte, error) {
	var hash, hOk = h.([]byte)
	if !hOk {
		return nil, mem.ErrWrongType
	}
	return hash, nil
}
