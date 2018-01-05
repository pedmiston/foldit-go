package main

import (
	"log"
	"os"

	minio "github.com/minio/minio-go"
)

var storage *minio.Client

func main() {
	var keys chan string
	connectStorage()
	keys, nKeys := findKeysInBucket("foldit")
	log.Printf("Found %s keys:", nKeys)

	i := 1
	for k := range keys {
		log.Printf("%s: %s", i, k)
		i++
	}

}

// connectStorage connects the global var storage to AWS S3.
func connectStorage() {
	var err error

	endpoint := "s3.amazonaws.com"
	accessKey := lookupEnv("AWS_ACCESS_KEY")
	secretKey := lookupEnv("AWS_SECRET_KEY")
	useSSL := true

	storage, err = minio.New(endpoint, accessKey, secretKey, useSSL)

	if err != nil {
		log.Fatal(err)
	}
}

// lookupEnv tries to find a key in the environment
func lookupEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found || value == "" {
		log.Fatalf("Environment variable '%s' not defined", key)
	}
	return value
}

// findKeysInBucket fills a chan with keys.
func findKeysInBucket(bucket string) (keys chan string, nKeys int) {
	keys = make(chan string)

	prefix := "" // list all objects in bucket
	isRecursive := false

	doneCh := make(chan struct{})
	defer close(doneCh)

	objectInfoCh := storage.ListObjects(bucket, prefix, isRecursive, doneCh)

	for objectInfo := range objectInfoCh {
		if objectInfo.Err != nil {
			return keys, nKeys
		}

		go func(k string) { keys <- k }(objectInfo.Key)
		nKeys++
	}

	return keys, nKeys
}
