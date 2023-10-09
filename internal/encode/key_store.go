package encode

import (
	"sort"
)

type KeyStore interface {
	Keys() []string
	IndexOf(key string) int
	Add(key string) int
	Reset()
}

type keyStore struct {
	keys   []string
	sorted bool
}

func NewKeyStoreWithKeys(keys []string) *keyStore {
	if keys == nil {
		return NewKeyStore()
	}
	return &keyStore{keys: keys}
}

func NewKeyStore() *keyStore {
	return &keyStore{keys: []string{}}
}

func (k *keyStore) Keys() []string {
	k.sort()
	return k.keys
}

func (k *keyStore) IndexOf(key string) int {
	k.sort()
	return sort.SearchStrings(k.keys, key)
}

func (k *keyStore) Add(key string) int {
	idx := k.IndexOf(key)
	if len(k.keys) <= idx || k.keys[idx] != key {
		k.sorted = false
		k.keys = append(k.keys, key)
		return len(k.keys) - 1
	}
	return idx
}

func (k *keyStore) Reset() {
	k.keys = []string{}
	k.sorted = true
}

func (k *keyStore) sort() {
	if !k.sorted {
		sort.Strings(k.keys)
		k.sorted = true
	}
}
