//go:build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter aman digunakan secara konkuren.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc meningkatkan nilai dari key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock sehingga hanya satu goroutine pada satu waktu yang dapat
	// mengakses map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value mengembalikan nilai dari key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock sehingga hanya satu gorouting pada satu waktu yang dapat
	// mengakses map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
