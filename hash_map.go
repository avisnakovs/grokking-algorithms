package grokking_algorithms

import (
	"crypto/sha1"
	"encoding/binary"
)

type HashMap interface {
	Set(string, string)
	Get(string) (string, bool)
}

func NewHashMap() HashMap {
	return &hashMap{size: 0, len: 10, slice: make([]*keyValue, 10, 10)}
}

type hashMap struct {
	size  int
	len   uint64
	slice []*keyValue
}

func (h *hashMap) Set(key, value string) {
	h.size++
	h.addKeyValue(key, value)

	if h.mapOverflow() {
		h.increaseMapLen()
	}
}

func (h *hashMap) mapOverflow() bool {
	return float32(h.size)/float32(len(h.slice)) >= 0.7
}

func (h *hashMap) increaseMapLen() {
	h.len = uint64(len(h.slice) * 2)
	oldSlice := h.slice
	h.slice = make([]*keyValue, h.len, h.len)
	for _, kv := range oldSlice {
		for kv != nil {
			h.addKeyValue(kv.key, kv.value)
			kv = kv.next
		}
	}
}

func (h *hashMap) addKeyValue(key string, value string) {
	index := h.getKeyIndex(key)
	if kv := h.slice[index]; kv == nil {
		h.slice[index] = &keyValue{key: key, value: value}
	} else {
		kv.add(key, value)
	}
}

func (h *hashMap) Get(key string) (string, bool) {
	index := h.getKeyIndex(key)
	return h.slice[index].find(key)
}

func (h *hashMap) getKeyIndex(key string) uint64 {
	return generateHash(key) % h.len
}

func generateHash(value string) uint64 {
	hash := sha1.New()
	hash.Write([]byte(value))
	return binary.BigEndian.Uint64(hash.Sum(nil))
}

type keyValue struct {
	key   string
	value string
	next  *keyValue
}

func (kv *keyValue) add(key, value string) {
	next := *kv
	*kv = keyValue{key: key, value: value, next: &next}
}

func (kv *keyValue) find(key string) (string, bool) {
	if kv == nil {
		return "", false
	} else if kv.key == key {
		return kv.value, true
	}
	return kv.next.find(key)
}
