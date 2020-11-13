package main

import (
	"sync"
)

func main() {
	m := sync.Mutex{}
	m.Lock()
	go func() {

	}()
	m.Unlock()
}
