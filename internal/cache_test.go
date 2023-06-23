package internal

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	const basetime = 5 * time.Millisecond
	const waitTime = basetime + 5*time.Millisecond
	cache := Cache{}
	cache.NewCache(basetime)
	cache.add("test", []byte("test"))
	_, ok := cache.get("test")
	if !ok {
		t.Errorf("Expected to find a key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.get("test")
	if ok {
		t.Errorf("expected to not find a key")
		return
	}
}