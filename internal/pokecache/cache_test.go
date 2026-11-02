package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("https://example.com", []byte("testdata"))

	val, ok := cache.Get("https://example.com")
	if !ok {
		t.Fatal("expected to find key")
	}
	if string(val) != "testdata" {
		t.Errorf("expected testdata, got %s", string(val))
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("testdata"))

	time.Sleep(interval * 3)

	_, ok := cache.Get("https://example.com")
	if ok {
		t.Error("expected entry to be reaped")
	}
}