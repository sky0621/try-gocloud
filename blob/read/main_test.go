package main

import (
	"context"
	"fmt"
	"testing"

	_ "gocloud.dev/blob/fileblob"
)

func TestGetObject(t *testing.T) {
	ctx := context.Background()
	obj, err := getObject(ctx, "file://anyway", "abc.txt")
	if err != nil {
		t.Fatalf("failed to getObject: %s", err)
	}
	fmt.Println(obj)
}
