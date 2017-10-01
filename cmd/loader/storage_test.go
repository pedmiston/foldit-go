package main

import (
	"testing"
)

func TestListKeysInBucket(t *testing.T) {
	authenticate("secrets.yml")
	connectStorage()
}
