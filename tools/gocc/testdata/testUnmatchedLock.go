package main

import (
	"sync"
)

func foo(l *sync.Mutex) {
	l.Lock()
}

func main() {
	m := &sync.Mutex{}
	foo(m)
	m.Unlock()
	m.Lock()
	m.Unlock()
}
