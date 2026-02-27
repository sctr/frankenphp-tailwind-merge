package twmerge

import (
	"fmt"
	"sync"
	"testing"
)

func TestCache_SetAndGet(t *testing.T) {
	c := NewLRUCache(10)
	c.Set("key1", "value1")
	c.Set("key2", "value2")

	val, ok := c.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("expected value1, got %q (ok=%v)", val, ok)
	}

	val, ok = c.Get("key2")
	if !ok || val != "value2" {
		t.Errorf("expected value2, got %q (ok=%v)", val, ok)
	}
}

func TestCache_MissReturnsEmpty(t *testing.T) {
	c := NewLRUCache(10)

	val, ok := c.Get("nonexistent")
	if ok || val != "" {
		t.Errorf("expected empty miss, got %q (ok=%v)", val, ok)
	}
}

func TestCache_Eviction(t *testing.T) {
	c := NewLRUCache(3)

	// Fill main cache to capacity
	c.Set("a", "1")
	c.Set("b", "2")
	c.Set("c", "3")

	// This should trigger rotation: main → previous, new main
	c.Set("d", "4")

	// "d" should be in the new main cache
	val, ok := c.Get("d")
	if !ok || val != "4" {
		t.Errorf("expected d=4, got %q (ok=%v)", val, ok)
	}

	// "a" should be promoted from previous cache
	val, ok = c.Get("a")
	if !ok || val != "1" {
		t.Errorf("expected a=1 from previous, got %q (ok=%v)", val, ok)
	}
}

func TestCache_TierRotation(t *testing.T) {
	c := NewLRUCache(2)

	// Fill and rotate
	c.Set("a", "1")
	c.Set("b", "2")
	c.Set("c", "3") // triggers rotation: {a,b,c} → previous, main={} then c stored... actually main gets {a,b,c} which > 2, so previous={a,b,c}, main={}

	// a and b are in previous, c triggered rotation and is now in main... wait let's trace:
	// Set c: cache={a:1, b:2, c:3}, len=3 > 2, so previous={a:1,b:2,c:3}, cache={}
	// So c is in previous, NOT in main

	// Get c: not in main (empty), found in previous → promotes to main
	val, ok := c.Get("c")
	if !ok || val != "3" {
		t.Errorf("expected c=3, got %q (ok=%v)", val, ok)
	}

	// Get a: not in main, found in previous → promotes to main
	val, ok = c.Get("a")
	if !ok || val != "1" {
		t.Errorf("expected a=1, got %q (ok=%v)", val, ok)
	}

	// Now main={c:3, a:1}, previous still has {a:1,b:2,c:3}
	// Adding new item to trigger another rotation
	c.Set("x", "10") // main={c:3, a:1, x:10} len=3 > 2 → rotate

	// b was in previous, now previous is the old main {c,a,x}
	// b is gone from both caches
	val, ok = c.Get("b")
	if ok {
		t.Errorf("expected b to be evicted, got %q", val)
	}
}

func TestCache_ConcurrentAccess(t *testing.T) {
	c := NewLRUCache(100)
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			c.Set(key, fmt.Sprintf("val%d", i))
		}(i)
	}

	wg.Wait()

	// Concurrent reads
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			c.Get(key)
		}(i)
	}

	wg.Wait()
}

func TestCache_ZeroSize(t *testing.T) {
	c := NewLRUCache(0)

	c.Set("key", "value")

	val, ok := c.Get("key")
	if ok || val != "" {
		t.Errorf("zero-size cache should never return values, got %q (ok=%v)", val, ok)
	}
}

func TestCache_NegativeSize(t *testing.T) {
	c := NewLRUCache(-5)

	c.Set("key", "value")

	val, ok := c.Get("key")
	if ok || val != "" {
		t.Errorf("negative-size cache should never return values, got %q (ok=%v)", val, ok)
	}
}
