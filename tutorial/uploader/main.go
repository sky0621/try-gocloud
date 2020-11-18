package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

func main() {
	// Define our input.
	if len(os.Args) != 3 {
		log.Fatal("usage: upload BUCKET_URL FILE")
	}
	bucketURL := os.Args[1]
	file := os.Args[2]

	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, bucketURL)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}
	defer func() {
		if err := b.Close(); err != nil {
			log.Printf("Failed to close: %s", err)
		}
	}()

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	w, err := b.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %s", err)
	}

	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to bucket: %s", err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			log.Printf("Failed to close: %s", err)
		}
	}()
}
