package kvstore

import (
	"sync"
	"time"
)

type Data struct {
	value      string
	expiration time.Time
}

type KeyValueStore struct {
	mu   sync.RWMutex
	data map[string]Data
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]Data),
	}
}

func (kv *KeyValueStore) Set(key, val string, ttl time.Duration) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[key] = Data{
		value:      val,
		expiration: time.Now().Add(ttl),
	}
}

func (kv *KeyValueStore) Get(key string) (string, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	val, ok := kv.data[key]

	// check for item expiry
	if val.expiration.IsZero() || time.Now().Before(val.expiration) {
		return val.value, ok
	}

	// delete if expired
	delete(kv.data, key)
	return "", false
}

func (kv *KeyValueStore) Delete(key string) bool {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	_, ok := kv.data[key]
	if !ok {
		return false
	}

	delete(kv.data, key)
	return true
}
