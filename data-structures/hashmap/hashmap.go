package hashmap

import (
//	"fmt"
)

const SIZE = 1000

//type HashMap interface {
//	InitHashTable() *HashTable
//}


type HashTable struct {
	HashTable interface{}
}

func InitHashTable() *HashTable  {
	var array [SIZE]string
	hash := &HashTable{array} 
	return hash
}
