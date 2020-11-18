package main

import (
	"context"
	"fmt"
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

	bucket, err := blob.OpenBucket(ctx, bucketURL)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}
	defer func() {
		if err := bucket.Close(); err != nil {
			log.Printf("Failed to close: %s", err)
		}
	}()

	b, err := bucket.ReadAll(ctx, file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	fmt.Println(string(b))
}
