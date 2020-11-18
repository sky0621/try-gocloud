package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/xerrors"

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

	obj, err := getObject(ctx, bucketURL, file)
	if err != nil {
		log.Fatalf("Failed to getObject: %s", err)
	}

	fmt.Println(obj)
}

func getObject(ctx context.Context, bucketURL, file string) (string, error) {
	bucket, err := blob.OpenBucket(ctx, bucketURL)
	if err != nil {
		return "", xerrors.Errorf("Failed to setup bucket: %w", err)
	}
	defer func() {
		if err := bucket.Close(); err != nil {
			log.Printf("Failed to close: %s", err)
		}
	}()

	b, err := bucket.ReadAll(ctx, file)
	if err != nil {
		return "", xerrors.Errorf("Failed to read file: %w", err)
	}

	return string(b), nil
}
