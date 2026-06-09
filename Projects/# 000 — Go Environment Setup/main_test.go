package main

import (
	"runtime"
	"testing"
)

func TestRuntimeValuesAreAvailable(t *testing.T) {
	if runtime.GOOS == "" {
		t.Fatal("GOOS is empty")
	}

	if runtime.GOARCH == "" {
		t.Fatal("GOARCH is empty")
	}
}
