package hashmap

import (
	"fmt"
)

const SIZE = 1000

//type HashMap interface {
//	InitHashTable() *HashTable
//}


type HashTable struct {
	hashTable interface{}
}

func main() {
	nw := InitHashTable()
	fmt.Println(nw)
}

func InitHashTable() *HashTable  {
	var array [SIZE]string
	hash := &HashTable{array} 
	return hash
}
