package hashmap

import (
//	"fmt"
)

const SIZE = 10000000

//type HashMap interface {
//	InitHashTable() *HashTable
//}

type Key string

type HashTable struct {
	HashTable [SIZE]KeyValue
}

type KeyValue struct {
	key string
	value string 
}

func InitHashTable() *HashTable  {
	var hash HashTable
	return &hash
}

func (s Key) hashCode() int {
	hash := 13
	return hash * int(s[0]) * int(s[len(s)-1]) * 17 + len(s)
}

func (ht *HashTable) Put(key, value string) {
	keyMap := Key(key)
	hash := keyMap.hashCode()
	for ht.HashTable[hash].key != "" {
		hash++
	}
	ht.HashTable[hash].key = key
	ht.HashTable[hash].value = value
}

func (ht *HashTable) Get(key string) string {
	keyMap := Key(key)
	hash := keyMap.hashCode()
	for ht.HashTable[hash].key != key {
		hash++
	}
	return ht.HashTable[hash].value
}

func (ht *HashTable) Delete(key string) string {
	keyMap := Key(key)
	hash := keyMap.hashCode()
	for ht.HashTable[hash].key != key {
		hash++
	}
	value := ht.HashTable[hash].value
	ht.HashTable[hash].key = ""
	ht.HashTable[hash].value = ""
	return value
}
