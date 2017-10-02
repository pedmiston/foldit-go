package main

import (
	"fmt"
	"log"
	"os"

	minio "github.com/minio/minio-go"
)

var storage *minio.Client

// connectStorage connects to Digital Ocean Spaces via the global var storage
func connectStorage() {
	var err error
	endpoint := "nyc3.digitaloceanspaces.com"
	useSSL := true
	storage, err = minio.New(endpoint, auth.DOAccessKey, auth.DOSecretKey, useSSL)
	if err != nil {
		log.Fatal(err)
	}
}

// loadKeysFromBucket stores all object keys in a bucket into a channel.
func loadKeysFromBucket(bucket string) (chan string, int) {
	var n int
	keys := make(chan string)

	// Create a done channel to control 'ListObjects' go routine.
	done := make(chan struct{})
	defer close(done)

	var prefix string
	isRecursive := true
	objectCh := storage.ListObjects(bucket, prefix, isRecursive, done)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Fprintf(os.Stderr, "error in listing objects: %v\n", object.Err)
			break
		}
		go func(k string) { keys <- k }(object.Key)
		n++
	}
	return keys, n
}

func loadKey(key, bucket string) chan string {
	ok := checkKey(key, bucket)
	if !ok {
		fmt.Fprintf(os.Stderr, "key '%v' not found in bucket '%v'\n", key, bucket)
		os.Exit(1)
	}
	keys := make(chan string)
	go func(k string) { keys <- k }(key)
	return keys
}

func checkKey(key, bucket string) bool {
	_, err := storage.StatObject(bucket, key, minio.StatObjectOptions{})
	return err == nil
}

func getObject(key, bucket string) (*minio.Object, error) {
	object, err := storage.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}
