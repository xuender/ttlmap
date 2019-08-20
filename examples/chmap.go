package main

import (
	"fmt"

	"github.com/xuender/toolkit"
)

func main() {
	chMap := toolkit.NewChMap()
	defer chMap.Close()
	chMap.Put("key1", "value1")
	chMap.Put("key2", "value2")

	fmt.Println(chMap.Get("key1"))
	fmt.Println(chMap.Count())
}