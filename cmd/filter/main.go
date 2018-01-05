package main

import (
	"log"
	"os"

	minio "github.com/minio/minio-go"
)

var storage *minio.Client
var bucket = "foldit"

func main() {
	var keysCh chan string
	var objsCh chan *minio.Object
	connectStorage()
	keysCh, nKeys := findKeysInBucket()
	objsCh, nObjs = getObjectsFromKeys(keysCh, nKeys)

	// Manage goroutines pulling objects from the channel
	//
	// for obj in objsCh:
	//   for each line in obj:
	//     extract filepath
	//     is this filepath unique?
	//     yes -> save line to output
	//     no -> next line
	//
	//   is there any output?
	//   yes -> push output to S3, overwriting source
	//   no  -> delete source on S3
	//
	//   signal done

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
func findKeysInBucket() (keys chan string, nKeys int) {
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

func getObjectsFromKeys(keysCh chan string, nKeys int) (objsCh chan *minio.Object, nObjs int) {
	objsCh = make(chan *minio.Object)

	for i := 0; i < nKeys; i++ {
		key := <-keysCh

		go func(k string) {
			obj, _ := storage.GetObject(bucket, k, minio.GetObjectOptions{})
			objsCh <- obj
		}(key)

		nObjs++
	}

	return objsCh, nObjs
}
