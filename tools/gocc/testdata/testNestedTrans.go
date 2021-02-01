package main

import (
	"sync"
)

func main() {
	m := &sync.Mutex{}
	n := &sync.Mutex{}
	count := 0
	m.Lock()
	n.Lock()
	count++
	n.Unlock()
	m.Unlock()
}
