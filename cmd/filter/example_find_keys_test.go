package main

import (
	"testing"
)

// ExampleFindKeys fills a channel with all of the keys
// in a bucket and drains it safely.
func ExampleFindKeys(t *testing.T) {
	connectStorage()
	keys, nKeys := findKeysInBucket("foldit")
	for i := 0; i < nKeys; i++ {
		_ = <-keys
	}
}
