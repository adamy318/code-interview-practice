package hashmap

/***************************

This is a basic implementation of a hashmap that only takes in 
a key-value pair of strings

possibility: extend to include any data type

****************************/


import (
//	"fmt"
)

const SIZE = 10000000

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

// hashing function to calculate where to store data
func (s Key) hashCode() int {
	hash := 13
	return hash * int(s[0]) * int(s[len(s)-1]) * 17 + len(s)
}

// puts a value in the hashmap
func (ht *HashTable) Put(key, value string) {
	keyMap := Key(key)
	hash := keyMap.hashCode()

	// handles hash collision, so searches empty buckets ahead
	for ht.HashTable[hash].key != "" {
		hash++
	}
	ht.HashTable[hash].key = key
	ht.HashTable[hash].value = value
}

// gets the value of a key in a hash map
func (ht *HashTable) Get(key string) string {
	keyMap := Key(key)
	hash := keyMap.hashCode()

	// handles hash collision, so searches empty buckets ahead
	for ht.HashTable[hash].key != key && ht.HashTable[hash].key != "" {
		hash++
	}
	if ht.HashTable[hash].key == "" {
		return "Key does not exist"
	}
	return ht.HashTable[hash].value
}

// deletes a key from the hash map given a key
func (ht *HashTable) Delete(key string) string {
	keyMap := Key(key)
	hash := keyMap.hashCode()

	// handles hash collision, so searches empty buckets ahead
	for ht.HashTable[hash].key != key && ht.HashTable[hash].key != "" {
		hash++
	}
	if ht.HashTable[hash].key == "" {
		return "Key does not exist"
	}
	value := ht.HashTable[hash].value
	ht.HashTable[hash].key = ""
	ht.HashTable[hash].value = ""
	return value
}
