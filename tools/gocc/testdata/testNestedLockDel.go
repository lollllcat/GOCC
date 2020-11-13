package main

import (
	"sync"
)

func main() {
	m := sync.Mutex{}
	m.Lock()
	func() {
		n := sync.Mutex{}
		n.Lock()
		n.Unlock()
	}()
	m.Unlock()
}
