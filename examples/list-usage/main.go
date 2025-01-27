package main

import (
	"fmt"
	"github.com/lehrcode/tutgo-collections"
	"github.com/lehrcode/tutgo-collections/linkedlist"
	"github.com/lehrcode/tutgo-collections/slicelist"
)

func main() {
	var list collections.List[string] = slicelist.New[string]()

	list.Push("eins")
	list.Push("zwei")
	list.Push("drei")
	list.Push("vier")

	fmt.Println(list)

	fmt.Println(list.Get(2))

	list.Delete(1)

	fmt.Println(list)

	var list2 collections.List[string] = linkedlist.New[string]()

	list2.Push("eins")
	list2.Push("zwei")
	list2.Push("drei")
	list2.Push("vier")

	fmt.Println(list2)

	fmt.Println(list2.Get(2))

	list2.Delete(1)

	fmt.Println(list2)
}
