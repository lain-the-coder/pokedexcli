package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	key := "https://pokeapi.co/api/v2/location-area/1"
	val := []byte("bulbasaur")

	cache.Add(key, val)

	actual, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected to find key %q in cache", key)
	}

	if string(actual) != string(val) {
		t.Errorf("expected value %q, got %q", string(val), string(actual))
	}
}

func TestReapLoop(t *testing.T) {
	// Use a tiny interval for testing so the test runs quickly!
	const interval = 10 * time.Millisecond
	cache := NewCache(interval)

	key := "https://pokeapi.co/api/v2/location-area/2"
	val := []byte("charmander")

	cache.Add(key, val)

	// Wait slightly longer than the expiration interval
	time.Sleep(interval + (5 * time.Millisecond))

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key %q to be reaped/deleted from cache", key)
	}
}
