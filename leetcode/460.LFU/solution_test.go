package _60_LFU

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	cache := Constructor(3)
	fmt.Println(cache)
	cache.Put(1, 11)
	fmt.Println(cache)
	cache.Put(2, 22)
	cache.Put(9, 99)
	cache.Put(1, 111)
	fmt.Println(cache)
	cache.Put(3, 33)
	fmt.Println(cache)
}
