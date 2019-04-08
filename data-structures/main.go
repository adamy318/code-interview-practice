package main

import (
	"fmt"
	"github.com/adamy318/code-interview-practice/data-structures/hashmap"
)

func main() {
	nw := hashmap.InitHashTable()

	nw.Put("test", "code")
	nw.Put("go", "yes")
	nw.Put("fox", "bear")
	nw.Put("fish", "crab")
	nw.Put("boom", "no")
	nw.Put("brrm", "dkfjdf")

	fmt.Println("Expected -> test: ", nw.Get("test"))
	fmt.Println("Expected -> go: ", nw.Get("go"))
	fmt.Println("Expected -> fox: ", nw.Get("fox"))
	fmt.Println("Expected -> fish: ", nw.Get("fish"))
	fmt.Println("Expected -> boom: ", nw.Get("boom"))
	fmt.Println("Expected -> brrm: ", nw.Get("brrm"))
	fmt.Println("Expected -> key does not exist: ", nw.Get("dkfjdfj"))

	fmt.Println("Deleted: ", nw.Delete("brrm"))
	fmt.Println("Deleted: ", nw.Delete("brjfm"))
}
