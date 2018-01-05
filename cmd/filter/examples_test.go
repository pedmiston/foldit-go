package main

import (
	"log"
)

// ExampleFindKeys fills a channel with all of the keys
// in a bucket and drains it safely.
func ExampleFindKeys() {
	connectStorage()
	keys, nKeys := findKeysInBucket()
	for i := 0; i < nKeys; i++ {
		_ = <-keys
	}
	log.Printf("Found %s keys", nKeys)
}

func ExampleGetObjects() {
	connectStorage()
	keysCh, nKeys := findKeysInBucket()
	objsCh, nObjs := getObjectsFromKeys(keysCh, nKeys)
	for i := 0; i < nObjs; i++ {
		_ = <-objsCh
	}
	log.Printf("Found %s objects", nObjs)
}
